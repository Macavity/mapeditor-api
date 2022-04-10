package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

type Server struct {
	DB  *gorm.DB
	Env string
}

func (server *Server) Run() {
	server.initDotEnv()
	server.initLogs()
	server.initRoutes()
	server.initDB(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
}

func (server *Server) initRoutes() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	err := router.Run()

	if err != nil {
		log.Warningln(err)
		return
	}
}

func (server *Server) initDB(Driver, User, Password, Port, Host, Name string) {
	var err error

	if Driver == "postgres" {
		url := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", Host, Port, User, Name, Password)
		server.DB, err = gorm.Open(Driver, url)
		if err != nil {
			log.Warnf("Cannot connect to %s database", Driver)
			log.Fatal("This is the error:", err)
		} else {
			log.Infof("We are connected to the %s database", Driver)
		}
	}
}

func (server *Server) initDotEnv() {
	env := os.Getenv("APP_ENV")
	if "" == env {
		env = "development"
	}
	server.Env = env

	_ = godotenv.Load(".env." + env + ".local")

	if "test" != env {
		_ = godotenv.Load(".env.local")
	}

	_ = godotenv.Load(".env." + env)

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Loading the .env file failed.")
		return
	}
}

func (server *Server) initLogs() {
	var file, err = os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	log.SetOutput(file)
}
