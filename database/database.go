package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB() (*mongo.Client, context.Context) {
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	credential := options.Credential{
		Username:    "api",
		Password:    "testpass",
		PasswordSet: true,
	}
	clientOptions.SetAuth(credential)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client, ctx
}

type Schema struct {
	ID     string `bson:"_id"`
	Schema []byte `bson:"schema"`
}

func PutSchema(db *mongo.Collection, key string, schema []byte) error {

	input := Schema{
		ID:     key,
		Schema: schema,
	}

	_, err := db.InsertOne(context.TODO(), input)

	return err

}

func GetSchema(db *mongo.Collection, key string) ([]byte, error) {

	filter := bson.M{"_id": key}

	res := Schema{}
	err := db.FindOne(context.TODO(), filter).Decode(&res)

	if err != nil {
		return nil, err
	}

	return res.Schema, nil
}
