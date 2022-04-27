package startup

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"profile-service/domain"
)

var profiles = []*domain.Profile{
	{
		Id:       getObjectId("623b0cc3a34d25d8567f9f82"),
		Name:     "milomir",
		LastName: "maric",
		Username: "mile",
		Password: "mile",
	},
	{
		Id:       getObjectId("623b0cc3a34d25d8567f9f83"),
		Name:     "pera",
		LastName: "peric",
		Username: "pera",
		Password: "pera",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
