package api

import (
	pb "github.com/XWS-DISLINKT/dislinkt/common/proto/profile-service"
	"profile-service/domain"
)

func mapProfile(profile *domain.Profile) *pb.Profile {
	profilePb := &pb.Profile{
		Id:       profile.Id.Hex(),
		Name:     profile.Name,
		LastName: profile.LastName,
		Username: profile.Username,
		Password: profile.Password,
	}
	return profilePb
}
