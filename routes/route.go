package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/satyamdash/Redis-GORM/controllers"
)

func PostRoutes(r *gin.Engine) {
	r.GET("/posts/:id", controllers.GetPost)
	r.POST("/posts", controllers.CreatePost)
}
