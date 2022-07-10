package startup

import (
	"fmt"
	profile "github.com/XWS-DISLINKT/dislinkt/common/proto/profile-service"
	saga "github.com/XWS-DISLINKT/dislinkt/common/saga/messaging"
	"github.com/XWS-DISLINKT/dislinkt/common/saga/messaging/nats"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
	"profile-service/application"
	"profile-service/domain"
	"profile-service/infrastructure/api"
	"profile-service/infrastructure/persistence"
	"profile-service/startup/config"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

const (
	QueueGroup = "profile_service"
)

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	iProfileService := server.initIProfileService(mongoClient)

	commandPublisher := server.initPublisher(server.config.CreateUserCommandSubject)
	replySubscriber := server.initSubscriber(server.config.CreateUserReplySubject, QueueGroup)
	orchestrator := server.initOrchestrator(commandPublisher, replySubscriber)

	profileService := server.initProfileService(iProfileService, orchestrator)

	commandSubscriber := server.initSubscriber(server.config.CreateUserCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.CreateUserReplySubject)
	server.initCreateUserHandler(profileService, replyPublisher, commandSubscriber)

	profileHandler := server.initProfileHandler(profileService)

	server.startGrpcServer(profileHandler)
}

func (server *Server) initPublisher(subject string) saga.Publisher {
	publisher, err := nats.NewNATSPublisher(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject)
	if err != nil {
		log.Fatal(err)
	}
	return publisher
}

func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}

func (server *Server) initOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) *application.CreateUserOrchestrator {
	orchestrator, err := application.NewCreateUserOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}

func (server *Server) initCreateUserHandler(service *application.ProfileService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewCreateOrderCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}
func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.ProfileDBHost, server.config.ProfileDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initIProfileService(client *mongo.Client) domain.IProfileService {
	collection := persistence.NewProfileMongoDb(client)
	collection.DeleteAll()
	for _, profile := range profiles {
		err := collection.Insert(profile)
		if err != nil {
			log.Fatal(err)
		}
	}
	for _, notification := range notifications {
		err := collection.InsertNotification(notification)
		if err != nil {
			log.Fatal(err)
		}
	}

	return collection
}

func (server *Server) initProfileService(iProfileService domain.IProfileService, orchestrator *application.CreateUserOrchestrator) *application.ProfileService {
	return application.NewProfileService(iProfileService, orchestrator)
}

func (server *Server) initProfileHandler(service *application.ProfileService) *api.ProfileHandler {
	return api.NewProfileHandler(service, *server.config)
}

func (server *Server) startGrpcServer(profileHandler *api.ProfileHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	profile.RegisterProfileServiceServer(grpcServer, profileHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
