package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Profile struct {
	Id          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	LastName    string             `bson:"lastName"`
	Username    string             `bson:"username"`
	Email       string             `bson:"email"`
	Password    string             `bson:"password"`
	DateOfBirth time.Time          `bson:"dateOfBirth"`
	PhoneNumber string             `bson:"phoneNumber"`
	Gender      Gender             `bson:"gender"`
	Biography   string             `bson:"biography"`
	Headline    string             `bson:"headline"`
	Experience  []Experience       `bson:"experience"`
	Education   []Education        `bson:"education"`
	Skills      []string           `bson:"skills"`
	Interests   []string           `bson:"interests"`
	IsPrivate   bool               `bson:"isPrivate"`
}

type Message struct {
	Id               primitive.ObjectID `bson:"_id"`
	Text             string             `bson:"text"`
	Date             time.Time          `bson:"date"`
	Seen             bool               `bson:"seen"`
	SenderUsername   string             `bson:"senderUsername"`
	SenderId         string             `bson:"senderId"`
	ReceiverUsername string             `bson:"receiverUsername"`
	ReceiverId       string             `bson:"receiverId"`
}
