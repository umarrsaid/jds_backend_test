package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ResponseObject struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

type ResponsePaginationObject struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

type PaginationMeta struct {
	TotalData  int64 `json:"total_data"`
	TotalPages int64 `json:"total_pages"`
	Page       int64 `json:"page"`
}

type ResponseValidationObject struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

// Response defines a generic response
func Response(c *fiber.Ctx, code int, status string, result interface{}, message string) error {
	resp := ResponseObject{
		Code:    code,
		Status:  status,
		Message: message,
		Result:  result,
	}
	return c.Status(code).JSON(resp)
}

// ErrorResponse defines a generic error response
func ErrorResponseHandler(c *fiber.Ctx, code int, status string, message string) error {
	resp := ResponseObject{
		Code:    code,
		Status:  status,
		Message: message,
		Result:  nil,
	}
	return c.Status(code).JSON(resp)
}

func ResponsePagination(c *fiber.Ctx, code int, status string, result interface{}, message string) error {
	resp := ResponsePaginationObject{
		Code:    code,
		Status:  status,
		Message: message,
		Result:  result,
	}
	return c.Status(code).JSON(resp)
}

func ErrorResponseValidation(c *fiber.Ctx, message string, err error) error {
	var validationErrors []map[string]string
	for _, e := range err.(validator.ValidationErrors) {
		errorMessage := fmt.Sprintf("Error on the field %s, condition: %s",
			e.Field(), e.ActualTag())
		errorMap := map[string]string{
			"key":     e.Field(),
			"message": errorMessage,
		}
		validationErrors = append(validationErrors, errorMap)
	}

	resp := ResponseValidationObject{
		Code:    400,
		Message: message,
		Data:    nil,
		Error:   validationErrors,
	}
	return c.Status(400).JSON(resp)
}

// Global error response
func GlobalErrorResponse(c *fiber.Ctx, err error) error {
	if err != nil {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}
		return c.Status(code).JSON(fiber.Map{"error": err.Error()})
	}
	return nil
}
