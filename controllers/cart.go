package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xvbnm48/golang-ecommerce/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	prodColletion  *mongo.Collection
	userCollection *mongo.Collection
}

func NewApplication(prodColletion, userCollection *mongo.Collection) *Application {
	return &Application{
		prodColletion:  prodColletion,
		userCollection: userCollection,
	}
}

func (app *Application) AddToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Query("id")
		if productQueryID == "" {
			log.Println("Product ID is empty")

			_ = c.AbortWithError(http.StatusInternalServerError, errors.New("Product ID is empty"))
			return
		}

		userQueryID := c.Query("user_id")
		if userQueryID == "" {
			log.Println("User ID is empty")

			_ = c.AbortWithError(http.StatusInternalServerError, errors.New("User ID is empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil {
			log.Println("Product ID is invalid")

			_ = c.AbortWithError(http.StatusInternalServerError, errors.New("Product ID is invalid"))
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = database.AddProductToCart(ctx, app.prodColletion, app.userCollection, productID, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}

		c.IndentedJSON(200, "successfully added to cart")
	}
}

func (app *Application) RemoveItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Query("id")
		if productQueryID == "" {
			log.Println("Product ID is empty")

			_ = c.AbortWithError(http.StatusInternalServerError, errors.New("Product ID is empty"))
			return
		}

		userQueryID := c.Query("user_id")
		if userQueryID == "" {
			log.Println("User ID is empty")

			_ = c.AbortWithError(http.StatusInternalServerError, errors.New("User ID is empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil {
			log.Println("Product ID is invalid")

			_ = c.AbortWithError(http.StatusInternalServerError, errors.New("Product ID is invalid"))
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = database.RemoveCartItem(ctx, app.prodColletion, app.userCollection, productID, userQueryID)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}

		c.IndentedJSON(200, "successfully removed from cart")
	}
}

func GetItemFromCart() gin.HandlerFunc {

}

func (app *Application) BuyFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		userQueryID := c.Query("ID")

		if userQueryID == "" {
			log.Panicln("User ID is empty")

			_ = c.AbortWithError(http.StatusInternalServerError, errors.New("User ID is empty"))
			return
		}

		context, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()

		err := database.BuyItemFromCart(ctx, app.prodColletion, app.userCollection, userQueryID)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}

		c.IndentedJSON(200, "successfully bought item")
	}
}

func (app *Application) InstantBuy() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Query("id")
		if productQueryID == "" {
			log.Println("Product ID is empty")

			_ = c.AbortWithError(http.StatusInternalServerError, errors.New("Product ID is empty"))
			return
		}

		userQueryID := c.Query("user_id")
		if userQueryID == "" {
			log.Println("User ID is empty")

			_ = c.AbortWithError(http.StatusInternalServerError, errors.New("User ID is empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil {
			log.Println("Product ID is invalid")

			_ = c.AbortWithError(http.StatusInternalServerError, errors.New("Product ID is invalid"))
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = database.InstantBuy(ctx, app.prodColletion, app.userCollection, productID, userQueryID)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}

		c.IndentedJSON(200, "successfully placed the order")
	}
}
