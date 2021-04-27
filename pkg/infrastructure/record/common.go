package record

import (
	"gorm.io/gorm"
	"time"

	"github.com/google/uuid"
)

type (
	// https://qiita.com/gold-kou/items/45a95d61d253184b0f33
	CommonRecord struct {
		ID        string `gorm:"primary_key;"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	CommonRecordSoftDelete struct {
		ID        string `gorm:"primary_key;"`
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt
	}
)

// https://gorm.io/docs/create.html#Create-Hooks
func (r *CommonRecord) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.New().String()
	return
}

func (r *CommonRecordSoftDelete) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.New().String()
	return
}
