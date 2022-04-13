package TileMaps

import (
	"Macavity/mapeditor-server/server/TileMaps/models"
	"Macavity/mapeditor-server/server/logwrapper"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"net/http"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func RegisterTileMapRoutes(route *gin.Engine) {
	route.GET("/tile-maps", getAll)
	route.GET("/tile-maps/:uuid", getById)
	route.POST("/tile-maps", create)
}

func getAll(c *gin.Context) {
	log := logwrapper.NewDebugLogger()
	tileMaps, err := FindAllTileMaps()

	if err != nil {
		log.Errorln(err)
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"tileMaps": tileMaps})
}

func create(c *gin.Context) {
	var input models.CreateTileMapDTO
	log := logwrapper.NewDebugLogger()
	validate = validator.New()

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := validate.Struct(input); err != nil {
		log.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	tileMap := models.TileMap{
		Name:   input.Name,
		UserID: input.UserID,
		Width:  input.Width,
		Height: input.Height,
	}
	err := StoreTileMap(&tileMap)
	if err != nil {
		log.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tileMap": &tileMap})
}

func getById(c *gin.Context) {
	tileMapId, err := uuid.Parse(c.Param("uuid"))
	tileMap, err := FindTileMapById(tileMapId)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"tileMap": &tileMap})
}
