package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Profile struct {
	Id       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	LastName string             `bson:"lastName"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
}
