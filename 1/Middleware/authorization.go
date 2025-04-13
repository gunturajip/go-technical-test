package Middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		productID, err := strconv.Atoi(c.Param("productID"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid photo parameter",
			})
			return
		}

		Photo := models.Photo{}
		err = db.First(&Photo, uint(productID)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": "photo doesn't exist",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		admin := userData["admin"].(bool)
		if !admin {
			userID := uint(userData["id"].(float64))
			if Photo.UserID != userID {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "Unauthorized",
					"message": "you are not allowed to update / delete this photo",
				})
				return
			}
		}
		c.Next()
	}
}
