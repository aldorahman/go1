package main

import (
	"go1/config"
	"go1/route"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
)

func main() {
	defer config.CloseDatabaseConnection(db)

	r := gin.Default()

	r.Use(gin.Logger())
	r.LoadHTMLGlob("static/*.html")
	r.Static("/static", "static")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	route.Routes(r)

	port := os.Getenv("PORT")

	r.Run(":" + port)
}
