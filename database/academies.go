package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	m "grabber/models"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *DB) GetAcademies() ([]*m.Academy, int, error) {
	collection := db.database.Collection("academies")

	var academies []*m.Academy

	cur, err := collection.Find(db.ctx, bson.D{{}})

	if err != nil || cur == nil {
		log.Fatal(err)
		return nil, 0, err
	}

	for cur.Next(db.ctx) {
		var t m.Academy
		err := cur.Decode(&t)
		if err != nil {
			log.Error(err)
		}
		academies = append(academies, &t)
	}

	// once exhausted, close the cursor
	cur.Close(db.ctx)

	if len(academies) == 0 {
		log.Warn("No academies founded")
	}

	return academies, len(academies), nil
}

func (db *DB) GetAcademy(id string) (m.Academy, error) {
	collection := db.database.Collection("academies")
	objID, _ := primitive.ObjectIDFromHex(id)

	var academy m.Academy
	err := collection.FindOne(db.ctx, bson.M{
		"_id": objID,
	}).Decode(&academy)

	if err != nil {
		log.Fatal(err)
	}

	return academy, nil
}
