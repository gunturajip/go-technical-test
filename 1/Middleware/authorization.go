package Middleware

import (
	"1/Config"
	"1/Model"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		userData := c.MustGet("userData").(jwt.MapClaims)
		admin := userData["admin"].(bool)
		if !admin {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to update / delete this product",
			})
			return
		}
		c.Next()
	}
}

func OrderAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := Config.GetDB()
		orderID, err := strconv.Atoi(c.Param("orderID"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid photo parameter",
			})
			return
		}

		Order := Model.Order{}
		err = db.First(&Order, uint(orderID)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": "order doesn't exist",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := userData["id"].(float64)
		if Order.UserID != uint(userID) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to update / delete this order",
			})
			return
		}
		c.Next()
	}
}
