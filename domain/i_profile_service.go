package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"profile-service/dto"
)

type IProfileService interface {
	Get(id primitive.ObjectID) (*Profile, error)
	GetAll() ([]*Profile, error)
	Insert(profile *Profile) error
	DeleteAll()
	Update(id primitive.ObjectID, profile *Profile) (*Profile, error)
	GetByName(name string) ([]*Profile, error)
	GetCredentials(username string) (*dto.CredentialsDTO, error)
}
