package models

import (
	"basic-trade-api/helpers"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	ID        int       `gorm:"primary_key" json:"-"`
	UUID      string    `json:"-"`
	Name      string    `json:"name" form:"name"`
	Email     string    `json:"email" form:"email"`
	Password  string    `json:"password" form:"password"`
	Salt      string    `json:"salt"`
	CreatedAt time.Time `gorm:"autoCreatedTime" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdatedTime" json:"-"`
	Products  []Product `gorm:"constraint:OnDelete:CASCADE" json:"products"`
}

// hook(s)
func (a *Admin) BeforeCreate(tx *gorm.DB) (err error) {
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
