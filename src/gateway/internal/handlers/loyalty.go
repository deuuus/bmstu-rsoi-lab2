package handlers

import (
	"github.com/deuuus/bmsru-rsoi-lab2/src/gateway/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (gwh *GateWayHandler) GetLoyaltyStatus(c *gin.Context) {
	username := c.GetHeader("X-User-Name")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Header X-User-Name is not presented."})
		return
	}

	loyaltyInfo, err := service.GetUserLoyaltyInfoController(gwh.Services.LoyaltyServiceAddress, username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, loyaltyInfo)
}
