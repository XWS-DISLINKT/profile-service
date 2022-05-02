package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Experience struct {
	Id          primitive.ObjectID `bson:"id"`
	JobTitle    string             `bson:"jobTitle"`
	CompanyName string             `bson:"companyName"`
	Description string             `bson:"description"`
	StartDate   time.Time          `bson:"startDate"`
	EndDate     time.Time          `bson:"endDate"`
}
