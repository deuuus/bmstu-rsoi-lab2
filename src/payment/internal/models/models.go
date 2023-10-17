package models

type PaymentInfo struct {
	Status string `json:"status"`
	Price  int    `json:"price"`
}

type PaymentFullInfo struct {
	PaymentUid string `json:"paymentUid"`
	Status     string `json:"status"`
	Price      int    `json:"price"`
}
