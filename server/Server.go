package server

import (
	"Macavity/mapeditor-server/server/TileMaps"
	"Macavity/mapeditor-server/server/Users"
	"Macavity/mapeditor-server/server/logwrapper"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"os"
)

type Server struct {
	DB  *gorm.DB
	Env string
}

func (server *Server) Run() {
	logger := logwrapper.NewDebugLogger()
	logger.Infoln("-----------------------------------------------")
	logger.Infoln("Server starting.")
	server.initDotEnv(logger)
	server.initRoutes()
	logger.Debugln("Init DB")
	server.initDatabase()
	logger.Debugln("Run finished")
}

func (server *Server) initRoutes() {
	logger := logwrapper.NewDebugLogger()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	Users.RegisterUserRoutes(router)
	TileMaps.RegisterTileMapRoutes(router)

	err := router.Run()

	if err != nil {
		logger.Errorln(err)
		return
	}
}

func (server *Server) initDatabase() {
	Users.MigrateUserTable()
	TileMaps.MigrateTileMapTables()
}

func (server *Server) initDotEnv(Logger *logwrapper.StandardLogger) {
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
		Logger.Warningln("Loading the .env file failed.")
		return
	}

	Logger.Debug("Environment: ", server.Env)
	Logger.Debugln("Dotenv loaded.")
}
