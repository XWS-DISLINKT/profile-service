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
	DATABASE                 = "profile"
	COLLECTION               = "profile"
	MESSAGES_COLLECTION      = "messages"
	NOTIFICATIONS_COLLECTION = "notifications"
)

type ProfileMongoDb struct {
	profiles      *mongo.Collection
	messages      *mongo.Collection
	notifications *mongo.Collection
}

func NewProfileMongoDb(client *mongo.Client) domain.IProfileService {
	profiles := client.Database(DATABASE).Collection(COLLECTION)
	messages := client.Database(DATABASE).Collection(MESSAGES_COLLECTION)
	notifications := client.Database(DATABASE).Collection(NOTIFICATIONS_COLLECTION)
	return &ProfileMongoDb{
		profiles:      profiles,
		messages:      messages,
		notifications: notifications,
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

func (collection *ProfileMongoDb) SendMessage(message *domain.Message) error {
	result, err := collection.messages.InsertOne(context.TODO(), message)
	if err != nil {
		return err
	}
	message.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (collection *ProfileMongoDb) GetChatMessages(senderId primitive.ObjectID, receiverId primitive.ObjectID) ([]*domain.Message, error) {
	//filter := bson.D{{}}
	filter := bson.M{
		"$or": []bson.M{
			{
				"senderId":   senderId.Hex(),
				"receiverId": receiverId.Hex(),
			},
			{
				"senderId":   receiverId.Hex(),
				"receiverId": senderId.Hex(),
			},
		},
	}

	return collection.filterMessages(filter)
}
func (collection *ProfileMongoDb) SendNotification(notification *domain.Notification) error {
	result, err := collection.notifications.InsertOne(context.TODO(), notification)
	if err != nil {
		return err
	}
	notification.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}
func (collection *ProfileMongoDb) GetNotificationsByUserId(receiverId string) ([]*domain.Notification, error) {
	//filter := bson.D{{}}
	filter := bson.M{"receiverId": receiverId}

	fmt.Println("{1}", receiverId)
	return collection.filterNotifications(filter)
}

func (collection *ProfileMongoDb) DeleteAll() {
	collection.profiles.DeleteMany(context.TODO(), bson.D{{}})
	collection.notifications.DeleteMany(context.TODO(), bson.D{{}})
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
				{"isPrivate", profile.IsPrivate},
				{"receivesConnectionNotifications", profile.ReceivesConnectionNotifications},
				{"receivesMessageNotifications", profile.ReceivesMessageNotifications},
				{"receivesPostNotifications", profile.ReceivesPostNotifications},
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

func (collection *ProfileMongoDb) filterMessages(filter interface{}) ([]*domain.Message, error) {
	cursor, err := collection.messages.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())
	if err != nil {
		return nil, err
	}

	return decodeMessages(cursor)
}

func (collection *ProfileMongoDb) filterNotifications(filter interface{}) ([]*domain.Notification, error) {
	cursor, err := collection.notifications.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())
	if err != nil {
		return nil, err
	}

	return decodeNotifications(cursor)
}

func (collection *ProfileMongoDb) InsertNotification(notification *domain.Notification) error {
	result, err := collection.notifications.InsertOne(context.TODO(), notification)
	if err != nil {
		return err
	}
	notification.Id = result.InsertedID.(primitive.ObjectID)
	return nil
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

func decodeMessages(cursor *mongo.Cursor) (messages []*domain.Message, err error) {
	for cursor.Next(context.TODO()) {
		var message domain.Message
		err = cursor.Decode(&message)
		if err != nil {
			return
		}
		messages = append(messages, &message)
	}
	err = cursor.Err()
	return
}

func decodeNotifications(cursor *mongo.Cursor) (notifications []*domain.Notification, err error) {
	for cursor.Next(context.TODO()) {
		var notification domain.Notification
		err = cursor.Decode(&notification)
		if err != nil {
			return
		}
		notifications = append(notifications, &notification)
	}
	err = cursor.Err()
	return
}
