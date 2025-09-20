package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/satyamdash/Redis-GORM/config"
	"github.com/satyamdash/Redis-GORM/models"
)

func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&post)
	c.JSON(http.StatusOK, post)

}

func GetPost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam) // convert string → int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	key := fmt.Sprintf("post:%d", id)

	// Try Redis cache
	val, err := config.GetCache(key)
	if err == nil {
		var post models.Post
		json.Unmarshal([]byte(val), &post)
		fmt.Println("Cache hit")
		c.JSON(http.StatusOK, post)
		return
	}

	// Fallback: DB
	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Save to cache
	jsonData, _ := json.Marshal(post)
	config.SetCache(key, string(jsonData), 10*time.Minute)
	fmt.Println("Cache miss → saved to Redis")

	c.JSON(http.StatusOK, post)
}
