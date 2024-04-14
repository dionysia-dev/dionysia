package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/learn-video/dionysia/internal/model"
	"gorm.io/gorm"
)

type InputStore interface {
	CreateInput(context.Context, *model.Input) error
	GetInput(context.Context, uuid.UUID) (model.Input, error)
	DeleteInput(context.Context, uuid.UUID) error
}

type DBInputStore struct {
	DB *gorm.DB
}

func NewDBInputStore(db *gorm.DB) *DBInputStore {
	return &DBInputStore{
		DB: db,
	}
}

func (s *DBInputStore) CreateInput(ctx context.Context, input *model.Input) error {
	return s.DB.WithContext(ctx).Create(input).Error
}

func (s *DBInputStore) GetInput(ctx context.Context, id uuid.UUID) (model.Input, error) {
	var input model.Input
	err := s.DB.WithContext(ctx).Preload("AudioProfiles").Preload("VideoProfiles").First(&input, "id = ?", id).Error

	return input, err
}

func (s *DBInputStore) DeleteInput(ctx context.Context, id uuid.UUID) error {
	return s.DB.WithContext(ctx).Delete(&model.Input{}, "id = ?", id).Error
}
