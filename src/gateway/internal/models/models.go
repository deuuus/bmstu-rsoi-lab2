package models

type UserInfo struct {
	Reservations *[]ReservationInfo `json:"reservations"`
	Loyalty      *LoyaltyShortInfo  `json:"loyalty"`
}

type ReservationInfo struct {
	ReservationUid string       `json:"reservationUid"`
	Hotel          *HotelInfo   `json:"hotel"`
	StartDate      string       `json:"startDate"`
	EndDate        string       `json:"endDate"`
	Status         string       `json:"status"`
	Payment        *PaymentInfo `json:"payment"`
}

type ReservationShortInfo struct {
	ReservationUid string     `json:"reservationUid"`
	PaymentUid     string     `json:"paymentUid"`
	Hotel          *HotelInfo `json:"hotel"`
	StartDate      string     `json:"startDate"`
	EndDate        string     `json:"endDate"`
	Status         string     `json:"status"`
}

type ReservationUpdateInfo struct {
	PaymentUid string `json:"paymentUid"`
}

type PaymentInfo struct {
	Status string `json:"status"`
	Price  int    `json:"price"`
}

type PaymentFullInfo struct {
	PaymentUid string `json:"paymentUid"`
	Status     string `json:"status"`
	Price      int    `json:"price"`
}

type ReservedShortInfo struct {
	ReservationUid string `json:"reservationUid"`
	//HotelUid       string `json:"hotelUid"`
	//StartDate      string `json:"startDate"`
	//EndDate        string `json:"endDate"`
	//Discount       int    `json:"discount"`
	Status string `json:"status"`
}

type ReservedInfo struct {
	ReservationUid string       `json:"reservationUid"`
	HotelUid       string       `json:"hotelUid"`
	StartDate      string       `json:"startDate"`
	EndDate        string       `json:"endDate"`
	Discount       int          `json:"discount"`
	Status         string       `json:"status"`
	Payment        *PaymentInfo `json:"payment"`
}

type HotelInfo struct {
	HotelUid    string `json:"hotelUid"`
	Name        string `json:"name"`
	FullAddress string `json:"fullAddress"`
	Stars       int    `json:"stars"`
}

type HotelCheckInfo struct {
	Price int `json:"price"`
}

type HotelFullInfo struct {
	HotelUid string `json:"hotelUid"`
	Name     string `json:"name"`
	Country  string `json:"country"`
	City     string `json:"city"`
	Address  string `json:"address"`
	Stars    int    `json:"stars"`
	Price    int    `json:"price"`
}

type HotelsLimited struct {
	Page          int              `json:"page"`
	PageSize      int              `json:"pageSize"`
	TotalElements int              `json:"totalElements"`
	Items         *[]HotelFullInfo `json:"items"`
}

type LoyaltyShortInfo struct {
	Status   string `json:"status"`
	Discount int    `json:"discount"`
}

type LoyaltyStatus struct {
	Status string `json:"status"`
}

type LoyaltyInfo struct {
	Username         string `json:"username"`
	Status           string `json:"status"`
	Discount         int    `json:"discount"`
	ReservationCount int    `json:"reservationCount"`
}

type ReserveInfo struct {
	HotelUid  string `json:"hotelUid"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type ReserveFullInfo struct {
	PaymentUid string `json:"paymentUid"`
	HotelUid   string `json:"hotelUid"`
	StartDate  string `json:"startDate"`
	EndDate    string `json:"endDate"`
}
