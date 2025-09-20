package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ctx = context.Background()

var DB *gorm.DB

var RDB *redis.Client

func ConnectDB() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	dsn := os.Getenv("DB_DSN")
	log.Println("DB_DSN:", os.Getenv("DB_DSN"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	DB = db
	fmt.Println("Database Connected")
}

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // adjust if using Docker
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	// test connection
	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("❌ Could not connect to Redis: %v", err)
	}
	log.Println("✅ Connected to Redis")
}

func SetCache(key string, value string, ttl time.Duration) {
	RDB.Set(ctx, key, value, ttl)
}

func GetCache(key string) (string, error) {
	return RDB.Get(ctx, key).Result()
}
