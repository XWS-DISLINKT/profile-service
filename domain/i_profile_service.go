package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IProfileService interface {
	Get(id primitive.ObjectID) (*Profile, error)
	GetAll() ([]*Profile, error)
	Insert(profile *Profile) error
	DeleteAll()
}
