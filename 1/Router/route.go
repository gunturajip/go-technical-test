package Router

import (
	"1/Controller"
	"1/Middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Mygram (Simple E-Commerce)
// @version 1.0
// @description This is a simple e-commerce API
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func StartApp() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Simple E-Commerce API")
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	guestRouter := r.Group("/users")
	{
		guestRouter.POST("/register", Controller.UserRegister)
		guestRouter.POST("/login", Controller.UserLogin)
	}

	r.Use(Middleware.Authentication())
	{
		productRouter := r.Group("/products")
		{
			productRouter.GET("/", Controller.GetProducts)
			productRouter.POST("/", Controller.CreateProduct)
			productRouter.GET("/:productID", Controller.GetProduct)
			productRouter.PUT("/:productID", Middleware.ProductAuthorization(), Controller.UpdateProduct)
			productRouter.DELETE("/:productID", Middleware.ProductAuthorization(), Controller.DeleteProduct)
		}

		orderRouter := r.Group("/orders")
		{
			orderRouter.GET("/", Controller.GetOrders)
			orderRouter.POST("/", Controller.CreateOrder)
			orderRouter.GET("/:orderID", Controller.GetOrder)
			orderRouter.PUT("/:orderID", Middleware.OrderAuthorization(), Controller.UpdateOrder)
			orderRouter.DELETE("/:orderID", Middleware.OrderAuthorization(), Controller.DeleteOrder)
		}
	}

	return r
}
