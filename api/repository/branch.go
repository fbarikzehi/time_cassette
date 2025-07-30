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

type Branch struct {
	ID                 primitive.ObjectID `bson: "_id"`
	HandlerUserId      primitive.ObjectID `bson:"handlerUserId"`
	HandlerUserConfirm bool               `bson:"handlerConfirm"`
	FragmentId         primitive.ObjectID `bson:"fragmentId"`
	Name               string             `bson:"name" validate:"min=2, max=50"`
	Color              string             `bson:"color"`
	CreateUserId       primitive.ObjectID `bson:"createUserId"`
	CreateDateTime     time.Time          `bson:"create_DateTime"`
	UpdateDateTime     time.Time          `bson:"update_DateTime"`
}

type BranchDeleteRequest struct {
	ID             primitive.ObjectID `bson: "_id"`
	BranchId       primitive.ObjectID `bson:"branchId"`
	HandlerUserId  primitive.ObjectID `bson:"handlerUserId"`
	HandlerConfirm bool               `bson:"handlerConfirm"`
	Description    string             `bson:"description"`
	SecretCode     string             `bson:"secretCode"`
	CreateUserId   primitive.ObjectID `bson:"createUserId"`
	CreateDateTime time.Time          `bson:"create_DateTime"`
	UpdateDateTime time.Time          `bson:"update_DateTime"`
}

var branchesCollection *mongo.Collection = OpenCollection(Client, "branches")
var brancheDeleteRequestCollection *mongo.Collection = OpenCollection(Client, "branche_delete_requests")

func CreateBranch(data models.BranchCreateRequestModel, ctx context.Context, cancel context.CancelFunc, ownerUserId primitive.ObjectID, handlerUserId primitive.ObjectID) (createdID string, message string, ok bool) {
	if branchesCollection == nil {
		return primitive.NilObjectID.String(), "", false
	}

	//is fragment for this user?
	fragment, message, ok := FindFragmentById(ctx, cancel, data.FragmentId)
	if !ok {
		defer cancel()
		return "", message, false
	}
	if fragment.UserId != ownerUserId {
		defer cancel()
		return "", "Fragment not found", false
	}
	//Create
	var branch Branch
	filter := bson.M{"fragmentId": data.FragmentId, "handlerUserId": handlerUserId, "createUserId": ownerUserId}
	err := branchesCollection.FindOne(ctx, filter).Decode(&branch)
	if err == mongo.ErrNoDocuments {
		filter := bson.M{"name": data.Name, "fragmentId": data.FragmentId, "createUserId": ownerUserId}
		err := branchesCollection.FindOne(ctx, filter).Decode(&branch)
		if err == mongo.ErrNoDocuments {
			branch.Name = data.Name
			branch.Color = utils.RandomColorString(6)
			branch.FragmentId = data.FragmentId
			branch.HandlerUserId = handlerUserId
			if ownerUserId == handlerUserId {
				branch.HandlerUserConfirm = true
			} else {
				branch.HandlerUserConfirm = false
			}
			branch.CreateUserId = ownerUserId
			branch.CreateDateTime = time.Now()
			insertOneResult, err := branchesCollection.InsertOne(ctx, branch)
			if err != nil || insertOneResult.InsertedID == nil {
				defer cancel()
				return primitive.NilObjectID.String(), "Sorry! Server side error.We will fix this ASAP", false

			} else {
				oid, _ := insertOneResult.InsertedID.(primitive.ObjectID)
				return oid.Hex(), ("Branch " + data.Name + " created"), true
			}

		} else {
			defer cancel()
			return primitive.NilObjectID.String(), "A branch with this name exist in this fragment", false
		}

	} else {
		defer cancel()
		return primitive.NilObjectID.String(), "A branch for this fragment and handler exist", false
	}
}

func UpdateBranch(data models.BranchUpdateRequestModel, ctx context.Context, cancel context.CancelFunc, ownerUserId primitive.ObjectID) (message string, ok bool) {
	if branchesCollection == nil {
		return "", false
	}

	//is fragment for this user?
	fragment, message, ok := FindFragmentById(ctx, cancel, data.FragmentId)
	if !ok {
		defer cancel()
		return message, false
	}
	if fragment.UserId != ownerUserId {
		defer cancel()
		return "Fragment not found", false
	}

	//Update
	var branch Branch
	filter := bson.M{"name": data.Name, "fragmentId": data.FragmentId, "createUserId": ownerUserId, "handlerConfirm": true}
	err := branchesCollection.FindOne(ctx, filter).Decode(&branch)
	if err == mongo.ErrNoDocuments {
		id, _ := primitive.ObjectIDFromHex(data.Id)
		filter := bson.M{"_id": id, "fragmentId": data.FragmentId, "createUserId": ownerUserId, "handlerConfirm": true}
		update := bson.M{"$set": bson.M{"name": data.Name}}
		result, err := branchesCollection.UpdateOne(ctx, filter, update)
		if err != nil || result.MatchedCount == 0 {
			defer cancel()
			return "Sorry! Server side error.We will fix this ASAP", false

		} else {
			return ("Branch name updated to" + data.Name), true
		}

	} else {
		defer cancel()
		return "Branch not found or other user confirmed-branch is close to update", false
	}

}

func BranchConfirm(data models.BranchConfirmRequestModel, ctx context.Context, cancel context.CancelFunc, handlerUserId primitive.ObjectID) (message string, ok bool) {
	if branchesCollection == nil {
		return "", false
	}

	var branch Branch
	branchId, _ := primitive.ObjectIDFromHex(data.Id)
	query := bson.M{"_id": branchId, "handlerUserId": handlerUserId}
	err := branchesCollection.FindOne(ctx, query).Decode(&branch)
	if err == nil {
		if data.Confirm {
			branchUpdate := bson.M{"$set": bson.M{"handlerConfirm": true, "update_DateTime": time.Now()}}
			result, err := branchesCollection.UpdateOne(ctx, query, branchUpdate)
			if err != nil || result.MatchedCount == 0 {
				defer cancel()
				return "Sorry! Server side error.We will fix this ASAP", false

			}
			return ("Branch Confirmed"), true
		} else {
			result, err := branchesCollection.DeleteOne(ctx, query)
			if err != nil || result.DeletedCount == 0 {
				defer cancel()
				return "Sorry! Server side error.We will fix this ASAP", false
			}

			return ("Branch Deleted"), true
		}

	}
	return "Branch not found.try again", false
}

func DeleteBranchConfirmation(data models.BranchDeleteConfirmationRequestModel, ctx context.Context, cancel context.CancelFunc, handlerUserId primitive.ObjectID) (message string, ok bool) {
	if brancheDeleteRequestCollection == nil {
		return "", false
	}

	var branchDelete BranchDeleteRequest
	branchDeleteRequestId, _ := primitive.ObjectIDFromHex(data.Id)
	branchDeleteFilter := bson.M{"_id": branchDeleteRequestId, "handlerUserId": handlerUserId}
	err := brancheDeleteRequestCollection.FindOne(ctx, branchDeleteFilter).Decode(&branchDelete)
	if err == nil {
		if branchDelete.SecretCode != data.SecretCode {
			return ("SecretCode is wrong"), true
		}
		if branchDelete.HandlerConfirm {
			return ("Branch already deleted"), true
		}

		branchId, _ := primitive.ObjectIDFromHex(data.Id)
		filter := bson.M{"_id": branchId, "handlerUserId": handlerUserId}
		result, err := branchesCollection.DeleteOne(ctx, filter)
		if err != nil || result.DeletedCount == 0 {
			defer cancel()
			return "Sorry! Server side error.We will fix this ASAP", false
		} else {
			branchDeleteRequestUpdate := bson.M{"$set": bson.M{"handlerConfirm": true, "update_DateTime": time.Now()}}
			result, err := brancheDeleteRequestCollection.UpdateOne(ctx, branchDeleteFilter, branchDeleteRequestUpdate)
			if err != nil || result.MatchedCount == 0 {
				defer cancel()
				return "Sorry! Server side error.We will fix this ASAP", false

			} else {
				message, ok := DeleteAllTimesByBranch(ctx, cancel, handlerUserId, branchId)
				if !ok {
					return message, true
				}
				return ("Branch deleted"), true
			}
		}

	}
	return "Delete Branch Confirmation error.try again", false
}

func DeleteBranchRequest(data models.BranchDeleteRequestModel, ctx context.Context, cancel context.CancelFunc, ownerUserId primitive.ObjectID) (message string, ok bool) {
	if brancheDeleteRequestCollection == nil {
		return "", false
	}

	var branchDeleteRequest BranchDeleteRequest
	branchId, _ := primitive.ObjectIDFromHex(data.Id)
	filter := bson.M{"branchId": branchId, "createUserId": ownerUserId}
	err := brancheDeleteRequestCollection.FindOne(ctx, filter).Decode(&branchDeleteRequest)
	if err == mongo.ErrNoDocuments {
		branch, _, ok := FindBrancheById(ctx, cancel, ownerUserId, branchId)
		if ok {
			if branch.HandlerUserId == ownerUserId {

				branchId, _ := primitive.ObjectIDFromHex(data.Id)
				deleteQuery := bson.M{"_id": branchId, "handlerUserId": branch.HandlerUserId, "createUserId": ownerUserId}
				result, err := branchesCollection.DeleteOne(ctx, deleteQuery)
				if err != nil || result.DeletedCount == 0 {
					defer cancel()
					return "Sorry! Server side error.We will fix this ASAP", false

				} else {
					return ("Branch deleted"), true
				}

			} else {
				branchDeleteRequest.BranchId = branchId
				branchDeleteRequest.HandlerUserId = branch.HandlerUserId
				branchDeleteRequest.HandlerConfirm = false
				branchDeleteRequest.Description = data.Description
				branchDeleteRequest.SecretCode = utils.RandomAlphanumericString(10)
				branchDeleteRequest.CreateUserId = ownerUserId
				branchDeleteRequest.CreateDateTime = time.Now()
				insertOneResult, err := brancheDeleteRequestCollection.InsertOne(ctx, branchDeleteRequest)
				if err != nil || insertOneResult.InsertedID == nil {
					defer cancel()
					return "Sorry! Server side error.We will fix this ASAP", false

				} else {
					return ("Branch delete confirmation sent"), true
				}
			}

		}

		return ("Branch not found"), false
	} else {
		return ("Branch delete confirmation already sent"), true
	}

}

func FindBranchesByFragment(ctx context.Context, cancel context.CancelFunc, userId primitive.ObjectID, fragmentId primitive.ObjectID) (branches []models.BranchModel, message string, ok bool) {

	var list []models.BranchModel
	if branchesCollection == nil {
		return list, "", false
	}

	// filter := bson.M{"createUserId": userId, "fragmentId": fragmentId}
	// cursor, err := branchesCollection.Find(ctx, filter)
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.M{"createUserId": userId, "fragmentId": fragmentId}}},
		{{Key: "$lookup", Value: bson.M{
			"from":         "times",
			"localField":   "_id",
			"foreignField": "branchId",
			"as":           "brancheTimes",
		}}},
		{{
			Key:   "$project",
			Value: bson.M{"_id": 1, "name": 1, "color": 1, "counts.time": bson.M{"$size": "$brancheTimes"}},
		}},
	}
	cursor, err := branchesCollection.Aggregate(context.TODO(), pipeline)
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

func FindBrancheById(ctx context.Context, cancel context.CancelFunc, userId primitive.ObjectID, id primitive.ObjectID) (branch Branch, message string, ok bool) {

	var entity Branch
	if branchesCollection == nil {
		return entity, "", false
	}

	filter := bson.M{"_id": id, "createUserId": userId}
	err := branchesCollection.FindOne(ctx, filter).Decode(&entity)
	if err != nil {
		defer cancel()
		return entity, "Sorry! Server side error.We will fix this ASAP", false
	}
	return entity, "", true
}

func FindBrancheByIdAndHandler(ctx context.Context, cancel context.CancelFunc, handlerUserId primitive.ObjectID, id primitive.ObjectID) (branch Branch, message string, ok bool) {

	var entity Branch
	if branchesCollection == nil {
		return entity, "", false
	}

	filter := bson.M{"_id": id, "handlerUserId": handlerUserId}
	err := branchesCollection.FindOne(ctx, filter).Decode(&entity)
	if err != nil {
		defer cancel()
		return entity, "Sorry! Server side error.We will fix this ASAP", false
	}
	return entity, "", true
}
