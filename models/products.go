package models

import (
	"log"
	"mime/multipart"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          int       `gorm:"primary_key" json:"id"`
	UUID        string    `gorm:"not null" json:"uuid"`
	ProductName string    `gorm:"not null" json:"product_name" validate:"required"`
	ImageURL    string    `gorm:"not null" json:"image_url" validate:"required"`
	AdminID     int       `gorm:"not null" json:"admin_id"`
	CreatedAt   time.Time `gorm:"autoCreatedTime" json:"-"`
	UpdatedAt   time.Time `gorm:"autoUpdatedTime" json:"-"`
	Variants    []Variant `gorm:"constraint:OnDelete:CASCADE" json:"variants"`
}

type ProductImage struct {
	ProductName string                `form:"product_name" binding:"required"`
	ImageFile   *multipart.FileHeader `form:"image_file" binding:"required"`
	AdminID     int                   `form:"admin_id"`
	Variants    []Variant             `form:"variants"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(p)
	if err != nil {
		log.Println(err.Error())
		return
	}
	p.UUID = uuid.New().String()
	return nil
}

// layer bawah untuk abstraksi products controller
func (c *Connection) CreateProduct() {

}
