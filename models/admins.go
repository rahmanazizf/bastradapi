package models

import (
	"basic-trade-api/helpers"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	ID        int       `gorm:"primary_key" json:"id"`
	UUID      string    `json:"-"`
	Name      string    `json:"name" form:"name" validate:"required"`
	Email     string    `json:"email" form:"email" validate:"required,email"`
	Password  string    `json:"password" form:"password" validate:"required"`
	Salt      string    `json:"-"`
	CreatedAt time.Time `gorm:"autoCreatedTime" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdatedTime" json:"-"`
	Products  []Product `gorm:"constraint:OnDelete:CASCADE" json:"-"`
}

// hook(s)
func (a *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	// validate := validator.New(validator.WithRequiredStructEnabled())
	validate := validator.New()
	err = validate.Struct(a)
	if err != nil {
		log.Println(err.Error())
		return
	}
	a.UUID = uuid.New().String()
	a.Salt = helpers.SaltGenerator()
	a.Password = helpers.HashPwd(a.Password, a.Salt)

	return nil
}

// layer bawah abstraksi untuk admin controller
func (c *Connection) CreateAdmin(admin *Admin) {
	res := c.IfExist(&admin)
	// _, isError := res.(error)
	_, isConnection := res.(Connection)
	if isConnection {
		// row := res.DB.Create(&admin)
	}
}
func (c *Connection) GetAdmin() {

}
