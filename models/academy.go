package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Academy struct {
	ID              primitive.ObjectID `bson:"_id"`
	Name            string             `bson:"name"`
	Abbreviation    string             `bson:"abbreviation"`
	Website         string             `bson:"website"`
	Endpoint        string             `bson:"endpoint"`
	Version         string             `bson:"version"`
	Alias           string             `bson:"alias"`
	DisabledSources []string           `bson:"disabledSources"`
	IsDisabled      bool               `bson:"isDisabled"`
}
