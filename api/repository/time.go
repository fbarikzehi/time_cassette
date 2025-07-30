package repository

import (
	"context"
	"strconv"
	"time"
	"timecassette_api/models"
	"timecassette_api/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Time struct {
	ID             primitive.ObjectID `bson: "_id"`
	Name           string             `bson:"name" validate:"min=2, max=50"`
	Description    string             `bson:"description"`
	Color          string             `bson:"color"`
	Duration       uint32             `bson:"duration"`
	StartDateTime  time.Time          `bson:"startDateTime"`
	EndDateTime    *time.Time         `bson:"endDateTime"`
	UserId         primitive.ObjectID `bson:"userId"`
	BranchId       primitive.ObjectID `bson:"branchId"`
	CreateDateTime time.Time          `bson:"create_DateTime"`
	UpdateDateTime time.Time          `bson:"update_DateTime"`
}

var timesCollection *mongo.Collection = OpenCollection(Client, "times")

func CreateTimeInstant(data models.TimeCreateRequestModel, ctx context.Context, cancel context.CancelFunc, userId primitive.ObjectID) (createdID string, message string, ok bool) {
	if timesCollection == nil {
		return primitive.NilObjectID.String(), "", false
	}

	runningTimes, ok := FindRunningTimes(ctx, cancel, userId)

	if ok {
		var m models.TimeEndRequestModel
		m.EndDateTime = data.EndDateTime
		for _, t := range runningTimes {
			m.Id = t.Id
			UpdateEndTime(ctx, cancel, userId, m)
		}
	}

	var newTime Time

	newTime.Name = "time " + strconv.FormatInt(time.Now().Unix(), 10)
	newTime.Description = data.Description
	newTime.Color = utils.RandomColorString(6)
	newTime.Duration = data.Duration
	newTime.StartDateTime = data.StartDateTime
	if data.Duration == 0 {
		newTime.EndDateTime = nil
	} else {
		var endDateTime = data.StartDateTime.Add(time.Minute * time.Duration(data.Duration))
		newTime.EndDateTime = &endDateTime
	}
	newTime.UserId = userId
	newTime.BranchId = data.BranchId
	newTime.CreateDateTime = time.Now()
	insertOneResult, err := timesCollection.InsertOne(ctx, newTime)
	if err != nil || insertOneResult.InsertedID == nil {
		defer cancel()
		return primitive.NilObjectID.String(), "Sorry! Server side error.We will fix this ASAP", false
	} else {
		oid, _ := insertOneResult.InsertedID.(primitive.ObjectID)
		return oid.Hex(), ("Time created"), true
	}
}

func UpdateTime(data models.TimeUpdateRequestModel, ctx context.Context, cancel context.CancelFunc, userId primitive.ObjectID) (message string, ok bool) {
	if timesCollection == nil {
		return "", false
	}

	var entity Time
	filter := bson.M{"_id": data.Id, "userId": userId}
	err := timesCollection.FindOne(ctx, filter).Decode(&entity)
	if err != nil {
		defer cancel()
		return "Sorry! Server side error.We will fix this ASAP", false
	}
	if entity.EndDateTime != nil {
		defer cancel()
		return "this time has started or ended already", false
	}

	update := bson.M{"$set": bson.M{"description": data.Description, "startDateTime": data.StartDateTime, "duration": data.Duration, "update_DateTime": time.Now()}}
	result, err := timesCollection.UpdateOne(ctx, filter, update)
	if err != nil || result.MatchedCount == 0 {
		defer cancel()
		return "Sorry! Server side error.We will fix this ASAP", false

	} else {
		return "Time started", true
	}
}

func UpdateTimeDescription(data models.TimeUpdateDescriptionRequestModel, ctx context.Context, cancel context.CancelFunc, userId primitive.ObjectID) (message string, ok bool) {
	if timesCollection == nil {
		return "", false
	}

	var entity Time
	filter := bson.M{"_id": data.Id, "userId": userId}
	err := timesCollection.FindOne(ctx, filter).Decode(&entity)
	if err != nil {
		defer cancel()
		return "Sorry! Server side error.We will fix this ASAP", false
	}

	update := bson.M{"$set": bson.M{"description": data.Description, "update_DateTime": time.Now()}}
	result, err := timesCollection.UpdateOne(ctx, filter, update)
	if err != nil || result.MatchedCount == 0 {
		defer cancel()
		return "Sorry! Server side error.We will fix this ASAP", false

	} else {
		return "Time description updated", true
	}
}

func UpdateStartTime(data models.TimeStartRequestModel, ctx context.Context, cancel context.CancelFunc, userId primitive.ObjectID) (message string, ok bool) {
	if timesCollection == nil {
		return "", false
	}

	var entity Time
	filter := bson.M{"_id": data.Id, "userId": userId}
	err := timesCollection.FindOne(ctx, filter).Decode(&entity)
	if err != nil {
		defer cancel()
		return "Sorry! Server side error.We will fix this ASAP", false
	}
	if entity.EndDateTime != nil {
		defer cancel()
		return "this time has started or ended already", false
	}
	if entity.StartDateTime.Unix() > time.Now().Unix() {
		defer cancel()
		return "This time has a start point which is not arrived yet", false
	}
	update := bson.M{"$set": bson.M{"startDateTime": data.StartDateTime, "update_DateTime": time.Now()}}
	result, err := timesCollection.UpdateOne(ctx, filter, update)
	if err != nil || result.MatchedCount == 0 {
		defer cancel()
		return "Sorry! Server side error.We will fix this ASAP", false

	} else {
		return "Time started", true
	}
}

func UpdateEndTime(ctx context.Context, cancel context.CancelFunc, userId primitive.ObjectID, requestPayload models.TimeEndRequestModel) (message string, ok bool) {
	if timesCollection == nil {
		return "", false
	}
	var entity Time
	query := bson.M{"_id": requestPayload.Id, "userId": userId}
	err := timesCollection.FindOne(ctx, query).Decode(&entity)
	if err != nil || entity.EndDateTime != nil {
		defer cancel()
		return "this time has ended already", false
	}

	if entity.Duration > 0 && entity.StartDateTime.Add(time.Minute*time.Duration(entity.Duration)).Unix() > requestPayload.EndDateTime.Unix() {
		defer cancel()
		return "This time has a duration period which is not completed yet", false
	}
	update := bson.M{"$set": bson.M{"endDateTime": requestPayload.EndDateTime, "update_DateTime": requestPayload.EndDateTime}}
	result, err := timesCollection.UpdateOne(ctx, query, update)
	if err != nil || result.MatchedCount == 0 {
		defer cancel()
		return "Sorry! Server side error.We will fix this ASAP", false

	} else {
		return "Time ended", true
	}

}

func DeleteTime(data models.TimeDeleteRequestModel, ctx context.Context, cancel context.CancelFunc, userId primitive.ObjectID) (message string, ok bool) {
	if timesCollection == nil {
		return "", false
	}

	filter := bson.M{"_id": data.Id, "userId": userId}
	result, err := timesCollection.DeleteOne(ctx, filter)
	if err != nil || result.DeletedCount == 0 {
		defer cancel()
		return "Sorry! Server side error.We will fix this ASAP", false
	} else {
		return "Time Deleted", true
	}

}

func DeleteAllTimesByBranch(ctx context.Context, cancel context.CancelFunc, userId primitive.ObjectID, branchId primitive.ObjectID) (message string, ok bool) {
	if timesCollection == nil {
		return "", false
	}

	filter := bson.M{"branchId": branchId, "userId": userId}
	result, err := timesCollection.DeleteMany(ctx, filter)
	if err != nil || result.DeletedCount == 0 {
		defer cancel()
		return "Sorry! Server side error.We will fix this ASAP", false
	} else {
		return "All Times Deleted", true
	}

}

func FindTimesByBranch(ctx context.Context, cancel context.CancelFunc, userId primitive.ObjectID, branchId primitive.ObjectID) (times []models.TimeModel, message string, ok bool) {

	var branchTimes []models.TimeModel
	if timesCollection == nil {
		return branchTimes, "", false
	}

	filter := bson.M{"branchId": branchId, "userId": userId}
	cursor, err := timesCollection.Find(ctx, filter)
	if err != nil {
		defer cancel()
		return branchTimes, "Sorry! Server side error.We will fix this ASAP", false
	} else {
		if err = cursor.All(ctx, &branchTimes); err != nil {
			return branchTimes, "Sorry! Server side error.We will fix this ASAP", false
		}

		return branchTimes, "", true
	}

}

func FindBranchRunningTimes(ctx context.Context, cancel context.CancelFunc, userId primitive.ObjectID, branchId primitive.ObjectID) (branches []models.TimeModel, ok bool) {

	var list []models.TimeModel
	if timesCollection == nil {
		return list, false
	}

	filter := bson.M{"startDateTime": bson.M{"$ne": nil}, "endDateTime": nil, "userId": userId, "branchId": branchId}
	cursor, err := timesCollection.Find(ctx, filter)
	if err != nil {
		defer cancel()
		return list, false
	} else {
		if err = cursor.All(ctx, &list); err != nil {
			return list, false
		}
		return list, true
	}

}

func FindRunningTimes(ctx context.Context, cancel context.CancelFunc, userId primitive.ObjectID) (branches []models.TimeModel, ok bool) {

	var list []models.TimeModel
	if timesCollection == nil {
		return list, false
	}

	filter := bson.M{"startDateTime": bson.M{"$ne": nil}, "endDateTime": nil, "userId": userId}
	cursor, err := timesCollection.Find(ctx, filter)
	if err != nil {
		defer cancel()
		return list, false
	} else {
		if err = cursor.All(ctx, &list); err != nil {
			return list, false
		}
		return list, true
	}

}
