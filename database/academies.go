package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	m "grabber/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (db *DB) GetAcademies() ([]*m.Academy, int, error) {
	collection := db.database.Collection("academies")
	var academies []*m.Academy
	cur, err := collection.Find(db.ctx, bson.D{{}})

	if err != nil {
		return nil, 0, err
	}

	for cur.Next(db.ctx) {
		var t m.Academy
		if err := cur.Decode(&t); err != nil {
			return nil, 0, err
		}
		academies = append(academies, &t)
	}

	// once exhausted, close the cursor
	if err = cur.Close(db.ctx); err != nil {
		return academies, len(academies), err
	}

	return academies, len(academies), nil
}

func (db *DB) GetAcademy(id string) (*m.Academy, error) {
	collection := db.database.Collection("academies")
	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var academy *m.Academy
	err = collection.FindOne(db.ctx, bson.M{
		"_id": objID,
	}).Decode(&academy)

	if err != nil || academy == nil {
		return nil, err
	}

	return academy, nil
}
