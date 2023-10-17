package models

type ReserveInfo struct {
	PaymentUid string `json:"paymentUid"`
	HotelUid   string `json:"hotelUid"`
	StartDate  string `json:"startDate"`
	EndDate    string `json:"endDate"`
}

type HotelFullInfo struct {
	HotelUid string `db:"hotel_uid" json:"hotelUid"`
	Name     string `json:"name"`
	Country  string `json:"country"`
	City     string `json:"city"`
	Address  string `json:"address"`
	Stars    int    `json:"stars"`
	Price    int    `json:"price"`
}

type ReservationSingleInfo struct {
	PaymentUid string `db:"payment_uid" json:"paymentUid"`
	HotelId    string `db:"hotel_id" json:"hotelUid"`
	Status     string `json:"status"`
	StartDate  string `db:"start_date" json:"startDate"`
	EndDate    string `db:"end_data" json:"endDate"`
}

type ReservationFullInfo struct {
	ReservationUid string `db:"reservation_uid" json:"reservationUid"`
	PaymentUid     string `db:"payment_uid" json:"paymentUid"`
	HotelId        string `db:"hotel_id" json:"hotelId"`
	Status         string `json:"status"`
	StartDate      string `db:"start_date" json:"startDate"`
	EndDate        string `db:"end_data" json:"endDate"`
}

type ReservationUpdateInfo struct {
	PaymentUid string `json:"paymentUid"`
}

type ReservationShortInfo struct {
	ReservationUid string     `json:"reservationUid"`
	PaymentUid     string     `json:"paymentUid"`
	Hotel          *HotelInfo `json:"hotel"`
	StartDate      string     `json:"startDate"`
	EndDate        string     `json:"endDate"`
	Status         string     `json:"status"`
}

type ReservedShortInfo struct {
	ReservationUid string `json:"reservationUid"`
	//HotelUid       string `json:"hotelUid"`
	//StartDate      string `json:"startDate"`
	//EndDate        string `json:"endDate"`
	//Discount       int    `json:"discount"`
	Status string `json:"status"`
}

type HotelInfo struct {
	HotelUid    string `json:"hotelUid"`
	Name        string `json:"name"`
	FullAddress string `json:"fullAddress"`
	Stars       int    `json:"stars"`
}

type HotelShortInfo struct {
	HotelUid string `db:"hotel_uid" json:"hotelUid"`
	Name     string `json:"name"`
	Country  string `json:"country"`
	City     string `json:"city"`
	Address  string `json:"address"`
	Stars    int    `json:"stars"`
}

type HotelCheckInfo struct {
	Price int `json:"price"`
}
