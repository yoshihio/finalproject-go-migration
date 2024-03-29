package data

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)
// struct for dummy data
type Event struct {
	Name string `bson:"name"`
	MakerUsername string `bson:makerUsername`
	Type string `bson:"type"`
	Games []string `bson:"games"`
	Category []string `bson:"category"`
	Description string `bson:"description"`
	Site string `bson:"site"`
	DateStart time.Time `bson:"dateStart"`
	DateEnd time.Time `bson:"dateEnd"`
	Location string `bson:"location"`
	Latitude string `bson:"latitude"`
	Longitude string `bson:"longitude"`
	Poster string `bson:"poster"`
}

func EventData(db *mongo.Database) {
	col := db.Collection("event")
	col.InsertOne(context.TODO(),
		Event{
			Name: "Event 1",
			MakerUsername: "TestUser",
			Type: "Competition",
			Games: []string{
				"Dota",
				"PUBG",
			},
			Category: []string{
				"Moba",
				"Shooting",
			},
			Description: "Description 1",
			DateStart: time.Now(),
			DateEnd: time.Now(),
			Site: "Online",
			Location: "Location 1",
			Poster: "Image 1",
			Latitude: "-6.195415",
			Longitude: "106.822329",
		})
	col.InsertOne(context.TODO(),
		Event{
			Name: "Event 2",
			MakerUsername: "Admin",
			Type: "Event",
			Games: []string{
				"YuGiOh",
			},
			Category: []string{
				"TCG",
			},
			Description: "Description 1",
			DateStart: time.Now(),
			DateEnd: time.Now(),
			Site: "Online",
			Location: "Location 1",
			Poster: "Image 1",
			Latitude: "-6.256272",
			Longitude: "106.618268",
		})
}