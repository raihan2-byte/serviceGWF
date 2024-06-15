package handler

import (
	"fmt"
	"net/http"
	"payment-gwf/auth"
	"payment-gwf/entity"
	"payment-gwf/formatter"
	"payment-gwf/helper"
	"payment-gwf/input"
	"payment-gwf/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type cartHandler struct {
	cartService service.ServiceCart
	authService auth.Service
}

func NewCartHandler(cartService service.ServiceCart, authService auth.Service) *cartHandler {
	return &cartHandler{cartService, authService}
}

func (h *cartHandler) AddToCart(c *gin.Context) {
	productID := c.Param("product_id")
	paramProduct, _ := strconv.Atoi(productID)

	var inputQty input.InputCart
	err := c.ShouldBindJSON(&inputQty)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	currentUser := c.MustGet("currentUser").(*entity.User)
	//ini inisiasi userID yang mana ingin mendapatkan id si user
	getUserId := currentUser.ID

	addCart, err := h.cartService.AddToCart(paramProduct, getUserId, inputQty)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, formatter.AddToCartFormatter(addCart))
	c.JSON(http.StatusOK, response)
}

func (h *cartHandler) GetCartByUserID(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(*entity.User)
	//ini inisiasi userID yang mana ingin mendapatkan id si user
	getUserId := currentUser.ID

	getCart, err := h.cartService.GetAllCartByUserId(getUserId)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, formatter.FormatterGetAllCartByUser(getCart))
	c.JSON(http.StatusOK, response)
}

func (h *cartHandler) UpdatedCartByID(c *gin.Context) {
	cartID := c.Param("cart_id")
	paramCart, _ := strconv.Atoi(cartID)

	var inputQty input.InputCart
	err := c.ShouldBindJSON(&inputQty)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(*entity.User)
	//ini inisiasi userID yang mana ingin mendapatkan id si user
	getUserId := currentUser.ID

	updatedCart, err := h.cartService.UpdatedCart(paramCart, getUserId, inputQty)

	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, formatter.AddToCartFormatter(updatedCart))
	c.JSON(http.StatusOK, response)
}

func (h *cartHandler) DeletedCartByID(c *gin.Context) {
	// cartID := c.Param("cart_id")

	// paramCart, _ := strconv.Atoi(cartID)

	var paramCart input.GetID
	err := c.ShouldBindUri(&paramCart)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(*entity.User)
	//ini inisiasi userID yang mana ingin mendapatkan id si user
	getUserId := currentUser.ID

	DeletedCartByID, err := h.cartService.DeleteCart(getUserId, paramCart)
	fmt.Println("iniparam")
	fmt.Println(paramCart)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, DeletedCartByID.ID)
	c.JSON(http.StatusOK, response)
}
