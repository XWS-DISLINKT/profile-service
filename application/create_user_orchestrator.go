package application

import (
	events "github.com/XWS-DISLINKT/dislinkt/common/saga/create_user"
	saga "github.com/XWS-DISLINKT/dislinkt/common/saga/messaging"
	"profile-service/domain"
)

type CreateUserOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewCreateUserOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*CreateUserOrchestrator, error) {
	o := &CreateUserOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *CreateUserOrchestrator) Start(user *domain.Profile) error {
	event := &events.CreateUserCommand{
		User: events.User{
			Id:        user.Id.Hex(),
			IsPrivate: user.IsPrivate,
		},
		Type: events.CreateUserConnection,
	}
	return o.commandPublisher.Publish(event)
}

func (o *CreateUserOrchestrator) handle(reply *events.CreateUserReply) {
	command := events.CreateUserCommand{
		User: reply.User,
	}
	command.Type = o.nextCommandType(reply.Type)
	if command.Type != events.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
	}
}

func (o *CreateUserOrchestrator) nextCommandType(reply events.CreateUserReplyType) events.CreateUserCommandType {
	switch reply {
	case events.UserConnectionCreated:
		return events.ApproveUser
	case events.UserConnectionNotCreated:
		return events.UnknownCommand
	default:
		return events.UnknownCommand
	}
}
