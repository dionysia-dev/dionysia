package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func handleValidationError(ctx *gin.Context, err error) {
	var details []ErrorDetail

	for _, fieldErr := range err.(validator.ValidationErrors) {
		details = append(details, ErrorDetail{
			Reason:  fieldErr.Field(),
			Message: fieldErr.ActualTag(),
		})
	}

	ctx.JSON(http.StatusBadRequest, ErrorResponse{
		Error: Error{
			Message: "Invalid request parameters",
			Details: details,
		},
	})
}
