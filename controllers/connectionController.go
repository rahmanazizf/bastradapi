package controllers

import (
	"basic-trade-api/models"
)

var connectDB *models.Connection

// func NewConnection(db *gorm.DB) *models.Connection {
// 	return &models.Connection{
// 		DB: db,
// 	}
// }

// func EstablishConnection(conn *models.Connection) {
// 	connectDB = conn
// }
