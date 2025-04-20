package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var DB *mongo.Database

func ConnectDB() {
	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")

	connectionString := "mongodb://" + username + ":" + password + "@" + host + ":" + port

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal("Ошибка создания клиента MongoDB:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Ошибка подключения к MongoDB:", err)
	}

	DB = client.Database(os.Getenv("DB_NAME"))
	log.Println("Успешное подключение к MongoDB")
}
