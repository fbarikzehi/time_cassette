package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"timecassette_api/models"
	"timecassette_api/utils"
)

type Cassette struct {
	ID             primitive.ObjectID `bson: "_id"`
	UserId         primitive.ObjectID `bson:"userId"`
	Name           string             `bson:"name" validate:"min=2, max=50"`
	Description    string             `bson:"description"`
	Color          string             `bson:"color"`
	IsPrivate      bool               `bson:"isPrivate" default:"true"`
	CreateDateTime time.Time          `bson:"create_DateTime"`
	UpdateDateTime time.Time          `bson:"update_DateTime"`
	// Fragments      []Fragment         `bson:"fragments"`
}

var cassettesCollection *mongo.Collection = OpenCollection(Client, "cassettes")

func CreateCassette(data models.CassetteCreateRequestModel, ctx context.Context, cancel context.CancelFunc, userId primitive.ObjectID) (createdID string, message string, ok bool) {
	if cassettesCollection == nil {
		return primitive.NilObjectID.String(), "", false
	}

	var cassette Cassette
	filter := bson.M{"name": data.Name, "userId": userId}
	err := cassettesCollection.FindOne(ctx, filter).Decode(&cassette)
	if err == mongo.ErrNoDocuments {
		cassette.Name = data.Name
		cassette.Description = data.Description
		cassette.IsPrivate = data.IsPrivate
		cassette.Color = utils.RandomColorString(6)
		cassette.UserId = userId
		cassette.CreateDateTime = time.Now()
		insertOneResult, err := cassettesCollection.InsertOne(ctx, cassette)
		if err != nil || insertOneResult.InsertedID == nil {
			defer cancel()
			return primitive.NilObjectID.String(), "Sorry! Server side error.We will fix this ASAP", false

		} else {
			oid, _ := insertOneResult.InsertedID.(primitive.ObjectID)
			return oid.Hex(), "Cassette created", true
		}

	} else {
		defer cancel()
		return primitive.NilObjectID.String(), "You have a cassette with this name", false
	}

}

func UpdateCassette(data models.CassetteUpdateRequestModel, ctx context.Context, cancel context.CancelFunc, userId primitive.ObjectID) (message string, ok bool) {
	if cassettesCollection == nil {
		return "", false
	}

	id, _ := primitive.ObjectIDFromHex(data.Id)
	filter := bson.M{"_id": id, "userId": userId}
	update := bson.M{"$set": bson.M{"name": data.Name, "description": data.Description}}
	result, err := cassettesCollection.UpdateOne(ctx, filter, update)
	if err != nil || result.MatchedCount == 0 {
		defer cancel()
		return "Sorry! Server side error.We will fix this ASAP", false

	} else {
		return "Cassette updated", true
	}

}

func DeleteCassette(data models.CassetteDeleteRequestModel, ctx context.Context, cancel context.CancelFunc, userId primitive.ObjectID) (message string, ok bool) {
	if cassettesCollection == nil {
		return "", false
	}

	id, _ := primitive.ObjectIDFromHex(data.Id)
	filter := bson.M{"_id": id, "userId": userId}
	result, err := cassettesCollection.DeleteOne(ctx, filter)
	if err != nil || result.DeletedCount == 0 {
		defer cancel()
		return "Sorry! Server side error.We will fix this ASAP", false
	} else {
		return "Cassette Deleted", true
	}

}

func FindCassettes(ctx context.Context, cancel context.CancelFunc, userId primitive.ObjectID) (cassettes []models.CassetteModel, message string, ok bool) {

	var cargo []models.CassetteModel
	if cassettesCollection == nil {
		return cargo, "", false
	}

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.M{"userId": userId}}},
		{{Key: "$lookup", Value: bson.M{
			"from":         "fragments",
			"localField":   "_id",
			"foreignField": "cassetteId",
			"as":           "cassetteFragments",
		}}},
		{{
			Key:   "$project",
			Value: bson.M{"_id": 1, "name": 1, "color": 1, "is_private": 1, "total_of_times": 1, "counts.fragment": bson.M{"$size": "$cassetteFragments"}},
		}},
	}
	cursor, err := cassettesCollection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		defer cancel()
		return cargo, "Sorry! Server side error.We will fix this ASAP", false
	} else {
		if err = cursor.All(ctx, &cargo); err != nil {
			return cargo, "Sorry! Server side error.We will fix this ASAP", false
		}
		return cargo, "", true
	}

}
