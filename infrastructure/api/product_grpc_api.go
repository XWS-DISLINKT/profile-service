package api

import (
	"context"
	"fmt"
	"github.com/XWS-DISLINKT/dislinkt/common/proto/connection-service"
	pb "github.com/XWS-DISLINKT/dislinkt/common/proto/profile-service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/emptypb"
	"profile-service/application"
	"profile-service/domain"
	"profile-service/infrastructure/services"
	"profile-service/startup/config"
)

type ProfileHandler struct {
	pb.UnsafeProfileServiceServer
	service                  *application.ProfileService
	connectionServiceAddress string
}

func NewProfileHandler(service *application.ProfileService, config config.Config) *ProfileHandler {
	return &ProfileHandler{
		service:                  service,
		connectionServiceAddress: fmt.Sprintf("%s:%s", config.ConnectionHost, config.ConnectionPort),
	}
}

func (handler *ProfileHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.Profile, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		emptyProfile := pb.Profile{}
		return &emptyProfile, nil
	}
	profile, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	response := mapProfile(profile)
	return response, nil
}

func (handler *ProfileHandler) UpdateNotificationSettings(ctx context.Context, request *pb.NotificationSettings) (*pb.Profile, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		emptyProfile := pb.Profile{}
		return &emptyProfile, nil
	}
	profile, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	profile.ReceivesMessageNotifications = request.ReceivesMessageNotifications
	profile.ReceivesPostNotifications = request.ReceivesPostNotifications
	profile.ReceivesConnectionNotifications = request.ReceivesConnectionNotifications
	profile, err = handler.service.Update(objectId, profile)
	if err != nil {
		return nil, err
	}
	response := mapProfile(profile)
	return response, nil
}

func (handler *ProfileHandler) GetNotificationsByUser(ctx context.Context, request *pb.GetNotificationsRequest) (*pb.GetNotificationsResponse, error) {
	receiverId, err := primitive.ObjectIDFromHex(request.UserId)
	if err != nil {
		return nil, err
	}
	response := &pb.GetNotificationsResponse{
		Notifications: []*pb.Notification{},
	}
	notifications, err := handler.service.GetNotificationsByUserId(receiverId.Hex())

	for _, notification := range notifications {
		current := mapNotification(notification)
		response.Notifications = append(response.Notifications, current)
	}
	return response, nil
}

func (handler *ProfileHandler) SeeNotificationsByUser(ctx context.Context, request *pb.GetNotificationsRequest) (*pb.GetNotificationsResponse, error) {
	receiverId, err := primitive.ObjectIDFromHex(request.UserId)
	if err != nil {
		return nil, err
	}
	response := &pb.GetNotificationsResponse{
		Notifications: []*pb.Notification{},
	}
	notifications, err := handler.service.SeeNotificationsByUserId(receiverId.Hex())

	for _, notification := range notifications {
		current := mapNotification(notification)
		response.Notifications = append(response.Notifications, current)
	}
	return response, nil
}

func (handler *ProfileHandler) SendNotification(ctx context.Context, request *pb.NewNotificationRequest) (*pb.Notification, error) {
	return handler.CreateNotification(request.SenderId, request.ReceiverId, request.NotificationType)
}
func (handler *ProfileHandler) CreateNotification(senderId string, receiverId string, notificationType string) (*pb.Notification, error) {
	id := senderId
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, nil
	}
	profile, err := handler.service.Get(objectId)
	content := ""
	senderName := profile.Name + " " + profile.LastName
	idR := receiverId
	objectIdR, err := primitive.ObjectIDFromHex(idR)
	if err != nil {
		return nil, nil
	}
	profileR, err := handler.service.Get(objectIdR)
	if notificationType == "message" && profileR.ReceivesMessageNotifications == true {
		content = senderName + " has sent you a message"
	} else if notificationType == "post" && profileR.ReceivesPostNotifications == true {
		content = senderName + " has shared a new post"
	} else if notificationType == "request" && profileR.ReceivesConnectionNotifications == true {
		content = senderName + " has requested to follow you"
	} else if notificationType == "connection" && profileR.ReceivesConnectionNotifications == true {
		content = senderName + " has started following you"
	} else {
		content = ""
	}
	notification := &domain.Notification{
		Id:               primitive.NewObjectID(),
		Content:          content,
		Seen:             false,
		NotificationType: notificationType,
		ReceiverId:       receiverId,
	}
	if content == "" {
		return mapNotification(notification), nil
	}
	err2 := handler.service.SendNotification(notification)
	if err2 != nil {
		return nil, err2
	}
	return mapNotification(notification), nil
}

func (handler *ProfileHandler) GetChatMessages(ctx context.Context, request *pb.GetMessagesRequest) (*pb.GetMessagesFromChat, error) {
	senderId, err := primitive.ObjectIDFromHex(request.SenderId)
	receiverId, err := primitive.ObjectIDFromHex(request.ReceiverId)
	if err != nil {
		return nil, err
	}
	response := &pb.GetMessagesFromChat{
		Messages: []*pb.Message{},
	}
	messages, err := handler.service.GetChatMessages(senderId, receiverId)

	for _, message := range messages {
		current := mapMessage(message)
		response.Messages = append(response.Messages, current)
	}
	return response, nil
}

func (handler *ProfileHandler) GetAll(ctx context.Context, request *emptypb.Empty) (*pb.GetAllResponse, error) {
	profiles, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Profiles: []*pb.Profile{},
	}
	for _, profile := range profiles {
		current := mapProfile(profile)
		response.Profiles = append(response.Profiles, current)
	}
	return response, nil
}

func (handler *ProfileHandler) GetByName(ctx context.Context, request *pb.GetByNameRequest) (*pb.GetAllResponse, error) {
	name := request.Name
	profiles, err := handler.service.GetByName(name)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{Profiles: []*pb.Profile{}}
	for _, profile := range profiles {
		current := mapProfile(profile)
		response.Profiles = append(response.Profiles, current)
	}
	return response, nil
}

func (handler *ProfileHandler) Create(ctx context.Context, request *pb.NewProfile) (*pb.Profile, error) {
	profile := mapNewProfile(request)
	profile.ReceivesConnectionNotifications = true
	profile.ReceivesPostNotifications = true
	profile.ReceivesMessageNotifications = true
	profile.IsPrivate = false
	err := handler.service.Create(profile)
	if err != nil {
		return nil, err
	}

	user := connection.User{UserId: profile.Id.Hex(), IsPrivate: profile.IsPrivate}
	response, err := services.ConnectionsClient(handler.connectionServiceAddress).InsertUser(ctx, &user)
	if !response.Success {
		//implementirati sagu
	}

	return mapProfile(profile), nil
}

func (handler *ProfileHandler) SendMessage(ctx context.Context, request *pb.Message) (*pb.Message, error) {
	message := mapToDomainMessage(request)
	err := handler.service.SendMessage(message)
	if err != nil {
		return nil, err
	}
	handler.CreateNotification(request.SenderId, request.ReceiverId, "message")
	return mapMessage(message), nil
}

func (handler *ProfileHandler) Update(ctx context.Context, request *pb.Profile) (*pb.Profile, error) {
	id := request.Id
	profile := request
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	updatedProfile, er := handler.service.Update(objectId, mapToDomain(profile))
	if er != nil {
		return nil, er
	}

	user := connection.User{UserId: updatedProfile.Id.Hex(), IsPrivate: updatedProfile.IsPrivate}
	response, err := services.ConnectionsClient(handler.connectionServiceAddress).UpdateUser(ctx, &user)
	if !response.Success {
		// saga
	}

	return mapProfile(updatedProfile), nil
}

func (handler *ProfileHandler) GetCredentials(ctx context.Context, request *pb.GetCredentialsRequest) (*pb.GetCredentialsResponse, error) {
	username := request.Username
	credentials, err := handler.service.GetCredentials(username)
	if err != nil {
		return nil, err
	}
	response := &pb.GetCredentialsResponse{Username: credentials.Username, Password: credentials.Password, Id: credentials.Id}
	return response, nil
}
