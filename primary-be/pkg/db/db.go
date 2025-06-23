package db

import (
	"fmt"

	"github.com/bikaxh/vid-gen/primary-be/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() {

	dbUrl := utils.GoDotEnvVariable("DATABASE_URL")
	var err error
	Db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		panic("Error while Connecting to DB 🔴")
	}

	fmt.Println("Connected successfully 🟢")

}
