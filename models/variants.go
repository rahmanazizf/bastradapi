package models

import "time"

type Variant struct {
	ID          int       `gorm:"primary_key" json:"id"`
	UUID        int       `gorm:"not null" json:"uuid"`
	VariantName string    `gorm:"not null" json:"variant_name"`
	Quantity    int       `gorm:"not null" json:"qty"`
	ProductID   int       `gorm:"not null" json:"product_id"`
	CreatedAt   time.Time `gorm:"autoCreatedTime" json:"-"`
	UpdatedAt   time.Time `gorm:"autoUpdatedTime" json:"-"`
}
