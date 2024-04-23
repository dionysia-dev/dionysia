package api

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

type SuccessResponse struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}

type ErrorDetail struct {
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

type Error struct {
	Message string        `json:"message"`
	Details []ErrorDetail `json:"details,omitempty"`
}

type ErrorResponse struct {
	Error `json:"error"`
}

type StatusCode int

// HandleValidationError validate and handle errors of structs that has `validator/v10` rules
func HandleValidationError(err error) (StatusCode, *ErrorResponse) {
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return 0, nil
	}

	details := make([]ErrorDetail, 0, len(validationErrors))

	for _, fieldErr := range err.(validator.ValidationErrors) {
		details = append(details, ErrorDetail{
			Reason:  fieldErr.Field(),
			Message: fieldErr.ActualTag(),
		})
	}

	return http.StatusBadRequest, &ErrorResponse{
		Error: Error{
			Message: "Invalid request parameters",
			Details: details,
		},
	}
}
