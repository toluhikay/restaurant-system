package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/toluhikay/resturant-system/cmd/database"
	"github.com/toluhikay/resturant-system/cmd/middleware"
	"github.com/toluhikay/resturant-system/cmd/routes"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	fmt.Println("Building Restaurant management system with golang, Jehovah please help me")
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	//using the gin-gonic http framework
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)

	// check middleware for authentication
	router.Use(middleware.Authentication())

	// other routes
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.InvoiceRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRouter(router)
}
