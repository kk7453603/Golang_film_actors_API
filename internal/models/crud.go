package models

import (
	"gorm.io/gorm"
)

func Create_user(db *gorm.DB, login string, password string, role bool) error {
	if tx := db.Create(&User{Login: login, Password: password, Role: role}); tx.Error != nil {
		return tx.Error
	}
	return nil
}
