package model

import "github.com/google/uuid"

type Input struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name" validate:"required"`
	Format string    `json:"format" validate:"required"`
}
