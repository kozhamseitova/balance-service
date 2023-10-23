package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kozhamseitova/balance-service/api"
	"github.com/kozhamseitova/balance-service/internal/models"
)

// @Summary Deposit funds into the user's account
// @Description Deposit a specified amount of funds into the user's account.
// @ID depositFunds
// @Produce json
// @Param body body models.BalanceInput true "Balance Input"
// @Success 200 {object} api.Ok
// @Failure 400 {object} api.Error
// @Router /balance/credit [post]
func(h *Handler) deposit(c *fiber.Ctx) error {
	var req models.BalanceInput 

	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&api.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	err = h.service.DepositFunds(c.UserContext(), req.Id, req.Amount)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&api.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(&api.Ok{
		Code:    http.StatusOK,
		Message: "succes",
	})
} 

// @Summary Get the balance of a user by user ID
// @Description Retrieve the balance of a user by their user ID.
// @ID getBalanceByUserID
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} api.Ok
// @Failure 400 {object} api.Error
// @Failure 500 {object} api.Error
// @Router /balance/{id} [get]
func(h *Handler) getBalance(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&api.Error{
			Code:    http.StatusBadRequest,
			Message: "invalid param",
		})
	}

	result, err := h.service.GetBalanceByUserID(c.UserContext(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&api.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(&api.Ok{
		Code:    http.StatusOK,
		Message: "succes",
		Data: map[string]int{
			"balance": result,
		},
	})
}

// @Summary Reserve funds for a service
// @Description Reserve a specified amount of funds for a service.
// @ID reserveFunds
// @Produce json
// @Param body body models.ReserveInput true "Reserve Input"
// @Success 200 {object} api.Ok
// @Failure 400 {object} api.Error
// @Failure 500 {object} api.Error
// @Router /balance/reserve [post]
func(h *Handler) reserveFunds(c *fiber.Ctx) error {
	var req models.ReserveInput 

	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&api.Error{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
		})
	}

	err = h.service.ReserveFunds(c.UserContext(), req.UserID, req.ServiceID, req.OrderID, req.Amount)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&api.Error{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
		})
	}

	return c.Status(http.StatusOK).JSON(&api.Ok{
		Code:    http.StatusOK,
		Message: "succes",
	})
}

// @Summary Recognize revenue for a service
// @Description Recognize revenue for a service and order.
// @ID recognizeRevenue
// @Produce json
// @Param body body models.ReserveInput true "Reserve Input"
// @Success 200 {object} api.Ok
// @Failure 400 {object} api.Error
// @Failure 500 {object} api.Error
// @Router /balance/recognize [post]
func(h *Handler) recognizeRevenue(c *fiber.Ctx) error {
	var req models.ReserveInput 

	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&api.Error{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
		})
	}

	err = h.service.RecognizeRevenue(c.UserContext(), req.UserID, req.ServiceID, req.OrderID, req.Amount)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&api.Error{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
		})
	}

	return c.Status(http.StatusOK).JSON(&api.Ok{
		Code:    http.StatusOK,
		Message: "succes",
	})
}

// @Summary Get a report of revenue recognition
// @Description Get a report of revenue recognition for all users and services.
// @ID getRevenueReport
// @Produce json
// @Success 200 {object} api.Ok
// @Failure 500 {object} api.Error
// @Router /report [get]
func(h *Handler) getReport(c *fiber.Ctx) error {
	var result []*models.Report

	result, err := h.service.GetReport(c.UserContext())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&api.Error{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
		})
	}

	return c.Status(http.StatusOK).JSON(&api.Ok{
		Code:    http.StatusOK,
		Message: "succes",
		Data: result,
	})
}