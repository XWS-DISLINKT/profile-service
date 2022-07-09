package application

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"profile-service/domain"
	"profile-service/dto"
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

func (service *ProfileService) Update(id primitive.ObjectID, profile *domain.Profile) (*domain.Profile, error) {
	return service.iProfileService.Update(id, profile)
}

func (service *ProfileService) GetByName(name string) ([]*domain.Profile, error) {
	return service.iProfileService.GetByName(name)
}

func (service *ProfileService) GetCredentials(username string) (*dto.CredentialsDTO, error) {
	return service.iProfileService.GetCredentials(username)
}

func (service *ProfileService) SendMessage(message *domain.Message) error {
	return service.iProfileService.SendMessage(message)
}

func (service *ProfileService) GetChatMessages(senderId primitive.ObjectID, receiverId primitive.ObjectID) ([]*domain.Message, error) {
	return service.iProfileService.GetChatMessages(senderId, receiverId)
}

func (service *ProfileService) GetNotificationsByUserId(receiverId string) ([]*domain.Notification, error) {
	return service.iProfileService.GetNotificationsByUserId(receiverId)
}

func (service *ProfileService) SendNotification(notification *domain.Notification) error {
	return service.iProfileService.SendNotification(notification)
}
