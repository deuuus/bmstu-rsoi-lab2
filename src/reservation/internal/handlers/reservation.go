package handlers

import (
	"github.com/deuuus/bmsru-rsoi-lab2/src/reservation/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetListOfHotels(c *gin.Context) {
	hotels, err := h.services.GetListOfHotels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, hotels)
}

func (h *Handler) GetHotelInfo(c *gin.Context) {
	hotelUid := c.Param("hotelUid")
	if hotelUid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Param hotelUid is not presented."})
		return
	}

	hotelCheckInfo, err := h.services.GetHotelInfo(hotelUid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, hotelCheckInfo)
}

func (h *Handler) CancelReservation(c *gin.Context) {
	username := c.GetHeader("X-User-Name")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Header X-User-Name is not presented."})
		return
	}

	reservationUid := c.Param("reservationUid")
	if reservationUid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Param ReservationUid is not presented."})
		return
	}

	reservationUpdateInfo, err := h.services.CancelReservation(username, reservationUid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reservationUpdateInfo)
}

func (h *Handler) CreateReservation(c *gin.Context) {
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

	reservedInfo, err := h.services.CreateNewReservation(username, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reservedInfo)
}

func (h *Handler) GetReservationsShortInfo(c *gin.Context) {
	username := c.GetHeader("X-User-Name")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Header X-User-Name is not presented."})
		return
	}

	reservationShortInfo, err := h.services.GetListOfReservationShortInfo(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reservationShortInfo)
}

func (h *Handler) GetReservationShortInfo(c *gin.Context) {
	username := c.GetHeader("X-User-Name")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Header X-User-Name is not presented."})
		return
	}

	reservationUid := c.Param("reservationUid")
	if reservationUid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Param ReservationUid is not presented."})
		return
	}

	reservation, err := h.services.GetReservationShortInfo(username, reservationUid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reservation)
}
