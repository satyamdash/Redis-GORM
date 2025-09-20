package main

import (
	"github.com/gin-gonic/gin"
	"github.com/satyamdash/Redis-GORM/config"
	"github.com/satyamdash/Redis-GORM/models"
	"github.com/satyamdash/Redis-GORM/routes"
)

func main() {
	config.ConnectDB()

	config.InitRedis()

	config.DB.AutoMigrate(&models.Post{})

	r := gin.Default()

	routes.PostRoutes(r)

	r.Run(":8080")

}
