package handler

import (
	"log"
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
	log.Println("Order ID received:", orderID)
	// params, err := strconv.Atoi(orderID)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order_id"})
	// 	return
	// }

	var req input.SubmitPaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	log.Printf("Request payload: %+v", req)

	currentUser := c.MustGet("currentUser").(*entity.User)
	//ini inisiasi userID yang mana ingin mendapatkan id si user
	getUserId := currentUser.ID

	resp, err := h.paymentService.DoPayment(req, orderID, getUserId)
	if err != nil {
		log.Printf("Error during payment: %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Payment response: %+v", formatter.FormatSimplePaymentResponse(resp))
	c.JSON(http.StatusOK, resp)
}

func (h *paymentHandler) GetStatusPayment(c *gin.Context) {

	orderID := c.Param("orderID")

	get, err := h.paymentService.FindStatus(orderID)

	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, get)
	c.JSON(http.StatusOK, response)
}

// will be called by midtrans to notify payment status
func (h *paymentHandler) GetPaymentNotification(c *gin.Context) {
	// orderID := c.Param("order_id")

	// params, err := strconv.Atoi(orderID)

	var input *entity.MidtransNotificationRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	err := h.paymentService.HandleNotification(input)
	if err != nil {
		log.Printf("Error during payment: %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "succes")

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
	response := helper.APIresponse(http.StatusOK, formatter.FormatterGetPayments(getOrder))
	c.JSON(http.StatusOK, response)
}

func (h *paymentHandler) GetAllPayment(c *gin.Context) {

	getOrder, err := h.paymentService.GetAllPayment()
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, formatter.FormatterGetPayments(getOrder))
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
