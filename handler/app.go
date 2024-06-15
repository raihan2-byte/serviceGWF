package handler

import (
	"log"
	"os"
	"payment-gwf/auth"
	"payment-gwf/database"
	"payment-gwf/middleware"
	"payment-gwf/repository"
	"payment-gwf/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func StartApp() {

	db, err := database.InitDb()
	if err != nil {
		log.Fatal("Eror Db Connection")
	}

	router := gin.Default()

	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file:", err)
	}
	//load config
	ServerKeys := os.Getenv("SERVER_KEY")
	if ServerKeys == "" {
		log.Fatal("Server key not found in environment variables")
	}
	gateway, err := service.NewMidtransGateway(&service.Config{
		ServerKey: ServerKeys,
	})

	if err != nil {
		log.Fatal("Failed to initialize Midtrans gateway: ", err)
	}

	// user
	userRepository := repository.NewRepositoryUser(db)
	userService := service.NewService(userRepository)
	authService := auth.NewService()
	userHandler := NewUserHandler(userService, authService)

	if err != nil {
		panic(err)
	}

	user := router.Group("/api/user")
	user.POST("/register", userHandler.RegisterUser)
	user.POST("/login", userHandler.Login)
	user.DELETE("/:slug", userHandler.DeletedUser)
	user.PUT("/:slug", userHandler.UpdateUser)

	// product
	productRepository := repository.NewRepositoryProduct(db)
	productService := service.NewServiceProduct(productRepository)
	productHandler := NewProductHandler(productService)

	api := router.Group("/api/products")
	api.POST("/", productHandler.CreateProduct)
	api.GET("/:id", productHandler.GetProduct)
	api.GET("/", productHandler.GetAllProduct)
	api.DELETE("/:id", productHandler.DeleteProduct)
	api.PUT("/:id", productHandler.UpdateProduct)

	cartRepository := repository.NewRepositoryCart(db)
	cartService := service.NewServiceCart(cartRepository, productRepository, userRepository)
	cartHandler := NewCartHandler(cartService, authService)

	api2 := router.Group("/api/user/cart")
	api2.POST("/:product_id", middleware.AuthMiddleware(authService, userService), cartHandler.AddToCart)
	api2.GET("/", middleware.AuthMiddleware(authService, userService), cartHandler.GetCartByUserID)
	api2.PUT("/:cart_id", middleware.AuthMiddleware(authService, userService), cartHandler.UpdatedCartByID)
	api2.DELETE("/delete/:id", middleware.AuthMiddleware(authService, userService), cartHandler.DeletedCartByID)

	addressRepository := repository.NewRepositoryAddress(db)
	addressService := service.NewServiceAddress(addressRepository, userRepository)
	addressHandler := NewAddressHandler(addressService, authService)

	api3 := router.Group("/api/user/address")
	api3.POST("/", middleware.AuthMiddleware(authService, userService), addressHandler.CreateAddress)
	//admin
	api3.GET("/", middleware.AuthRole(authService, userService), addressHandler.GetAllAddress)
	//admin
	api3.DELETE("/delete/:id", middleware.AuthRole(authService, userService), middleware.AuthMiddleware(authService, userService), addressHandler.DeleteAddress)

	rajaOngkirRepository := repository.NewRepositoryRajaOngkir(db)
	serviceRajaOngkir := service.NewServiceRajaOngkir(rajaOngkirRepository, addressRepository, userRepository)

	paymentRepository := repository.NewRepositoryPayment(db)

	orderRepository := repository.NewRepositoryOrder(db)
	orderService := service.NewServiceOrder(orderRepository, cartRepository, productRepository, userRepository, rajaOngkirRepository, paymentRepository, serviceRajaOngkir)
	orderHandler := NewOrderHandler(orderService, authService)

	// paymentService := service.NewServicePayment(paymentRepository, userRepository, orderRepository, )
	paymentService := service.NewServicePayment(paymentRepository, userRepository, orderRepository, gateway)
	paymentHandler := NewPaymentHandler(paymentService, authService)

	api4 := router.Group("/api/user/order")
	api4.POST("/", middleware.AuthMiddleware(authService, userService), orderHandler.CreateOrder)
	api4.GET("/", middleware.AuthMiddleware(authService, userService), orderHandler.GetOrderHistoryByUserID)
	api4.GET("/all-order", orderHandler.GetAllOrderHistory)

	apiPayment := router.Group("/api/user/payment")
	apiPayment.GET("/", middleware.AuthMiddleware(authService, userService), paymentHandler.GetAllPaymentByUserID)
	apiPayment.GET("/all-payment", paymentHandler.GetAllPayment)
	apiPayment.DELETE("/", paymentHandler.DeletePayment)
	apiPayment.POST("/:order_id", middleware.AuthMiddleware(authService, userService), paymentHandler.DoPayment)

	makeDonationRepository := repository.NewRepositoryMakeDonation(db)
	makeDonationService := service.NewServiceMakeDonation(makeDonationRepository, userRepository)
	makeDonationHandler := NewMakeDonationHandler(makeDonationService, authService)

	api5 := router.Group("/api/user/make-donation")
	api5.POST("/", middleware.AuthMiddleware(authService, userService), makeDonationHandler.CreateDonation)
	//admin
	api5.GET("/", middleware.AuthRole(authService, userService), middleware.AuthMiddleware(authService, userService), makeDonationHandler.GetAllDonation)
	//admin
	api5.DELETE("/delete/:id", middleware.AuthMiddleware(authService, userService), makeDonationHandler.DeleteDonation)

	api5.GET("/my-donation", middleware.AuthMiddleware(authService, userService), makeDonationHandler.GetDonation)

	api6 := router.Group("api/raja-ongkir/")

	provinceHandler := NewProvinceHandler(serviceRajaOngkir, addressService)

	api6.GET("/provinces", provinceHandler.GetProvinces)
	api6.GET("/city/:id", provinceHandler.GetCityByProvinceID)
	api6.GET("/calculate-shipping-fee", provinceHandler.CalculateShippingFee)
	api6.POST("/apply-shipping", middleware.AuthMiddleware(authService, userService), provinceHandler.ApplyShipping)
	// api6.POST("/apply-shipping-user", middleware.AuthMiddleware(authService, userService), provinceHandler.CreateAddressUser)
	// Port

	statusEkspedisiRepository := repository.NewRepositoryStatusEkspedisi(db)
	statusEkspedisiService := service.NewServiceStatusEkspedisi(statusEkspedisiRepository, orderRepository, userRepository)
	statusEkspedisiHandler := NewStatusEkspedisiHandler(statusEkspedisiService)

	api7 := router.Group("api/status-ekspedisi/")

	api7.POST("/:order_id/:user_id", statusEkspedisiHandler.CreateStatusEkspedisi)
	api7.GET("/user", middleware.AuthMiddleware(authService, userService), statusEkspedisiHandler.GetAllStatusEkspedisiByUserId)
	api7.GET("/", statusEkspedisiHandler.GetAllStatusEkspedisi)
	api7.DELETE("/:id", statusEkspedisiHandler.DeleteStatusEkspedisi)
	api7.PUT("/:id", statusEkspedisiHandler.UpdateStatusEkspedisi)

	router.Run(":8080")
}
