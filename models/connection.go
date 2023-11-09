package models

import "gorm.io/gorm"

type Connection struct {
	DB *gorm.DB
}

// fungsi-fungsi general
// IfExist mengecek apakah suatu record sudah tertulis di database
func (c *Connection) IfExist(model interface{}) interface{} {
	res := c.DB.First(&model)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return nil
	}
	return c
}

func (c *Connection) ReturnDynamic() interface{} {
	return c.DB
}

type ErrorReturn struct {
	Value error
}

func (e *ErrorReturn) ReturnDynamic() interface{} {
	return e.Value
}
