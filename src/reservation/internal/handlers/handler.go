package handlers

import (
	"github.com/deuuus/bmsru-rsoi-lab2/src/reservation/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
	return
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/manage/health", h.CheckHealth)

	api := router.Group("/api/v1")
	{
		api.GET("/hotels", h.GetListOfHotels)
		api.GET("/hotels/:hotelUid", h.GetHotelInfo)
		api.GET("/reservations", h.GetReservationsShortInfo)
		api.POST("/reservations", h.CreateReservation)
		api.GET("/reservations/:reservationUid", h.GetReservationShortInfo)
		api.POST("/reservations/:reservationUid", h.CancelReservation)
	}

	return router
}
