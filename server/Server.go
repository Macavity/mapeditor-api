package server

import (
	"Macavity/mapeditor-server/server/Users"
	"Macavity/mapeditor-server/server/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	database.Connect(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	database.MigrateDatabase()
}

func (server *Server) initRoutes() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	Users.RegisterUserRoutes(router)

	err := router.Run()

	if err != nil {
		log.Warningln(err)
		return
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
		fmt.Println(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	log.SetOutput(file)
	log.SetLevel(log.DebugLevel)
	log.Infoln("Logs initialised.")
}
