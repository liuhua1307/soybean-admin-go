package main

import (
	"github.com/gin-gonic/gin"
	"soybean-admin-go/config"
	"soybean-admin-go/db"
	"soybean-admin-go/middleware"
	"soybean-admin-go/router"
	"time"
)

func main() {
	var Loc, _ = time.LoadLocation("Asia/Shanghai")
	time.Local = Loc

	app := gin.Default()
	app.Use(middleware.Cors())
	db.Init()
	router.Init(app)
	err := app.Run(":8081")
	if err != nil {
		panic(err)
	}
	config.Logger.Info("Server is running on port 8081")
}
