package handler

import (
	"net/http"
	"payment-gwf/auth"
	"payment-gwf/entity"
	"payment-gwf/helper"
	"payment-gwf/input"
	"payment-gwf/service"

	"github.com/gin-gonic/gin"
)

type makeDonationHandler struct {
	makeDonationService service.ServiceMakeDonation
	authService         auth.Service
}

func NewMakeDonationHandler(makeDonationService service.ServiceMakeDonation, authService auth.Service) *makeDonationHandler {
	return &makeDonationHandler{makeDonationService, authService}
}

func (h *makeDonationHandler) CreateDonation(c *gin.Context) {
	var input input.MakeDonationInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(*entity.User)
	//ini inisiasi userID yang mana ingin mendapatkan id si user
	getUserId := currentUser.ID

	makeDonation, err := h.makeDonationService.CreateDonation(getUserId, input)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := (makeDonation)
	response := helper.APIresponse(http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)
}

func (h *makeDonationHandler) GetAllDonation(c *gin.Context) {
	products, err := h.makeDonationService.GetDonations()
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Eror to get product")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, (products))
	c.JSON(http.StatusOK, response)
}

func (h *makeDonationHandler) GetDonation(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(*entity.User)
	//ini inisiasi userID yang mana ingin mendapatkan id si user
	getUserId := currentUser.ID

	products, err := h.makeDonationService.GetDonation(getUserId)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Eror to get product")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, (products))
	c.JSON(http.StatusOK, response)
}

func (h *makeDonationHandler) DeleteDonation(c *gin.Context) {
	idString := c.Param("id")
	// id, _ := strconv.Atoi(idString)

	products, err := h.makeDonationService.DeleteDonation(idString)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Eror to get product")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, (products))
	c.JSON(http.StatusOK, response)
}
