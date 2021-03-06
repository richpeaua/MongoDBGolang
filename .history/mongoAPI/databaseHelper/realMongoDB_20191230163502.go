package databasehelper

import (
	clienthelper "MongoDBGolang/mongoAPI/clientHelper"
	collectionhelper "MongoDBGolang/mongoAPI/collectionHelper"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseHelper struct {
}

func (dd DatabaseHelper) NewClient(uri string) (clienthelper.IClientHelper, error) {
	var clientHelper clienthelper.RealClientHelper = clienthelper.RealClientHelper{}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return clienthelper.RealClientHelper{}, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return clienthelper.RealClientHelper{}, err
	}

	fmt.Println("Connected to MongoDB!")

	clientHelper.Client = client

	return client, nil
}

func (dd DatabaseHelper) Collection(name string) (collectionhelper.ICollectionHelper, error) {
	return &collectionhelper.DummyCollectionHelper{}, nil
}
