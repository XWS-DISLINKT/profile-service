package api

import (
	pb "github.com/XWS-DISLINKT/dislinkt/common/proto/profile-service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
	"profile-service/domain"
	"time"
)

func mapToDomain(profilePb *pb.Profile) *domain.Profile {
	profile := &domain.Profile{
		Id:          getObjectId(profilePb.Id),
		Name:        profilePb.Name,
		LastName:    profilePb.LastName,
		Username:    profilePb.Username,
		Email:       profilePb.Email,
		Password:    profilePb.Password,
		DateOfBirth: profilePb.DateOfBirth.AsTime(),
		PhoneNumber: profilePb.PhoneNumber,
		Gender:      domain.Gender(profilePb.Gender),
		Biography:   profilePb.Biography,
		Headline:    profilePb.Headline,
		Experience:  make([]domain.Experience, 0),
		Education:   make([]domain.Education, 0),
		Skills:      make([]string, 0),
		Interests:   make([]string, 0),
		IsPrivate:   profilePb.IsPrivate,
	}

	for _, skillPb := range profilePb.Skills {
		skill := skillPb
		profile.Skills = append(profile.Skills, skill)
	}

	for _, experiencePb := range profilePb.Experience {
		experience := domain.Experience{
			Id:          getObjectId(experiencePb.Id),
			JobTitle:    experiencePb.JobTitle,
			CompanyName: experiencePb.CompanyName,
			Description: experiencePb.Description,
			StartDate:   experiencePb.StartDate.AsTime(),
			EndDate:     experiencePb.EndDate.AsTime(),
		}
		profile.Experience = append(profile.Experience, experience)
	}

	for _, educationPb := range profilePb.Education {
		education := domain.Education{
			Id:           getObjectId(educationPb.Id),
			School:       educationPb.School,
			FieldOfStudy: educationPb.FieldOfStudy,
			Degree:       educationPb.Degree,
		}
		profile.Education = append(profile.Education, education)
	}

	return profile
}

func mapProfile(profile *domain.Profile) *pb.Profile {
	profilePb := &pb.Profile{
		Id:                              profile.Id.Hex(),
		Name:                            profile.Name,
		LastName:                        profile.LastName,
		Username:                        profile.Username,
		Email:                           profile.Email,
		Password:                        profile.Password,
		DateOfBirth:                     timestamppb.New(profile.DateOfBirth),
		PhoneNumber:                     profile.PhoneNumber,
		Gender:                          mapGender(profile.Gender),
		Biography:                       profile.Biography,
		Headline:                        profile.Headline,
		Experience:                      make([]*pb.Experience, 0),
		Education:                       make([]*pb.Education, 0),
		Skills:                          make([]string, 0),
		Interests:                       make([]string, 0),
		IsPrivate:                       profile.IsPrivate,
		ReceivesMessageNotifications:    profile.ReceivesMessageNotifications,
		ReceivesPostNotifications:       profile.ReceivesPostNotifications,
		ReceivesConnectionNotifications: profile.ReceivesConnectionNotifications,
	}

	for _, skill := range profile.Skills {
		skillPb := skill
		profilePb.Skills = append(profilePb.Skills, skillPb)
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
		Id:          primitive.NewObjectID(),
		Name:        profilePb.Name,
		LastName:    profilePb.LastName,
		Username:    profilePb.Username,
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
		IsPrivate:   profilePb.IsPrivate,
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

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}

func mapToDomainMessage(messagePb *pb.Message) *domain.Message {
	message := &domain.Message{
		Id:               primitive.NewObjectID(),
		Text:             messagePb.Text,
		Date:             messagePb.Date.AsTime(),
		Seen:             messagePb.Seen,
		SenderUsername:   messagePb.SenderUsername,
		SenderId:         messagePb.SenderId,
		ReceiverUsername: messagePb.ReceiverUsername,
		ReceiverId:       messagePb.ReceiverId,
	}
	return message
}

func mapMessage(message *domain.Message) *pb.Message {
	messagePb := &pb.Message{
		Text:             message.Text,
		Date:             timestamppb.New(message.Date),
		Seen:             message.Seen,
		SenderUsername:   message.SenderUsername,
		SenderId:         message.SenderId,
		ReceiverUsername: message.ReceiverUsername,
		ReceiverId:       message.ReceiverId,
	}

	return messagePb
}

func mapNotification(notification *domain.Notification) *pb.Notification {
	notificationPb := &pb.Notification{
		Id:               notification.Id.Hex(),
		Content:          notification.Content,
		Date:             timestamppb.New(notification.Date),
		Seen:             notification.Seen,
		NotificationType: notification.NotificationType,
		ReceiverId:       notification.ReceiverId,
	}

	return notificationPb
}
func mapToDomainNotification(notificationPb *pb.Notification) *domain.Notification {
	notification := &domain.Notification{
		Id:               primitive.NewObjectID(),
		Content:          notificationPb.Content,
		Date:             notificationPb.Date.AsTime(),
		Seen:             notificationPb.Seen,
		NotificationType: notificationPb.NotificationType,
		ReceiverId:       notificationPb.ReceiverId,
	}
	return notification
}
