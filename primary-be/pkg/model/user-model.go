package model

import (
	"errors"
	"fmt"

	"github.com/bikaxh/vid-gen/primary-be/pkg/db"
	"github.com/bikaxh/vid-gen/primary-be/pkg/utils"
	"github.com/google/uuid"
	
	"time"
)

type User struct {
	ID       string `gorm:"type:uuid;primaryKey" json:"id"`
	Email    string `gorm:"primaryKey" json:"email" validate:"required"`
	UserName string `json:"userName" validate:"required"`
	Password string `json:"password" validate:"required"`
	Projects  []Project `gorm:"foreignKey:UserId"`
	CreatedAt time.Time
	UpdatedAt time.Time
	
}

func init() {
	db.Connect()
	//  db.Db.AutoMigrate(&User{}, &Project{}, &Scene{})
}

func (user *User) Save() (*User, error) {
	hashPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return nil, err
	}
	user.ID = uuid.New().String()
	user.Password = *hashPassword

	result := db.Db.Create(user)

	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}
	user.Password = ""
	return user, nil
}

func FindUserById(userId string) (*User, error) {
	var user User = User{
		ID: userId,
	}
	result := db.Db.First(&user, &user)

	user.Password = ""
	if result.Error != nil {
		return nil, errors.New("User not found")
	}

	return &user, nil
}

func FindUserByEmail(email string) (*User, error) {

	var existingUser User = User{
		Email: email,
	}
	result := db.Db.First(&existingUser, &existingUser)

	if result.Error != nil {
		return nil, errors.New("User not found")
	}

	return &existingUser, nil
}
