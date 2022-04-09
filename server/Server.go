package src

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
}

func (server *Server) InitRoutes() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	controllers.HomeRoutes(router)

	err := router.Run()

	if err != nil {
		log.Warningln(err)
		return
	}
}
