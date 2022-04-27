package api

import (
	"context"
	pb "github.com/XWS-DISLINKT/dislinkt/common/proto/profile-service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"profile-service/application"
)

type ProfileHandler struct {
	pb.UnsafeProfileServiceServer
	service *application.ProfileService
}

func NewProfileHandler(service *application.ProfileService) *ProfileHandler {
	return &ProfileHandler{
		service: service,
	}
}

func (handler *ProfileHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	profile, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	profilePb := mapProfile(profile)
	response := &pb.GetResponse{
		Profile: profilePb,
	}
	return response, nil
}

func (handler *ProfileHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
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
