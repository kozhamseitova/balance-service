package models

import "time"

type BalanceInput struct {
	Id int `json:"id"`
	Amount int `json:"amount"`
}

type ReserveInput struct {
	UserID int `json:"user_id" db:"user_id"`
	ServiceID int `json:"service_id" db:"service_id"`
	OrderID int `json:"order_id" db:"order_id"`
	Amount int `json:"amount" db:"amount"`
}

type Report struct {
	Id int `json:"id" db:"id"`
	UserID int `json:"user_id" db:"user_id"`
	ServiceID int `json:"service_id" db:"service_id"`
	OrderID int `json:"order_id" db:"order_id"`
	Amount int `json:"amount" db:"amount"`
	RecognitionDate time.Time `json:"recognition_date" db:"recognition_date"`
}


type BalanceResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    BalanceData `json:"data"`
}
type BalanceData struct {
    Balance int `json:"balance"`
}