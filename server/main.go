package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	FB "server/firebaseconfig"
)

func main() {
	FB.ConnectDB()
	defer FB.Client.Close()
	//connect to firebase database
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:8081"},
		AllowMethods:  []string{"PUT", "POST", "GET", "DELETE"},
		AllowHeaders:  []string{"Content-Type", "Content-Length", "Accept-Encoding"},
		ExposeHeaders: []string{"Content-Length", "Content-Type"},
	}))
	//// gin framework
	//// CORS
	api := router.Group("/root")
	{
		api.GET("/getItemAll", func(ctx *gin.Context) {})
	}
	////routes for API
	router.Run(":8080")

}
