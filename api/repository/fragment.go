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

type Fragment struct {
	ID             primitive.ObjectID `bson: "_id"`
	UserId         primitive.ObjectID `bson:"userId"`
	CassetteId     primitive.ObjectID `bson:"cassetteId"`
	Name           string             `bson:"name" validate:"min=2, max=50"`
	Description    string             `bson:"description"`
	Color          string             `bson:"color"`
	CreateDateTime time.Time          `bson:"create_DateTime"`
	UpdateDateTime time.Time          `bson:"update_DateTime"`
}

var fragmentsCollection *mongo.Collection = OpenCollection(Client, "fragments")

func CreateFragment(data models.FragmentCreateRequestModel, ctx context.Context, cancel context.CancelFunc, userId primitive.ObjectID) (createdID string, message string, ok bool) {
	if fragmentsCollection == nil {
		return primitive.NilObjectID.String(), "", false
	}

	var fragment Fragment
	filter := bson.M{"name": data.Name, "cassetteId": data.CassetteId, "userId": userId}
	err := fragmentsCollection.FindOne(ctx, filter).Decode(&fragment)
	if err == mongo.ErrNoDocuments {
		fragment.Name = data.Name
		fragment.Description = data.Description
		fragment.Color = utils.RandomColorString(6)
		fragment.CassetteId = data.CassetteId
		fragment.UserId = userId
		fragment.CreateDateTime = time.Now()
		insertOneResult, err := fragmentsCollection.InsertOne(ctx, fragment)
		if err != nil || insertOneResult.InsertedID == nil {
			defer cancel()
			return primitive.NilObjectID.String(), "Sorry! Server side error.We will fix this ASAP", false

		} else {
			oid, _ := insertOneResult.InsertedID.(primitive.ObjectID)
			var branch models.BranchCreateRequestModel
			handlerUser, ok := FindUserById(ctx, cancel, userId)
			if ok {
				branch.FragmentId = oid
				branch.HandlerUserEmail = handlerUser.Email
				branch.Name = "Default Branch"
				CreateBranch(branch, ctx, cancel, userId, userId)
			}
			return oid.Hex(), "Fragment created", true
		}

	} else {
		defer cancel()
		return primitive.NilObjectID.String(), "You have a fragment with this name", false
	}

}

func UpdateFragment(data models.FragmentUpdateRequestModel, ctx context.Context, cancel context.CancelFunc, userId primitive.ObjectID) (message string, ok bool) {
	if fragmentsCollection == nil {
		return "", false
	}

	id, _ := primitive.ObjectIDFromHex(data.Id)
	filter := bson.M{"_id": id, "cassetteId": data.CassetteId, "userId": userId}
	update := bson.M{"$set": bson.M{"name": data.Name, "description": data.Description}}
	result, err := fragmentsCollection.UpdateOne(ctx, filter, update)
	if err != nil || result.MatchedCount == 0 {
		defer cancel()
		return "Sorry! Server side error.We will fix this ASAP", false

	} else {
		return "Fragment updated", true
	}

}

func DeleteFragment(data models.FragmentDeleteRequestModel, ctx context.Context, cancel context.CancelFunc, userId primitive.ObjectID) (message string, ok bool) {
	if fragmentsCollection == nil {
		return "", false
	}

	id, _ := primitive.ObjectIDFromHex(data.Id)
	filter := bson.M{"_id": id, "userId": userId}
	result, err := fragmentsCollection.DeleteOne(ctx, filter)
	if err != nil || result.DeletedCount == 0 {
		defer cancel()
		return "Sorry! Server side error.We will fix this ASAP", false
	} else {
		return "Fragment Deleted", true
	}

}

func FindFragmentsByCassette(ctx context.Context, cancel context.CancelFunc, userId primitive.ObjectID, cassetteId primitive.ObjectID) (cassettes []models.FragmentModel, message string, ok bool) {
	var list []models.FragmentModel
	if fragmentsCollection == nil {
		return list, "", false
	}

	// filter := bson.M{"userId": userId, "cassetteId": cassetteId}
	// cursor, err := fragmentsCollection.Find(ctx, filter)
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.M{"userId": userId, "cassetteId": cassetteId}}},
		{{Key: "$lookup", Value: bson.M{
			"from":         "branches",
			"localField":   "_id",
			"foreignField": "fragmentId",
			"as":           "fragmentBranches",
		}}},
		{{
			Key:   "$project",
			Value: bson.M{"_id": 1, "name": 1, "description": 1, "color": 1, "counts.branch": bson.M{"$size": "$fragmentBranches"}},
		}},
	}
	cursor, err := fragmentsCollection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		defer cancel()
		return list, "Sorry! Server side error.We will fix this ASAP", false
	} else {
		if err = cursor.All(ctx, &list); err != nil {
			return list, "Sorry! Server side error.We will fix this ASAP", false
		}
		return list, "", true
	}
}

func FindFragmentById(ctx context.Context, cancel context.CancelFunc, fragmentId primitive.ObjectID) (fragment Fragment, message string, ok bool) {

	var entity Fragment
	if fragmentsCollection == nil {
		return entity, "", false
	}

	filter := bson.M{"_id": fragmentId}
	err := fragmentsCollection.FindOne(ctx, filter).Decode(&entity)
	if err != nil {
		defer cancel()
		return entity, "Sorry! Server side error.We will fix this ASAP", false
	}
	return entity, "", true
}
