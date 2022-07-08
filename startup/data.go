package startup

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"profile-service/domain"
)

var profiles = []*domain.Profile{
	{
		Id:          getObjectId("623b0cc3a34d25d8567f9f82"),
		Name:        "Milan",
		LastName:    "Savic",
		Username:    "milan",
		Email:       "milan@gmail.com",
		Password:    "milan",
		PhoneNumber: "+381 63 123 45 67",
		Biography:   "",
		Education: []domain.Education{
			{
				Id:           getObjectId("623b0cc3a34d25d8567f9f09"),
				School:       "Fakultet tehnickih nauka, Novi Sad",
				FieldOfStudy: "Softversko inzenjerstvo",
				Degree:       "Master's degree",
			},
		},
		Experience: []domain.Experience{
			{
				Id:          getObjectId("623b0cc3a34d25d8567f9f02"),
				JobTitle:    "Full Stack Developer",
				CompanyName: "Levi9, Serbia",
				Description: "Full time",
			},
			{
				Id:          getObjectId("623b0cc3a34d25d8567f9r02"),
				JobTitle:    "Backend Java Developer",
				CompanyName: "IT Labs",
				Description: "Full time",
			},
		},
		Skills:                          []string{"Java", "C", "SQL", "Java Script"},
		Interests:                       []string{"Web Programming"},
		IsPrivate:                       false,
		ReceivesMessageNotifications:    true,
		ReceivesConnectionNotifications: true,
		ReceivesPostNotifications:       true,
	},
	{
		Id:          getObjectId("623b0cc3a34d25d8567f9f83"),
		Name:        "Aleksandra",
		LastName:    "Novakovic",
		Username:    "aleksandra",
		Email:       "aleksandra@gmail.com",
		Password:    "aleksandra",
		PhoneNumber: "+381 62 123 45 67",
		Biography:   "",
		Education: []domain.Education{
			{
				Id:           getObjectId("623b0cc3a34d25d8567f9f09"),
				School:       "Fakultet tehnickih nauka, Novi Sad",
				FieldOfStudy: "Racunarstvo i automatika",
				Degree:       "Bachelor's degree",
			},
		},
		Experience: []domain.Experience{
			{
				Id:          getObjectId("623b0cc3a34d25d8567f9f02"),
				JobTitle:    "DevOps Engineer",
				CompanyName: "Symphony.is, Serbia",
				Description: "Full time",
			},
			{
				Id:          getObjectId("623b0cc3a34d25d8567f9f02"),
				JobTitle:    "Software Developer",
				CompanyName: "Endava",
				Description: "Internship",
			},
		},
		Skills:                          []string{"Java", "Docker", "AWS"},
		Interests:                       []string{},
		IsPrivate:                       false,
		ReceivesMessageNotifications:    true,
		ReceivesConnectionNotifications: true,
		ReceivesPostNotifications:       true,
	},
	{
		Id:          getObjectId("623b0cc3a34d25d8567f9f84"),
		Name:        "Stefan",
		LastName:    "Stefanovic",
		Username:    "stefan",
		Email:       "stefan@gmail.com",
		Password:    "stefan",
		PhoneNumber: "+381 64 123 45 67",
		Biography:   "",
		Education: []domain.Education{
			{
				Id:           getObjectId("623b0cc3a34d25d8567f9f09"),
				School:       "Elektrotehnicki fakultet, Beograd",
				FieldOfStudy: "Softversko inzenjerstvo",
				Degree:       "Master-s degree",
			},
		},
		Experience: []domain.Experience{
			{
				Id:          getObjectId("623b0cc3a34d25d8567f9f02"),
				JobTitle:    "Data Scientist",
				CompanyName: "Seven Bridges, Serbia",
				Description: "Full time",
			},
		},
		Skills:                          []string{"Python", "C", "SQL", "AWS"},
		Interests:                       []string{"Artificial Intelligence", "Data Science", "Machine Learning"},
		IsPrivate:                       false,
		ReceivesMessageNotifications:    true,
		ReceivesConnectionNotifications: true,
		ReceivesPostNotifications:       true,
	},
	{
		Id:          getObjectId("623b0cc3a34d25d8567f9f87"),
		Name:        "Sara",
		LastName:    "Jovanovic",
		Username:    "sara",
		Email:       "sara@gmail.com",
		Password:    "sara",
		PhoneNumber: "+381 69 123 45 67",
		Biography:   "",
		Education: []domain.Education{
			{
				Id:           getObjectId("623b0cc3a34d25d8567f9f09"),
				School:       "Elektrotehnicki fakultet, Beograd",
				FieldOfStudy: "Softversko inzenjerstvo",
				Degree:       "Bachelor's degree",
			},
		},
		Experience: []domain.Experience{
			{
				Id:          getObjectId("623b0cc3a34d25d8567f9f02"),
				JobTitle:    "Software Developer",
				CompanyName: "Seven Bridges, Serbia",
				Description: "Full time",
			},
			{
				Id:          getObjectId("623b0cc3a34d25d8567f9f02"),
				JobTitle:    "Frontend Developer",
				CompanyName: "IT Lab, Serbia",
				Description: "Part time",
			},
		},
		Skills:                          []string{"Python", "C#", "NoSQL"},
		Interests:                       []string{"Artificial Intelligence", "Data Science", "Machine Learning"},
		IsPrivate:                       true,
		ReceivesMessageNotifications:    true,
		ReceivesConnectionNotifications: true,
		ReceivesPostNotifications:       true,
	},
	{
		Id:          getObjectId("623b0cc3a34d25d8567f9f88"),
		Name:        "Filip",
		LastName:    "Filipovic",
		Username:    "filip",
		Email:       "filip@gmail.com",
		Password:    "filip",
		PhoneNumber: "+381 66 123 45 67",
		Biography:   "",
		Education: []domain.Education{
			{
				Id:           getObjectId("623b0cc3a34d25d8567f9f09"),
				School:       "Elektrotehnicki fakultet, Beograd",
				FieldOfStudy: "Softversko inzenjerstvo",
				Degree:       "Master's degree",
			},
		},
		Experience: []domain.Experience{
			{
				Id:          getObjectId("623b0cc3a34d25d8567f9f02"),
				JobTitle:    "Software Developer",
				CompanyName: "Nordeus, Serbia",
				Description: "Full time",
			},
		},
		Skills:                          []string{"Python", "C#", "NoSQL", "Java"},
		Interests:                       []string{"Artificial Intelligence", "Gaming", "Machine Learning"},
		IsPrivate:                       true,
		ReceivesMessageNotifications:    true,
		ReceivesConnectionNotifications: true,
		ReceivesPostNotifications:       true,
	},
}

var notifications = []*domain.Notification{
	{
		Id:               getObjectId("158b0cc3a34d25d8567f9f01"),
		Seen:             true,
		Content:          "User Name Surname has followed you.",
		NotificationType: "connection",
		ReceiverId:       "623b0cc3a34d25d8567f9f82",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
