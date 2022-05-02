package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Education struct {
	Id           primitive.ObjectID `bson:"id"`
	School       string             `bson:"school"`
	FieldOfStudy string             `bson:"fieldOfStudy"`
	Degree       string             `bson:"degree"`
}
