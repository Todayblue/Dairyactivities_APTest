package main

import (
	"Dairyactivities_API/controller"
	"Dairyactivities_API/entity"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	os.Remove("./nutrition.db")
	r := gin.Default()

	// Connect to database
	entity.SetupDatabase()

	// Routes
	r.POST("/create_activity", controller.CreateActivity)
	r.GET("/activities", controller.ListActivities)
	r.GET("/activity/:id", controller.GetActivity)

	// Run the server
	r.Run()
}
