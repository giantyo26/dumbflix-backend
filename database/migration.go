package database

import (
	"dumbflix-api/models"
	"dumbflix-api/pkg/mysql"
	"fmt"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.User{}, &models.Film{}, &models.Category{}, &models.Transaction{}, &models.Episode{})
	if err != nil {
		fmt.Println(err)
		panic("Migration failed")
	}
	fmt.Println("Migration sucess")

}
