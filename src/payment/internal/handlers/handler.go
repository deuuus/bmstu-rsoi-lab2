package handlers

import (
	"github.com/deuuus/bmsru-rsoi-lab2/src/payment/internal/service"
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
		api.GET("/reservations", h.GetPaymentInfo)
		api.POST("/reservations", h.CreatePayment)
		api.POST("/reservations/:paymentUid", h.CancelPayment)
	}

	return router
}
