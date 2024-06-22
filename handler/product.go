package handler

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"payment-gwf/formatter"
	"payment-gwf/helper"
	"payment-gwf/imagekits"
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

	var imagesKitURLs []string

	// Loop through all file input fields (e.g., "file1", "file2", etc.)
	for i := 1; ; i++ {
		fileKey := fmt.Sprintf("file%d", i)
		file, err := c.FormFile(fileKey)

		// If there are no more files to upload, break the loop
		if err == http.ErrMissingFile {
			break
		}

		if err != nil {
			fmt.Printf("Error when opening file %s: %v\n", fileKey, err)
			continue // Skip to the next file
		}

		src, err := file.Open()
		if err != nil {
			fmt.Printf("Error when opening file %s: %v\n", fileKey, err)
			continue
		}
		defer src.Close()

		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, src); err != nil {
			fmt.Printf("Error reading file %s: %v\n", fileKey, err)
			continue
		}

		img, err := imagekits.Base64toEncode(buf.Bytes())
		if err != nil {
			fmt.Printf("Error reading image %s: %v\n", fileKey, err)
			continue
		}

		fmt.Printf("Image base64 format %s: %v\n", fileKey, img)

		imageKitURL, err := imagekits.ImageKit(context.Background(), img)
		if err != nil {
			fmt.Printf("Error uploading image %s to ImageKit: %v\n", fileKey, err)
			continue
		}

		imagesKitURLs = append(imagesKitURLs, imageKitURL)
	}

	var input input.ProductInput

	err := c.ShouldBind(&input)
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

	for _, imageURL := range imagesKitURLs {
		// Create a new BeritaImage record for each image and associate it with the news item
		err := h.productService.CreateProductImage(newProduct.ID, imageURL)
		if err != nil {
			response := helper.APIresponse(http.StatusUnprocessableEntity, err)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
	}

	response := helper.APIresponse(http.StatusOK, formatter.FormatterProduct(newProduct))
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
	response := helper.APIresponse(http.StatusOK, formatter.FormatterGetProducts(products))
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
	response := helper.APIresponse(http.StatusOK, formatter.FormatterProduct(products))
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) DeleteProduct(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	err := h.productService.DeleteProduct(int(id))
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Eror to get product")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, "products has succesfuly deleted")
	c.JSON(http.StatusOK, response)
}
