package main

import (
	"flea-market-app/controllers"
	"flea-market-app/infra"
	"flea-market-app/middlewares"

	// "flea-market-app/models"
	"flea-market-app/repositories"
	"flea-market-app/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	// items := []models.Item{
	// 	{ID: 1, Name: "Item1", Price: 1000, Description: "Description1", SoldOut: false},
	// 	{ID: 2, Name: "Item2", Price: 2000, Description: "Description2", SoldOut: true},
	// 	{ID: 3, Name: "Item3", Price: 1000, Description: "Description3", SoldOut: false},
	// }

	// itemRepository := repositories.NewItemMemoryRepository(items)
	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	r := gin.Default()
	// NOTE: This is a local development setting. You should specify cors settings in production.
	r.Use(cors.Default())
	itemRouter := r.Group("/items")
	itemRouterWithAuth := r.Group("/items", middlewares.AuthMiddleware(authService))
	authRouter := r.Group("/auth")

	itemRouter.GET("", itemController.FindAll)
	itemRouterWithAuth.GET("/:id", itemController.FindById)
	itemRouterWithAuth.POST("", itemController.Create)
	itemRouterWithAuth.PUT("/:id", itemController.Update)
	itemRouterWithAuth.DELETE("/:id", itemController.Delete)

	authRouter.POST("/signup", authController.Signup)
	authRouter.POST("/login", authController.Login)

	return r
}

func main() {
	infra.Initialize()

	db := infra.SetupDB()
	r := setupRouter(db)

	r.Run(":8082")
}
