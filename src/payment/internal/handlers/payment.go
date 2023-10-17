package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreatePayment(c *gin.Context) {
	var price int
	if err := c.BindJSON(&price); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	paymentFullInfo, err := h.services.CreatePayment(price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, paymentFullInfo)
}

func (h *Handler) GetPaymentInfo(c *gin.Context) {
	paymentUid := c.GetHeader("paymentUid")
	if paymentUid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Header paymentUid is not presented."})
		return
	}

	paymentInfo, err := h.services.GetPaymentInfo(paymentUid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, paymentInfo)
}

func (h *Handler) CancelPayment(c *gin.Context) {
	paymentUid := c.Param("paymentUid")
	if paymentUid == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "Header paymentUid is not presented."})
		return
	}

	err := h.services.CancelPayment(paymentUid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}
