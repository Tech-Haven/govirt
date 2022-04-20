package models

import (
	"context"
	"govirt/configs"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Ovirt		OvirtUser
	Discord	DiscordUser
}

type OvirtUser struct {
	Username			string	`json:"username,omitempty"`
	Access_token	string	`json:"access_token,omitempty"`
}

type DiscordUser struct {
	Id				string	`json:"id,omitempty"`
	Username	string	`json:"username,omitempty"`
}

// Accepts a user model and returns Mongo update results
func SaveUser(model User, db *mongo.Client) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

	var userCollection *mongo.Collection = configs.GetCollection(db, "users")

	filter := bson.D{primitive.E{Key: "ovirt.username", Value: model.Ovirt.Username}}
	update := bson.D{primitive.E{Key: "$set", Value: model}}
	opts := options.Update().SetUpsert(true)

	result, err := userCollection.UpdateOne(ctx, filter, update, opts)
		if err != nil {
			return nil, err
	}

	return result, nil
}