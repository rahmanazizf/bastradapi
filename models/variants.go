package models

import (
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TOOD: tambahkan validasi
type Variant struct {
	ID          int       `gorm:"primary_key" json:"id"`
	UUID        string    `gorm:"not null" json:"uuid"`
	VariantName string    `gorm:"not null" json:"variant_name" validate:"required"`
	Quantity    int       `gorm:"not null" json:"qty" validate:"required,gte=0"`
	ProductID   int       `gorm:"not null" json:"product_id" validate:"required"`
	CreatedAt   time.Time `gorm:"autoCreatedTime" json:"-"`
	UpdatedAt   time.Time `gorm:"autoUpdatedTime" json:"-"`
}

func (v *Variant) BeforeCreate(tx *gorm.DB) (err error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(v)
	if err != nil {
		log.Println(err.Error())
		return
	}
	v.UUID = uuid.New().String()
	return nil
}
