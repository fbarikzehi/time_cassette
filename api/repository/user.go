package repository

import (
	"context"
	"time"
	"timecassette_api/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id"`
	Email          string             `bson:"email" validate:"required, min=5,email"`
	FirstName      string             `bson:"first_name" validate:"required,min=50"`
	LastName       string             `bson:"last_name" validate:"required,min=50"`
	Avatar         *string            `bson:"avatar"  validate:"url"`
	PhoneNumber    *string            `bson:"phone_number" validate:"omitempty"`
	CreateDateTime time.Time          `bson:"create_DateTime"`
	UpdateDateTime time.Time          `bson:"update_DateTime"`
}

var usersCollection *mongo.Collection = OpenCollection(Client, "users")

func FindUserById(ctx context.Context, cancel context.CancelFunc, userId primitive.ObjectID) (User, bool) {
	var user User
	if usersCollection == nil {
		return user, false
	}
	filter := bson.M{"_id": userId}
	err := usersCollection.FindOne(ctx, filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		cancel()
		return user, false
	} else {

		//TODO: jwt VerifyToken
		// jwtPayload := utils.JwtPayload{Username: user.Username, Email: user.Email}
		// if _, err := utils.VerifyToken(token, jwtPayload); err != nil {
		// 	return user, false
		// }

		return user, true
	}
}

func FindUserByEmail(ctx context.Context, cancel context.CancelFunc, email string) (User, bool) {
	var user User
	if usersCollection == nil {
		return user, false
	}
	filter := bson.M{"email": email}
	err := usersCollection.FindOne(ctx, filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		cancel()
		return user, false
	} else {
		return user, true
	}
}

func SearchEmails(ctx context.Context, cancel context.CancelFunc, searchValue string) (users []models.EmailModel, ok bool) {
	var list []models.EmailModel
	if usersCollection == nil {
		return list, false
	}

	query := bson.M{"email": bson.M{"$in": searchValue}}
	cursor, err := usersCollection.Find(ctx, query)
	if err != mongo.ErrNoDocuments {
		defer cancel()
		return list, false
	} else {
		if err = cursor.All(ctx, &list); err != nil {
			return list, false
		}
		return list, true
	}
}
