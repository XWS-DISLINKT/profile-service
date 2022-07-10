package api

import (
	events "github.com/XWS-DISLINKT/dislinkt/common/saga/create_user"
	saga "github.com/XWS-DISLINKT/dislinkt/common/saga/messaging"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"profile-service/application"
	"profile-service/domain"
)

type CreateUserCommandHandler struct {
	profileService    *application.ProfileService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewCreateOrderCommandHandler(profileService *application.ProfileService, publisher saga.Publisher, subscriber saga.Subscriber) (*CreateUserCommandHandler, error) {
	o := &CreateUserCommandHandler{
		profileService:    profileService,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *CreateUserCommandHandler) handle(command *events.CreateUserCommand) {
	id, err := primitive.ObjectIDFromHex(command.User.Id)
	if err != nil {
		return
	}
	profile := &domain.Profile{Id: id}

	reply := events.CreateUserReply{User: command.User}

	switch command.Type {
	case events.ApproveUser:
		err := handler.profileService.Approve(profile)
		if err != nil {
			return
		}
		reply.Type = events.UnknownReply
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
