package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

var title = "Echo app"

func init() {
	t, found := os.LookupEnv("SITE_TITLE")

	if found {
		title = t
	}
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
		return
	})
	router.GET("/", handleIndex)

	err := router.Run()

	log.Fatal(err)
}

func handleIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":   title,
		"body":    title,
		"request": c.Request,
	})
}
