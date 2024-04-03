package model

import "github.com/google/uuid"

type Input struct {
	ID     uuid.UUID
	Name   string `json:"name" validate:"required"`
	Format string `json:"format" validate:"required"`
}
