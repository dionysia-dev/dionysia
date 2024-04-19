package db

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/learn-video/dionysia/internal/model"
	"gorm.io/gorm"
)

var (
	ErrNotFound = errors.New("record not found")
)

type InputStore interface {
	CreateInput(context.Context, *model.Input) error
	GetInput(context.Context, uuid.UUID) (model.Input, error)
	DeleteInput(context.Context, uuid.UUID) error
}

type InputStoreDB struct {
	DB *gorm.DB
}

func NewDBInputStore(db *gorm.DB) InputStore {
	return &InputStoreDB{
		DB: db,
	}
}

func (s *InputStoreDB) CreateInput(ctx context.Context, input *model.Input) error {
	return s.DB.WithContext(ctx).Create(input).Error
}

func (s *InputStoreDB) GetInput(ctx context.Context, id uuid.UUID) (model.Input, error) {
	var input model.Input
	err := s.DB.WithContext(ctx).Preload("AudioProfiles").Preload("VideoProfiles").First(&input, "id = ?", id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.Input{}, ErrNotFound
	}

	return input, err
}

func (s *InputStoreDB) DeleteInput(ctx context.Context, id uuid.UUID) error {
	return s.DB.WithContext(ctx).Delete(&model.Input{}, "id = ?", id).Error
}
