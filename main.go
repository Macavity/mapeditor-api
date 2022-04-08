package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	initLogs()
	initDotEnv()
	//server.InitDB(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	initRoutes()
	fmt.Println("Map-Editor Server started.")
}

func initRoutes() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	err := router.Run()

	if err != nil {
		log.Warningln(err)
		return
	}
}

func initDotEnv() {
	env := os.Getenv("APP_ENV")
	if "" == env {
		env = "development"
	}

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

func initLogs() {
	var file, err = os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	log.SetOutput(file)
}
