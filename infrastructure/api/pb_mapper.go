package api

import (
	pb "github.com/XWS-DISLINKT/dislinkt/common/proto/profile-service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
	"profile-service/domain"
	"time"
)

func mapProfile(profile *domain.Profile) *pb.Profile {
	profilePb := &pb.Profile{
		Id:          profile.Id.Hex(),
		Name:        profile.Name,
		LastName:    profile.LastName,
		Username:    profile.Username,
		Email:       profile.Email,
		Password:    profile.Password,
		DateOfBirth: timestamppb.New(profile.DateOfBirth),
		PhoneNumber: profile.PhoneNumber,
		Gender:      mapGender(profile.Gender),
		Biography:   profile.Biography,
		Headline:    profile.Headline,
		Experience:  make([]*pb.Experience, 0),
		Education:   make([]*pb.Education, 0),
		Skills:      make([]*pb.Skill, 0),
		Interests:   make([]*pb.Interest, 0),
	}

	for _, experience := range profile.Experience {
		experiencePb := &pb.Experience{
			Id:          experience.Id.Hex(),
			JobTitle:    experience.JobTitle,
			CompanyName: experience.CompanyName,
			Description: experience.Description,
			StartDate:   timestamppb.New(experience.StartDate),
			EndDate:     timestamppb.New(experience.EndDate),
		}
		profilePb.Experience = append(profilePb.Experience, experiencePb)
	}

	for _, education := range profile.Education {
		educationPb := &pb.Education{
			Id:           education.Id.Hex(),
			School:       education.School,
			FieldOfStudy: education.FieldOfStudy,
			Degree:       education.Degree,
		}
		profilePb.Education = append(profilePb.Education, educationPb)
	}

	return profilePb
}

func mapNewProfile(profilePb *pb.NewProfile) *domain.Profile {
	profile := &domain.Profile{
		Id:          primitive.ObjectID{},
		Name:        profilePb.Name,
		LastName:    profilePb.LastName,
		Username:    "",
		Email:       profilePb.Email,
		Password:    profilePb.Password,
		DateOfBirth: time.Time{},
		PhoneNumber: "",
		Gender:      0,
		Biography:   "",
		Headline:    "",
		Experience:  nil,
		Education:   nil,
		Skills:      nil,
		Interests:   nil,
	}
	return profile
}

func mapGender(gender domain.Gender) pb.Profile_Gender {
	switch gender {
	case domain.Male:
		return pb.Profile_Male
	case domain.Female:
		return pb.Profile_Female
	}
	return pb.Profile_Other
}
