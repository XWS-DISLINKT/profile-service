package startup

import (
	"fmt"
	profile "github.com/XWS-DISLINKT/dislinkt/common/proto/profile-service"
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

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	iProfileService := server.initIProfileService(mongoClient)

	profileService := server.initProfileService(iProfileService)

	profileHandler := server.initProfileHandler(profileService)

	server.startGrpcServer(profileHandler)
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
	return collection
}

func (server *Server) initProfileService(iProfileService domain.IProfileService) *application.ProfileService {
	return application.NewProfileService(iProfileService)
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
