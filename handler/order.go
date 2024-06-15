package handler

import (
	"net/http"
	"payment-gwf/auth"
	"payment-gwf/entity"
	"payment-gwf/formatter"
	"payment-gwf/helper"
	"payment-gwf/input"
	"payment-gwf/service"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	orderService service.ServiceOrder
	authService  auth.Service
}

func NewOrderHandler(orderService service.ServiceOrder, authService auth.Service) *orderHandler {
	return &orderHandler{orderService, authService}
}

func (h *orderHandler) CreateOrder(c *gin.Context) {
	var inputOrder input.CreateOrder

	err := c.ShouldBindBodyWithJSON(&inputOrder)
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

	create, err := h.orderService.CreateOrders(getUserId, inputOrder)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, formatter.FormatterPostOrder(create))
	c.JSON(http.StatusOK, response)
}

func (h *orderHandler) GetOrderHistoryByUserID(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(*entity.User)
	//ini inisiasi userID yang mana ingin mendapatkan id si user
	getUserId := currentUser.ID

	getOrder, err := h.orderService.GetOrderHistoryByUserID(getUserId)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, formatter.FormatterGetAllOrder(getOrder))
	c.JSON(http.StatusOK, response)
}

func (h *orderHandler) GetAllOrderHistory(c *gin.Context) {

	getOrder, err := h.orderService.GetAllOrderHistory()
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, formatter.FormatterGetAllOrder(getOrder))
	c.JSON(http.StatusOK, response)
}
