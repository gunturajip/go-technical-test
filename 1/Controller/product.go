package Controller

import (
	"1/Config"
	"1/Model"
	"1/Util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProducts godoc
// @Summary Get all products
// @Description Get all products data
// @Tags comment
// @Accept json
// @Produce json
// @Param Authorization header string true "Type Bearer your_token"
// @Success 200 {object} Model.Response
// @Failure 400 {object} Model.Response
// @Router /products [get]
func GetProducts(c *gin.Context) {
	db := Config.GetDB()
	response := Model.Response{}

	Products := []Model.Product{}

	err := db.Find(&Products).Error
	if err != nil {
		response.Error = "Bad Request"
		response.Message = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response.Data = Products
	c.JSON(http.StatusOK, response)
}

// GetProduct godoc
// @Summary Get product details for the given ID
// @Description Get details of a product corresponding to the input ID
// @Tags product
// @Accept json
// @Produce json
// @Param Authorization header string true "Type Bearer your_token"
// @Param ID path int true "ID of the product"
// @Success 200 {object} Model.Response
// @Failure 400 {object} Model.Response
// @Router /products/{ID} [get]
func GetProduct(c *gin.Context) {
	db := Config.GetDB()
	response := Model.Response{}

	Product := Model.Product{}
	productID := c.Param("productID")

	err := db.First(&Product, "id = ?", productID).Error
	if err != nil {
		response.Error = "Bad Request"
		response.Message = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response.Data = Product
	c.JSON(http.StatusOK, response)
}

// CreateProduct godoc
// @Summary Post a new product
// @Description Post details of a new product
// @Tags product
// @Accept json
// @Produce json
// @Param Authorization header string true "Type Bearer your_token"
// @Param Model.Product body Model.Product true "create a product"
// @Success 201 {object} Model.Response
// @Failure 400 {object} Model.Response
// @Router /products [post]
func CreateProduct(c *gin.Context) {
	db := Config.GetDB()
	response := Model.Response{}

	Product := Model.Product{}
	contentType := Util.GetContentType(c)
	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	errProduct := db.Create(&Product).Error
	if errProduct != nil {
		response.Error = "Bad Request"
		response.Message = errProduct.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response.Data = Product
	c.JSON(http.StatusCreated, response)
}

// UpdateProduct godoc
// @Summary Update product for the given ID
// @Description Update details of a product corresponding to the input ID
// @Tags product
// @Accept json
// @Produce json
// @Param Authorization header string true "Type Bearer your_token"
// @Param ID path int true "ID of the product"
// @Param Model.Product body Model.Product true "update product"
// @Success 200 {object} Model.Response
// @Failure 400 {object} Model.Response
// @Router /products/{ID} [put]
func UpdateProduct(c *gin.Context) {
	db := Config.GetDB()
	response := Model.Response{}

	Product := Model.Product{}
	productID := c.Param("productID")

	contentType := Util.GetContentType(c)
	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	errProduct := db.Model(&Product).Where("id = ?", productID).Updates(Model.Product{
		Name:  Product.Name,
		Price: Product.Price,
	}).Error
	if errProduct != nil {
		response.Error = "Bad Request"
		response.Message = errProduct.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	c.JSON(http.StatusNoContent, response)
}

// DeleteProduct godoc
// @Summary Delete product details for a given ID
// @Description Delete details of a product corresponding to the input ID
// @Tags product
// @Accept json
// @Produce json
// @Param Authorization header string true "Type Bearer your_token"
// @Param ID path int true "ID of the product"
// @Success 200 {object} Model.Response
// @Failure 400 {object} Model.Response
// @Router /products/{ID} [delete]
func DeleteProduct(c *gin.Context) {
	db := Config.GetDB()
	response := Model.Response{}

	Product := Model.Product{}
	productID := c.Param("productID")

	err := db.Where("id = ?", productID).Delete(&Product).Error
	if err != nil {
		response.Error = "Bad Request"
		response.Message = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	c.JSON(http.StatusNoContent, response)
}
