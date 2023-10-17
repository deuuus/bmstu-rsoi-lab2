package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetLoyaltyStatus(c *gin.Context) {
	username := c.GetHeader("X-User-Name")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Header X-User-Name is not presented."})
		return
	}

	loyaltyInfo, err := h.services.GetLoyaltyStatus(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, loyaltyInfo)
}

func (h *Handler) UpdateReservationsCount(c *gin.Context) {
	username := c.GetHeader("X-User-Name")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Header X-User-Name is not presented."})
		return
	}
	acc := c.GetHeader("Acc")
	accInt, err := strconv.Atoi(acc)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	loyaltyInfo, err := h.services.UpdateReservationCount(username, accInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, loyaltyInfo)
}
