package Router

import (
	"1/Controller"
	"1/Middleware"
	"net/http"

	"github.com/gin-gonic/gin"
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

	// Read homepage content
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Simple E-Commerce API")
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	guestRouter := r.Group("/users")
	{
		// Register user
		guestRouter.POST("/register", Controller.UserRegister)

		// Login user
		guestRouter.POST("/login", Controller.UserLogin)
	}

	r.Use(Middleware.Authentication())
	{
		productRouter := r.Group("/products")
		{
			productRouter.GET("/", controllers.GetProducts)
			productRouter.POST("/", controllers.CreateProduct)
			productRouter.GET("/:productID", controllers.GetProduct)
			productRouter.PUT("/:productID", middlewares.PhotoAuthorization(), controllers.UpdateProduct)
			productRouter.DELETE("/:productID", middlewares.PhotoAuthorization(), controllers.DeleteProduct)
		}

		commentRouter := r.Group("/comments")
		{
			// Get all comments
			commentRouter.GET("/", controllers.GetComments)

			// Post comment
			commentRouter.POST("/", controllers.CreateComment)

			// Get comment by id
			commentRouter.GET("/:commentID", controllers.GetComment)

			// Update comment by id
			commentRouter.PUT("/:commentID", middlewares.CommentAuthorization(), controllers.UpdateComment)

			// Delete comment by id
			commentRouter.DELETE("/:commentID", middlewares.CommentAuthorization(), controllers.DeleteComment)
		}

		socialMediaRouter := r.Group("/socialmedia")
		{
			// Get all social media
			socialMediaRouter.GET("/", controllers.GetSocialMedias)

			// Post social media
			socialMediaRouter.POST("/", controllers.CreateSocialMedia)

			// Get social media by id
			socialMediaRouter.GET("/:socialMediaID", controllers.GetSocialMedia)

			// Update social media by id
			socialMediaRouter.PUT("/:socialMediaID", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)

			// Delete social media by id
			socialMediaRouter.DELETE("/:socialMediaID", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
		}
	}

	return r
}
