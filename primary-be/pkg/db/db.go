package db

import (
	"context"
	"fmt"

	"github.com/bikaxh/vid-gen/primary-be/pkg/utils"
	"github.com/redis/go-redis/v9"
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

var RedisClient *redis.Client

func ConnectRedisClient() {

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     utils.GoDotEnvVariable("REDIS_URL"),
		Password: "",
		DB:       0,
	})

	ctx := context.Background()
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}

	fmt.Println("🟢 Connected to Redis")

}
