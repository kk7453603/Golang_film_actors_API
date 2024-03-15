package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Login    string
	Password string
	Role     bool
}

type Actor struct {
	gorm.Model
	Name     string
	Sex      string
	Birthday time.Time
}

type Film struct {
	gorm.Model
	Name   string
	Actors []Actor `gorm:"many2many:film_actors;"`
	Date   time.Time
	Rating int `gorm:"check: rating < 10 AND rating >= 0"`
}
