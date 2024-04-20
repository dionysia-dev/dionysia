package db

import (
	"github.com/dionysia-dev/dionysia/internal/model"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Input{}, &model.AudioProfile{}, &model.VideoProfile{}) //nolint:errcheck // skip error check
}
