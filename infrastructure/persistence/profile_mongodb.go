package persistence

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"profile-service/domain"
	"profile-service/dto"
)

const (
	DATABASE   = "profile"
	COLLECTION = "profile"
)

type ProfileMongoDb struct {
	profiles *mongo.Collection
}

func NewProfileMongoDb(client *mongo.Client) domain.IProfileService {
	profiles := client.Database(DATABASE).Collection(COLLECTION)
	return &ProfileMongoDb{
		profiles: profiles,
	}
}

func (collection *ProfileMongoDb) Get(id primitive.ObjectID) (*domain.Profile, error) {
	filter := bson.M{"_id": id}
	return collection.filterOne(filter)
}

func (collection *ProfileMongoDb) GetAll() ([]*domain.Profile, error) {
	filter := bson.D{{}}
	return collection.filter(filter)
}

func (collection *ProfileMongoDb) Insert(profile *domain.Profile) error {
	//profile.Id = primitive.NewObjectID()
	result, err := collection.profiles.InsertOne(context.TODO(), profile)
	if err != nil {
		return err
	}
	profile.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (collection *ProfileMongoDb) DeleteAll() {
	collection.profiles.DeleteMany(context.TODO(), bson.D{{}})
}

func (collection *ProfileMongoDb) Update(id primitive.ObjectID, profile *domain.Profile) (*domain.Profile, error) {
	result, err := collection.profiles.UpdateOne(
		context.TODO(), bson.M{"_id": id}, bson.D{
			{"$set", bson.D{
				{"name", profile.Name},
				{"lastName", profile.LastName},
				{"username", profile.Username},
				{"email", profile.Email},
				{"password", profile.Password},
				{"dateOfBirth", profile.DateOfBirth},
				{"phoneNumber", profile.PhoneNumber},
				{"gender", profile.Gender},
				{"biography", profile.Biography},
				{"experience", profile.Experience},
				{"education", profile.Education},
				{"skills", profile.Skills},
				{"interests", profile.Interests},
			}},
		})
	if err != nil {
		return nil, err
	}
	fmt.Printf("Updated %v", result.ModifiedCount)
	filter := bson.M{"_id": id}
	return collection.filterOne(filter)
}

func (collection *ProfileMongoDb) GetByName(name string) ([]*domain.Profile, error) {
	filter := bson.D{{"name", name}, {"isPrivate", false}}
	return collection.filter(filter)
}

func (collection *ProfileMongoDb) GetCredentials(username string) (*dto.CredentialsDTO, error) {
	filter := bson.M{"username": username}
	profile, err := collection.filterOne(filter)
	if err != nil {
		return nil, err
	}
	credentialsDTO := &dto.CredentialsDTO{
		Username: profile.Username,
		Password: profile.Password,
		Id:       profile.Id.Hex(),
	}
	return credentialsDTO, nil
}

func (collection *ProfileMongoDb) filter(filter interface{}) ([]*domain.Profile, error) {
	cursor, err := collection.profiles.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (collection *ProfileMongoDb) filterOne(filter interface{}) (profile *domain.Profile, err error) {
	result := collection.profiles.FindOne(context.TODO(), filter)
	err = result.Decode(&profile)
	return
}

func decode(cursor *mongo.Cursor) (profiles []*domain.Profile, err error) {
	for cursor.Next(context.TODO()) {
		var profile domain.Profile
		err = cursor.Decode(&profile)
		if err != nil {
			return
		}
		profiles = append(profiles, &profile)
	}
	err = cursor.Err()
	return
}
