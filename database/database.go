package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Schema struct {
	ID     string `bson:"_id"`
	schema []byte `bson:"schema"`
}

func PutSchema(db *mongo.Collection, key string, schema []byte) error {

	input := Schema{
		ID:     key,
		schema: schema,
	}

	_, err := db.InsertOne(context.TODO(), input)
	if err != nil {
		return err
	}
	// could log insertion?

	return nil

}

func GetSchema(db *mongo.Collection, key string) ([]byte, error) {

	filter := bson.M{"_id": key}

	res := Schema{}

	err := db.FindOne(context.TODO(), filter).Decode(&res)

	if err != nil {
		//handle err
		return nil, err
	}

	return res.schema, nil
}
