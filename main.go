package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type JsonRequet struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type SleepTime struct {
	Time int `json:"time"`
}

func main() {
	r := gin.Default()
	r.POST("/notime", func(c *gin.Context) {
		var json JsonRequet
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"title": json.Title, "text": json.Title})
	})
	r.POST("/4min", func(c *gin.Context) {
		var json JsonRequet
		time.Sleep(time.Minute * 4)
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"title": json.Title, "text": json.Title})
	})
	r.POST("/6min", func(c *gin.Context) {
		var json JsonRequet
		time.Sleep(time.Minute * 6)
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"title": json.Title, "text": json.Title})
	})
	r.POST("/10min", func(c *gin.Context) {
		var json JsonRequet
		time.Sleep(time.Minute * 10)
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"title": json.Title, "text": json.Title})
	})
	r.POST("/20min", func(c *gin.Context) {
		var json JsonRequet
		time.Sleep(time.Minute * 20)
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"title": json.Title, "text": json.Title})
	})
	r.POST("/custom-time", func(c *gin.Context) {
		var json SleepTime
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		time.Sleep(time.Second * time.Duration(json.Time))
		c.JSON(http.StatusOK, gin.H{"msg": fmt.Sprintf("sleep for %d seconds", json.Time)})
	})
	r.Run()
}
