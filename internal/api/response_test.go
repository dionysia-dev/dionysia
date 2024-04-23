package api_test

import (
	"net/http"
	"testing"

	"github.com/dionysia-dev/dionysia/internal/api"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func TestHandleValidationError(t *testing.T) {
	validate = validator.New(validator.WithRequiredStructEnabled())

	testCases := []struct {
		name               string
		inputError         *api.InputData
		expectedStatusCode api.StatusCode
		expectedErrorResp  *api.ErrorResponse
	}{
		{
			name: "Test with bad request body",
			inputError: &api.InputData{
				Name: "just a name",
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedErrorResp: &api.ErrorResponse{
				Error: api.Error{
					Message: "Invalid request parameters",
					Details: []api.ErrorDetail{
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
			inputError: &api.InputData{
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
			actualStatusCode, actualErrorResp := api.HandleValidationError(err)

			assert.Equal(t, tc.expectedStatusCode, actualStatusCode, "Status code mismatch")
			assert.Equal(t, tc.expectedErrorResp, actualErrorResp, "Error response mismatch")
		})
	}
}
