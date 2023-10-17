package models

type LoyaltyStatus struct {
	Username         string `json:"username"`
	ReservationCount int    `db:"reservation_count" json:"reservationCount"`
	Status           string `json:"status"`
	Discount         int    `json:"discount"`
}

type LoyaltyStatusShort struct {
	Status string `json:"status"`
}
