package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/xvbnm48/golang-ecommerce/database"
	"github.com/xvbnm48/golang-ecommerce/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection = database.UserData(database.Client, "Users")
var ProductCollection *mongo.Collection = database.ProductData(database.Client, "Products")
var Validate = validator.New()

func HashPassword(password string) string {

}

func VerifyPassword(userPassword, givenPassword string) (bool, string) {

}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		ValidationErr := Validate.Struct(user)
		if ValidationErr != nil {
			c.JSON(400, gin.H{
				"error": ValidationErr,
			})
			return
		}
		count, err := UserCollection.CountDocument(ctx, bson.M{
			"email": user.Email,
		})

		if err != nil {
			log.Panic(err)
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		if count > 0 {
			c.JSON(400, gin.H{
				"error": "User already exists",
			})
		}
		count, err = UserCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})

		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		if count > 0 {
			c.JSON(400, gin.H{
				"error": "Phone already exists",
			})
			return
		}

		password := HashPassword(*user.Password)
		user.Password = &password
		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()

		token, refreshtoken, _ := generate.TokenGenerator(*user.Mail, *user.Firts_Name, *user.Last_Name, user.user_ID)
		user.Token = &token
		user.Refresh_Token = &refreshtoken
		user.UserCart = make([]models.ProductUser, 0)
		user.Address_Details = make([]models.Address, 0)
		user.Order_Status = make([]models.Order, 0)
		_, inserterr := UserCollection.InsertOne(ctx, user)
		if inserterr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "the user did not get created",
			})
			return
		}
		defer cancel()
		c.JSON(http.StatusCreated, "successfully signed in!")
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user = models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error", err,
			})
			return
		}

		err := UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&founduser)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "login or passsword is incorrect",
			})
			return
		}
		passwordInvalid, msg := VerifyPassword(*user.Password, *founduser.Password)
		defer cancel()
		if !passwordInvalid {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": msg,
			})
			fmt.Println(msg)
			return
		}
		token, refreshToken, _ := generate.TokenGenerator(*founduser.Email, *founduser.Firts_Name, *founeduser.Last_name, founduser.user.ID)
		defer cancel()

		generate.UpdateAllToken(token, refreshToken, founduser.User_ID)

		c.JSON(http.StatusFound, founduser)

	}
}

func ProductViewerAdmin() gin.HandlerFunc {

}

func SearchProductByQuery() gin.HandlerFunc {

}

func SearchProduct() gin.HandlerFunc {

}
