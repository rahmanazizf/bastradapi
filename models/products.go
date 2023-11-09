package models

import "time"

type Product struct {
	ID          int       `gorm:"primary_key" json:"id"`
	UUID        string    `gorm:"not null" json:"uuid"`
	ProductName string    `gorm:"not null" json:"product_name"`
	ImageURL    string    `gorm:"not null" json:"image_url"`
	AdminID     int       `gorm:"not null" json:"admin_id"`
	CreatedAt   time.Time `gorm:"autoCreatedTime" json:"-"`
	UpdatedAt   time.Time `gorm:"autoUpdatedTime" json:"-"`
	Variants    []Variant `gorm:"constraint:OnDelete:CASCADE" json:"variants"`
}

// layer bawah untuk abstraksi products controller
func (c *Connection) CreateProduct() {

}
