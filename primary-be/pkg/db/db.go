package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() {

	dbUrl := "postgresql://neondb_owner:npg_xlomHr18DVgt@ep-tiny-boat-a8btcdhj-pooler.eastus2.azure.neon.tech/neondb?sslmode=require"
	var err error
	Db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		panic("Error while Connecting to DB 🔴")
	}

	fmt.Println("Connected successfully 🟢")

}
