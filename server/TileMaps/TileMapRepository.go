package TileMaps

import (
	"Macavity/mapeditor-server/server/TileMaps/models"
	"Macavity/mapeditor-server/server/database"
	"Macavity/mapeditor-server/server/logwrapper"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func createDatabaseClient() *gorm.DB {
	MigrateTileMapTables()
	return database.NewConnection()
}

func MigrateTileMapTables() {
	client := database.NewConnection()
	logger := logwrapper.NewLogger(logrus.DebugLevel)
	logger.Debugln("[TileMaps.Migrate] Start")

	if client.Migrator().HasTable(&models.Layer{}) == false {
		logger.Debugln("[TileMaps.Migrate] Create Layer Table")
		err := client.Migrator().CreateTable(&models.Layer{})
		if err != nil {
			logger.Errorln("[TileMaps.Migrate] Layer Table creation failed", err)
			return
		}
	}

	if client.Migrator().HasTable(models.TileMap{}) == false {
		logger.Debugln("[TileMaps.Migrate] Create TileMap Table")
		err := client.Migrator().CreateTable(&models.TileMap{})
		if err != nil {
			logger.Errorln("[TileMaps.Migrate] Table creation failed", err)
			return
		}
	}

	err := client.AutoMigrate(&models.TileMap{}, &models.Layer{})
	if err != nil {
		logger.Errorln("[TileMaps.Migrate] AutoMigrate failed", err)
		return
	}
}

func FindAllTileMaps() (*[]models.TileMap, error) {
	var err error
	var tileMaps []models.TileMap
	logger := logwrapper.NewDebugLogger()
	logger.Debugln("FindAllTileMaps")
	client := createDatabaseClient()

	err = client.Model(&models.TileMap{}).Limit(20).Find(&tileMaps).Error

	if err != nil {
		logger.Errorln(err)
		return &[]models.TileMap{}, err
	}

	return &tileMaps, err
}

func StoreTileMap(tileMap *models.TileMap) (err error) {
	client := createDatabaseClient()
	if err != nil {
		log.Println(err)
		return err
	}

	err = client.Model(&models.TileMap{}).Create(tileMap).Error

	return err
}

func FindTileMapById(tileMapId uuid.UUID) (tileMap models.TileMap, err error) {
	client := createDatabaseClient()
	result := client.First(&tileMap, tileMapId)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err = fmt.Errorf("TileMap with ID %d not found", tileMapId)
	}

	return tileMap, err
}
