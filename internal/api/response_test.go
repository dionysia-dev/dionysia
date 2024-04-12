package api

import (
	"net/http"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/learn-video/dionysia/internal/model"
	"github.com/stretchr/testify/assert"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func TestHandleValidationError(t *testing.T) {
	validate = validator.New(validator.WithRequiredStructEnabled())

	// Test cases
	testCases := []struct {
		name               string
		inputError         *model.Input
		expectedStatusCode StatusCode
		expectedErrorResp  *ErrorResponse
	}{
		{
			name: "Test with bad request body",
			inputError: &model.Input{
				Name: "just a name",
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedErrorResp: &ErrorResponse{
				Error: Error{
					Message: "Invalid request parameters",
					Details: []ErrorDetail{
						{
							Reason:  "Format",
							Message: "required",
						},
					},
				},
			},
		},
		{
			name: "Test with valid request body",
			inputError: &model.Input{
				Name:   "beautiful name",
				Format: "nice format",
			},
			expectedStatusCode: 0,
			expectedErrorResp:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validate.Struct(tc.inputError)
			actualStatusCode, actualErrorResp := handleValidationError(err)

			assert.Equal(t, tc.expectedStatusCode, actualStatusCode, "Status code mismatch")
			assert.Equal(t, tc.expectedErrorResp, actualErrorResp, "Error response mismatch")
		})
	}
}
