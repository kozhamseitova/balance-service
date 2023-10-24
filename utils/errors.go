package utils

import "errors"

var (
	ErrInternalError = errors.New("internal server error")
	ErrNotFound = errors.New("user not found error")
	ErrReservationNotFound = errors.New("reservation not found")
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrInvalidAmount = errors.New("invalid amount")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrDuplicateKey = errors.New("service already ordered by this user")
) 