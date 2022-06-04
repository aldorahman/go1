package main

import (
	"go1/config"
	"go1/route"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
)

func main() {
	defer config.CloseDatabaseConnection(db)

	r := gin.Default()
	route.Routes(r)

	r.Run(":8080")
}
