package migration

import (
	"log"

	"github.com/devianwahyu/farmigo/database"
	"github.com/devianwahyu/farmigo/model/entity"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.Role{}, &entity.User{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("Successfully migrated database")
}
