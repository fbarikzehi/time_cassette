package repository

import (
	"context"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type UserCredential struct {
	ID             primitive.ObjectID `bson:"_id"`
	UserId         primitive.ObjectID `bson:"userId"`
	Username       string             `bson:"username" validate:"min=5"`
	HashedPassword string             `bson:"hashedpassword" validate:"required"`
	Role           string             `bson:"role" validate:"required, eq=ADMIN|eq=USER|eq=VISITOR"`
	Token          string             `bson:"token" validate:"required"`
	TokenExpire    time.Time          `bson:"token_expire" validate:"required"`
	ReturnUrl      string             `bson:"return_url"`
	CreateDateTime time.Time          `bson:"create_DateTime"`
	UpdateDateTime time.Time          `bson:"update_DateTime"`
}

var usersCredentialCollection *mongo.Collection = OpenCollection(Client, "users_credential")

func CreateUserCredential(ctx context.Context, cancel context.CancelFunc, userCredential UserCredential) (ok bool) {
	if usersCredentialCollection == nil {
		return false
	}

	insertOneResult, err := usersCredentialCollection.InsertOne(ctx, userCredential)
	defer cancel()
	if err != nil || insertOneResult.InsertedID == nil {
		return false
	} else {
		//
		return true
	}

}

func UpdateUserCredentialToken(ctx context.Context, cancel context.CancelFunc, userId primitive.ObjectID, token string, expireAt time.Time) (ok bool) {
	if usersCredentialCollection == nil {
		return false
	}

	filter := bson.M{"userId": userId}
	update := bson.M{"$set": bson.M{"token": token, "token_expire": expireAt}}
	result, err := usersCredentialCollection.UpdateOne(ctx, filter, update)
	if err != nil || result.MatchedCount == 0 {
		defer cancel()
		return false

	}
	defer cancel()
	return true

}

func FindUserCredentialByToken(c *gin.Context, ctx context.Context, cancel context.CancelFunc) (UserCredential, bool) {
	var userCredential UserCredential
	if usersCredentialCollection == nil {
		return userCredential, false
	}
	if c.GetHeader("Authorization") == "" {
		return userCredential, false
	}

	token := strings.Split(c.GetHeader("Authorization"), " ")[1]
	if token == "" {
		return userCredential, false
	} else {

		//Find by token
		filter := bson.M{"token": token}
		err := usersCredentialCollection.FindOne(ctx, filter).Decode(&userCredential)
		if err == mongo.ErrNoDocuments {
			return userCredential, false
		} else {

			//TODO:  Security check  jwt VerifyToken
			// jwtPayload := utils.JwtPayload{Username: user.Username, Email: user.Email}
			// if _, err := utils.VerifyToken(token, jwtPayload); err != nil {
			// 	return user, false
			// }

			//Todo: Security check for token expire
			return userCredential, true
		}
	}

}

func FindUserCredentialByUserId(ctx context.Context, cancel context.CancelFunc, userId primitive.ObjectID) (UserCredential, bool) {
	var userCredential UserCredential
	if usersCredentialCollection == nil {
		return userCredential, false
	}

	//Find by userId
	filter := bson.M{"userId": userId}
	err := usersCredentialCollection.FindOne(ctx, filter).Decode(&userCredential)
	if err == mongo.ErrNoDocuments {
		cancel()
		return userCredential, false
	} else {

		//TODO:  Security check  jwt VerifyToken
		// jwtPayload := utils.JwtPayload{Username: user.Username, Email: user.Email}
		// if _, err := utils.VerifyToken(token, jwtPayload); err != nil {
		// 	return user, false
		// }

		//Todo: Security check for token expire
		if userCredential.Token == "" || userCredential.TokenExpire.Unix() <= time.Now().Unix() {
			//TODO: Save last url for new signin
			return userCredential, false
		}
		return userCredential, true
	}

}

func FindSigninUserCredentialByUserId(ctx context.Context, cancel context.CancelFunc, userId primitive.ObjectID) (UserCredential, bool) {
	var userCredential UserCredential
	if usersCredentialCollection == nil {
		return userCredential, false
	}

	//Find by userId
	filter := bson.M{"userId": userId}
	err := usersCredentialCollection.FindOne(ctx, filter).Decode(&userCredential)
	if err == mongo.ErrNoDocuments {
		cancel()
		return userCredential, false
	} else {

		//TODO:  Security check  jwt VerifyToken
		// jwtPayload := utils.JwtPayload{Username: user.Username, Email: user.Email}
		// if _, err := utils.VerifyToken(token, jwtPayload); err != nil {
		// 	return user, false
		// }

		return userCredential, true
	}

}
