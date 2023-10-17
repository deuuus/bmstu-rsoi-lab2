package handlers

import (
	"github.com/deuuus/bmsru-rsoi-lab2/src/gateway/internal/models"
	"github.com/deuuus/bmsru-rsoi-lab2/src/gateway/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (gwh *GateWayHandler) GetUserInfo(c *gin.Context) {
	username := c.GetHeader("X-User-Name")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Header X-User-Name is not presented."})
		return
	}

	userInfo, err := service.GetUserInfoController(gwh.Services.ReservationServiceAddress, gwh.Services.PaymentServiceAddress, gwh.Services.LoyaltyServiceAddress, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userInfo)
}

func (gwh *GateWayHandler) GetListOfUserReservations(c *gin.Context) {
	username := c.GetHeader("X-User-Name")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Header X-User-Name is not presented."})
		return
	}

	reservationsInfo, err := service.GetListOfUserReservationsController(gwh.Services.ReservationServiceAddress, gwh.Services.PaymentServiceAddress, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reservationsInfo)
}

func (gwh *GateWayHandler) GetReservationInfo(c *gin.Context) {
	reservationUid := c.Param("reservationUid")
	if reservationUid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ReservationUid is not presented."})
		return
	}

	username := c.GetHeader("X-User-Name")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Header X-User-Name is not presented."})
		return
	}

	reservationInfo, err := service.GetReservationInfoController(gwh.Services.ReservationServiceAddress, gwh.Services.PaymentServiceAddress, reservationUid, username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reservationInfo)
}

func (gwh *GateWayHandler) GetListOfHotels(c *gin.Context) {
	pageStr := c.Query("page")
	if pageStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Query param <page> is not presented."})
		return
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid param <page>."})
		return
	}

	sizeStr := c.Query("size")
	if sizeStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Query param <size> is not presented."})
		return
	}
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid param <size>."})
		return
	}

	hotels, err := service.GetListOfHotelsController(gwh.Services.ReservationServiceAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	right := page * size
	if len(*hotels) < right {
		right = len(*hotels)
	}

	hotelsStripped := (*hotels)[(page-1)*size : right]
	hs := models.HotelsLimited{
		Page:          page,
		PageSize:      size,
		TotalElements: len(hotelsStripped),
		Items:         &hotelsStripped,
	}
	c.JSON(http.StatusOK, hs)
}

func (gwh *GateWayHandler) ReserveHotel(c *gin.Context) {
	username := c.GetHeader("X-User-Name")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Header X-User-Name is not presented."})
		return
	}

	var input models.ReserveInfo
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	reservedInfo, err := service.CreateReservationController(gwh.Services.ReservationServiceAddress, gwh.Services.PaymentServiceAddress, gwh.Services.LoyaltyServiceAddress, username, &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reservedInfo)
}

func (gwh *GateWayHandler) CancelReservation(c *gin.Context) {
	reservationUid := c.Param("reservationUid")
	if reservationUid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ReservationUid is not presented."})
		return
	}

	username := c.GetHeader("X-User-Name")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Header X-User-Name is not presented."})
		return
	}

	err := service.CancelReservationController(gwh.Services.ReservationServiceAddress, gwh.Services.PaymentServiceAddress, gwh.Services.LoyaltyServiceAddress, reservationUid, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Reservation successfully canceled."})
}
