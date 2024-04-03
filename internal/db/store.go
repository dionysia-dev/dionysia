package db

import "context"

type InputStore interface {
	CreateInput(context.Context, CreateInputParams) (Input, error)
}
