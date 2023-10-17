package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ServicesAddresses struct {
	ReservationServiceAddress string
	PaymentServiceAddress     string
	LoyaltyServiceAddress     string
}

type GateWayHandler struct {
	Services ServicesAddresses
}

func NewGateWayHandler(services *ServicesAddresses) *GateWayHandler {
	return &GateWayHandler{Services: *services}
}

func (gwh *GateWayHandler) CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
	return
}

func InitRoutes() *gin.Engine {
	router := gin.Default()

	services := ServicesAddresses{
		ReservationServiceAddress: "http://reservation-service:8070",
		PaymentServiceAddress:     "http://payment-service:8060",
		LoyaltyServiceAddress:     "http://loyalty-service:8050",
	}

	gwh := NewGateWayHandler(&services)

	router.GET("/manage/health", gwh.CheckHealth)

	api := router.Group("/api/v1")
	{
		api.GET("/loyalty", gwh.GetLoyaltyStatus)
		api.GET("/hotels", gwh.GetListOfHotels)
		api.GET("/reservations", gwh.GetListOfUserReservations)
		api.GET("/reservations/:reservationUid", gwh.GetReservationInfo)
		api.GET("/me", gwh.GetUserInfo)
		api.POST("/reservations", gwh.ReserveHotel)
		api.DELETE("/reservations/:reservationUid", gwh.CancelReservation)
	}

	return router
}
