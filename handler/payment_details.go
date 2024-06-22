package handler

import (
	"log"
	"net/http"
	"payment-gwf/auth"
	"payment-gwf/entity"
	"payment-gwf/service"

	"github.com/gin-gonic/gin"
)

type paymentDetailsHandler struct {
	paymentDetailsService service.ServicePaymentDetails
	authService           auth.Service
}

func NewPaymentDetailsHandler(paymentDetailsService service.ServicePaymentDetails, authService auth.Service) *paymentDetailsHandler {
	return &paymentDetailsHandler{paymentDetailsService, authService}
}

func (h *paymentDetailsHandler) GetPaymentDonationNotification(c *gin.Context) {
	// orderID := c.Param("order_id")

	// params, err := strconv.Atoi(orderID)

	var input *entity.MidtransNotificationRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	err := h.paymentDetailsService.HandleNotificationPaymentDetails(input)
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
