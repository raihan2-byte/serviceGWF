package handler

import (
	"net/http"
	"payment-gwf/helper"
	"payment-gwf/input"
	"payment-gwf/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService service.ServiceProduct
}

func NewProductHandler(service service.ServiceProduct) *productHandler {
	return &productHandler{service}
}

func (h *productHandler) CreateProduct(c *gin.Context) {
	var input input.ProductInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newProduct, err := h.productService.CreateProduct(input)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := (newProduct)
	response := helper.APIresponse(http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) UpdateProduct(c *gin.Context) {
	var input input.ProductInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	getID := c.Param("id")
	param, _ := strconv.Atoi(getID)
	newProduct, err := h.productService.UpdateProduct(param, input)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := (newProduct)
	response := helper.APIresponse(http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) GetAllProduct(c *gin.Context) {
	products, err := h.productService.GetProducts()
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Eror to get product")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, (products))
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) GetProduct(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	products, err := h.productService.GetProduct(int(id))
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Eror to get product")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, (products))
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) DeleteProduct(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	_, err := h.productService.DeleteProduct(int(id))
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Eror to get product")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, "products has succesfuly deleted")
	c.JSON(http.StatusOK, response)
}
