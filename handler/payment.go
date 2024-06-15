package handler

import (
	"net/http"
	"payment-gwf/auth"
	"payment-gwf/entity"
	"payment-gwf/helper"
	"payment-gwf/input"
	"payment-gwf/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type paymentHandler struct {
	paymentService service.ServicePayment
	authService    auth.Service
}

func NewPaymentHandler(paymentService service.ServicePayment, authService auth.Service) *paymentHandler {
	return &paymentHandler{paymentService, authService}
}

// will be called by user through payment endpoint
func (h *paymentHandler) DoPayment(c *gin.Context) {
	//1. get order data
	orderID := c.Param("order_id")
	params, err := strconv.Atoi(orderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order_id"})
		return
	}

	var req input.SubmitPaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	resp, err := h.paymentService.DoPayment(req, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// will be called by midtrans to notify payment status
func (h *paymentHandler) GetPaymentNotification(c *gin.Context) {
	//1. get order data from db
	//2. check request transaction_status
	//3. map transaction_status to db payment status
	//4. update db payment status
}

func (h *paymentHandler) GetAllPaymentByUserID(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(*entity.User)
	//ini inisiasi userID yang mana ingin mendapatkan id si user
	getUserId := currentUser.ID

	getOrder, err := h.paymentService.GetAllByUserID(getUserId)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, (getOrder))
	c.JSON(http.StatusOK, response)
}

func (h *paymentHandler) GetAllPayment(c *gin.Context) {

	getOrder, err := h.paymentService.GetAllPayment()
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, (getOrder))
	c.JSON(http.StatusOK, response)
}

func (h *paymentHandler) DeletePayment(c *gin.Context) {
	param := c.Param("id")
	paramPayment, _ := strconv.Atoi(param)

	_, err := h.paymentService.DeletePayment(paramPayment)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, "psyment has succesfuly deleted")
	c.JSON(http.StatusOK, response)
}
