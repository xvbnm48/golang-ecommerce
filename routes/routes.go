package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xvbnm48/golang-ecommerce/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp)
	incomingRoutes.POST("/users/login", controllers.login)
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin)
	incomingRoutes.GET("/users/productreview", controllers.GetProductReview)
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery)
}
