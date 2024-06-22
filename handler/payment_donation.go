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

	"github.com/gin-gonic/gin"
)

type paymentDonationHandler struct {
	paymentDonationService service.ServicePaymentDonation
	authService            auth.Service
}

func NewPaymentDonationHandler(paymentDonationService service.ServicePaymentDonation, authService auth.Service) *paymentDonationHandler {
	return &paymentDonationHandler{paymentDonationService, authService}
}

// will be called by user through payment endpoint
func (h *paymentDonationHandler) DoPaymentDonation(c *gin.Context) {
	makeDonationID := c.Param("make_donation_id")

	var req input.SubmitPaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	resp, err := h.paymentDonationService.DoPaymentDonation(req, makeDonationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, formatter.FormatSimplePaymentResponse(resp))
}

func (h *paymentDonationHandler) GetPaymentDonationNotification(c *gin.Context) {
	// orderID := c.Param("order_id")

	// params, err := strconv.Atoi(orderID)

	var input *entity.MidtransNotificationRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	err := h.paymentDonationService.HandleNotificationPaymentDonation(input)
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

func (h *paymentDonationHandler) GetAllPaymentByUserID(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(*entity.User)
	//ini inisiasi userID yang mana ingin mendapatkan id si user
	getUserId := currentUser.ID

	getOrder, err := h.paymentDonationService.GetAllDonationByUserID(getUserId)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, formatter.FormatterGetAllPaymentDonation(getOrder))
	c.JSON(http.StatusOK, response)
}

func (h *paymentDonationHandler) GetAllPayment(c *gin.Context) {

	getOrder, err := h.paymentDonationService.GetAllPaymentDonation()
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, formatter.FormatterGetAllPaymentDonation(getOrder))
	c.JSON(http.StatusOK, response)
}

func (h *paymentDonationHandler) DeletePayment(c *gin.Context) {
	param := c.Param("id")
	// paramPayment, _ := strconv.Atoi(param)

	_, err := h.paymentDonationService.DeletePaymentDonation(param)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, "psyment has succesfuly deleted")
	c.JSON(http.StatusOK, response)
}
