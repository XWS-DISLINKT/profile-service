package application

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"profile-service/domain"
)

type ProfileService struct {
	iProfileService domain.IProfileService
}

func NewProfileService(iProfileService domain.IProfileService) *ProfileService {
	return &ProfileService{
		iProfileService: iProfileService,
	}
}

func (service *ProfileService) Get(id primitive.ObjectID) (*domain.Profile, error) {
	return service.iProfileService.Get(id)
}

func (service *ProfileService) GetAll() ([]*domain.Profile, error) {
	return service.iProfileService.GetAll()
}

func (service *ProfileService) Create(profile *domain.Profile) error {
	return service.iProfileService.Insert(profile)
}
