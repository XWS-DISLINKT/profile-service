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
		Email:    "mile@gmail.com",
		Password: "mile",
		Education: []domain.Education{
			{
				Id:           getObjectId("623b0cc3a34d25d8567f9f09"),
				School:       "Fakultet tehnickih nauka, Novi Sad",
				FieldOfStudy: "Softversko inzenjerstvo",
				Degree:       "Master",
			},
		},
		Skills: []string{"Java", "C"},
	},
	{
		Id:       getObjectId("623b0cc3a34d25d8567f9f83"),
		Name:     "pera",
		LastName: "peric",
		Username: "pera",
		Password: "pera",
	},
	{
		Id:       getObjectId("623b0cc3a34d25d8567f9f84"),
		Name:     "pera",
		LastName: "petrovic",
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
