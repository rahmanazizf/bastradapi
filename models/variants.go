package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TOOD: tambahkan validasi
type Variant struct {
	ID          int       `gorm:"primary_key" json:"id"`
	UUID        string    `gorm:"not null" json:"uuid"`
	VariantName string    `gorm:"not null" json:"variant_name"`
	Quantity    int       `gorm:"not null" json:"qty"`
	ProductID   int       `gorm:"not null" json:"product_id"`
	CreatedAt   time.Time `gorm:"autoCreatedTime" json:"-"`
	UpdatedAt   time.Time `gorm:"autoUpdatedTime" json:"-"`
}

func (v *Variant) BeforeCreate(tx *gorm.DB) (err error) {
	v.UUID = uuid.New().String()
	return nil
}
