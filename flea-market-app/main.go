package main

import (
	"flea-market-app/controllers"
	"flea-market-app/models"
	"flea-market-app/repositories"
	"flea-market-app/services"

	"github.com/gin-gonic/gin"
)

func main() {
	items := []models.Item{
		{ID: 1, Name: "Item1", Price: 1000, Description: "Description1", SoldOut: false},
		{ID: 2, Name: "Item2", Price: 2000, Description: "Description2", SoldOut: true},
		{ID: 3, Name: "Item3", Price: 1000, Description: "Description3", SoldOut: false},
	}

	itemRepository := repositories.NewItemMemoryRepository(items)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	r := gin.Default()
	r.GET("/items", itemController.FindAll)
	r.GET("/items/:id", itemController.FindById)
	r.Run() // listen and serve on 0.0.0.0:8080
}
