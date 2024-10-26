package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"root/configs"
	greeter "root/protobuf/greeter"
)

type DB struct {
	client *mongo.Client
}

func ConnectDB() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI(configs.EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	return &DB{client: client}
}

// // colHelper возвращает коллекцию по имени
// func colHelper(db *DB, collectionName string) *mongo.Collection {
// 	return db.client.Database("projectMngt").Collection(collectionName)
// }

// // InsertMessage вставляет новое сообщение в коллекцию
// func (db *DB) InsertMessage(message *model.Message) error {
// 	collection := colHelper(db, "messages") // Используем colHelper для получения коллекции

// 	_, err := collection.InsertOne(context.Background(), bson.M{
// 		"id":   message.ID,
// 		"text": message.Text,
// 	})
// 	log.Println("вроде норм бд")
// 	return err
// }

func (db *DB) InsertMember(member *greeter.CreateMemberRequest) error {
	collection := db.client.Database("projectMngt").Collection("members") // Укажите имя вашей коллекции

	_, err := collection.InsertOne(context.Background(), bson.M{
		"name": member.Name,
		// Добавьте другие поля по необходимости
	})
	if err != nil {
		return err
	}
	return nil
}