package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/xvbnm48/golang-ecommerce/controllers"
	"github.com/xvbnm48/golang-ecommerce/database"
	"github.com/xvbnm48/golang-ecommerce/middlewares"
	"github.com/xvbnm48/golang-ecommerce/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middlewares.Authication())

	router.GET("/addtochart", app.AddToChart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))

}
