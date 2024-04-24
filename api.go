package main

import (
	"api/controllers"
	"api/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	repo := repository.NewDatabaseRepository()
	controllers := controllers.NewController(repo)
	router.POST("/deletedata", controllers.Deletedata)
	router.POST("/adddata", controllers.AddData)
	router.Run()
}
