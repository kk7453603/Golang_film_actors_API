package main

import (
	"log"

	"github.com/kinofilm_go_vk_test/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{}, &models.Actor{}, &models.Film{})
	user := models.User{Login: "admin", Password: "123", Role: true}
	res := db.Create(&user)
	log.Println(res.Error)
}
