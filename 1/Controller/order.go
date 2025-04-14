package Controller

import (
	"1/Config"
	"1/Model"
	"1/Util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetOrders godoc
// @Summary Get all orders
// @Description Get all orders data
// @Tags comment
// @Accept json
// @Produce json
// @Param Authorization header string true "Type Bearer your_token"
// @Success 200 {object} Model.Response
// @Failure 400 {object} Model.Response
// @Router /orders [get]
func GetOrders(c *gin.Context) {
	db := Config.GetDB()
	response := Model.Response{}

	Orders := []Model.Order{}

	err := db.Preload("Items").Find(&Orders).Error
	if err != nil {
		response.Error = "Bad Request"
		response.Message = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response.Data = Orders
	c.JSON(http.StatusOK, response)
}

// GetOrder godoc
// @Summary Get order details for the given ID
// @Description Get details of a order corresponding to the input ID
// @Tags order
// @Accept json
// @Produce json
// @Param Authorization header string true "Type Bearer your_token"
// @Param ID path int true "ID of the order"
// @Success 200 {object} Model.Response
// @Failure 400 {object} Model.Response
// @Router /orders/{ID} [get]
func GetOrder(c *gin.Context) {
	db := Config.GetDB()
	response := Model.Response{}

	Order := Model.Order{}
	orderID := c.Param("orderID")

	err := db.Preload("Items").First(&Order, "id = ?", orderID).Error
	if err != nil {
		response.Error = "Bad Request"
		response.Message = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response.Data = Order
	c.JSON(http.StatusOK, response)
}

// CreateOrder godoc
// @Summary Post a new order
// @Description Post details of a new order
// @Tags order
// @Accept json
// @Produce json
// @Param Authorization header string true "Type Bearer your_token"
// @Param Model.Order body Model.Order true "create a order"
// @Success 201 {object} Model.Response
// @Failure 400 {object} Model.Response
// @Router /orders [post]
func CreateOrder(c *gin.Context) {
	db := Config.GetDB()
	response := Model.Response{}

	Order := Model.Order{}
	contentType := Util.GetContentType(c)
	if contentType == appJSON {
		c.ShouldBindJSON(&Order)
	} else {
		c.ShouldBind(&Order)
	}

	if Order.PurchaseProofLink != "" {
		Order.Status = "processing"
	} else {
		Order.Status = "draft"
	}

	errOrder := db.Create(&Order).Error
	if errOrder != nil {
		response.Error = "Bad Request"
		response.Message = errOrder.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response.Data = Order
	c.JSON(http.StatusCreated, response)
}

// UpdateOrder godoc
// @Summary Update order for the given ID
// @Description Update details of a order corresponding to the input ID
// @Tags order
// @Accept json
// @Produce json
// @Param Authorization header string true "Type Bearer your_token"
// @Param ID path int true "ID of the order"
// @Param Model.Order body Model.Order true "update order"
// @Success 200 {object} Model.Response
// @Failure 400 {object} Model.Response
// @Router /orders/{ID} [put]
func UpdateOrder(c *gin.Context) {
	db := Config.GetDB()
	response := Model.Response{}

	Order := Model.Order{}
	orderID := c.Param("orderID")

	var input struct {
		Address           string       `json:"address"`
		PurchaseProofLink string       `json:"purchase_proof_link"`
		Status            string       `json:"status"`
		UserID            uint         `json:"user_id"`
		Items             []Model.Item `json:"items"`
	}

	contentType := Util.GetContentType(c)
	if contentType == appJSON {
		c.ShouldBindJSON(&input)
	} else {
		c.ShouldBind(&input)
	}

	if input.PurchaseProofLink != "" {
		input.Status = "processing"
	} else {
		input.Status = "draft"
	}

	errOrderItems := db.Where("order_id = ?", orderID).Delete(&input.Items).Error
	if errOrderItems != nil {
		response.Error = "Bad Request"
		response.Message = errOrderItems.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	errOrder := db.Model(&Order).Where("id = ?", orderID).Updates(map[string]interface{}{
		"address":             input.Address,
		"purchase_proof_link": input.PurchaseProofLink,
		"status":              input.Status,
	}).Error
	if errOrder != nil {
		response.Error = "Bad Request"
		response.Message = errOrder.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	items := make([]Model.Item, len(input.Items))
	uintOrderID, _ := strconv.ParseUint(orderID, 10, 32)
	for i, item := range input.Items {
		items[i] = Model.Item{
			OrderID:   uint(uintOrderID),
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}
	}

	errItem := db.Create(&items).Error
	if errItem != nil {
		response.Error = "Bad Request"
		response.Message = errItem.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusNoContent, response)
}

// DeleteOrder godoc
// @Summary Delete order details for a given ID
// @Description Delete details of a order corresponding to the input ID
// @Tags order
// @Accept json
// @Produce json
// @Param Authorization header string true "Type Bearer your_token"
// @Param ID path int true "ID of the order"
// @Success 200 {object} Model.Response
// @Failure 400 {object} Model.Response
// @Router /orders/{ID} [delete]
func DeleteOrder(c *gin.Context) {
	db := Config.GetDB()
	response := Model.Response{}

	Order := Model.Order{}
	orderID := c.Param("orderID")

	errOrderItems := db.Where("order_id = ?", orderID).Delete(&Order.Items).Error
	if errOrderItems != nil {
		response.Error = "Bad Request"
		response.Message = errOrderItems.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	errOrder := db.Where("id = ?", orderID).Delete(&Order).Error
	if errOrder != nil {
		response.Error = "Bad Request"
		response.Message = errOrder.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	c.JSON(http.StatusNoContent, response)
}
