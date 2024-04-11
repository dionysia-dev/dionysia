package api

import (
	"net/http"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/learn-video/streaming-platform/internal/model"
	"github.com/stretchr/testify/assert"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func TestHandleValidationError(t *testing.T) {
	validate = validator.New(validator.WithRequiredStructEnabled())

	testCases := []struct {
		input *model.Input
		code  StatusCode
		// I don't feel the need to validate the error message, at least of now
		// In the case that this change, the logic of the validation would have to change too
		e any
	}{
		{
			input: &model.Input{
				Name:   "alex",
				Format: "format",
			},
			code: StatusCode(0),
			e:    nil,
		},
		{
			input: &model.Input{
				Format: "format",
			},
			code: StatusCode(http.StatusBadRequest),
			e:    (*ErrorResponse)(nil),
		},
	}

	for _, tt := range testCases {
		err := validate.Struct(tt.input)
		code, response := handleValidationError(err)

		assert.Equal(t, tt.code, code)
		assert.IsType(t, tt.e, response)
	}
}
