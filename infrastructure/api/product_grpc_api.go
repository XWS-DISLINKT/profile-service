package api

import (
	"context"
	"encoding/json"
	pb "github.com/XWS-DISLINKT/dislinkt/common/proto/profile-service"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
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

func (handler *ProfileHandler) Create(ctx context.Context, request *pb.CreateProfileRequest) (*pb.CreateProfileResponse, error) {
	profile := mapNewProfile(request.Profile)
	err := handler.service.Create(profile)
	if err != nil {
		return nil, err
	}
	return &pb.CreateProfileResponse{Profile: mapProfile(profile)}, nil
}

func (handler *ProfileHandler) Update(ctx context.Context, request *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error) {
	id := request.Id
	profile := request.Profile
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	updatedProfile, er := handler.service.Update(objectId, mapToDomain(profile))
	if er != nil {
		return nil, er
	}
	return &pb.UpdateProfileResponse{Profile: mapProfile(updatedProfile)}, nil
}

func (handler *ProfileHandler) GetCredentials(writer http.ResponseWriter, request *http.Request) {
	username := mux.Vars(request)["username"]
	credentialsDTO, err := handler.service.GetCredentials(username)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(credentialsDTO)
}
