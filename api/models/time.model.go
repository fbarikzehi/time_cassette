package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TimeModel struct {
	Id            primitive.ObjectID `bson:"_id"`
	Name          string             `json:"name" validate:"min=2, max=50"`
	Color         string             `json:"color"`
	StartDateTime *time.Time         `json:"startDateTime"`
	EndDateTime   *time.Time         `json:"endDateTime"`
	BranchName    string             `json:"branchName"`
}

type TimeCreateRequestModel struct {
	Description   string             `json:"description" validate:"max=150"`
	Duration      uint32             `json:"duration"`
	StartDateTime time.Time          `json:"startDateTime"`
	EndDateTime   time.Time          `json:"endDateTime"`
	BranchId      primitive.ObjectID `bson:"branchId"`
}

type TimeCreateResponseModel struct {
	Meta MetaBaseModel `json:"meta"`
	Data IdBaseModel   `json:"data"`
}

type TimeUpdateResponseModel struct {
	Meta MetaBaseModel `json:"meta"`
}

type TimeUpdateRequestModel struct {
	Id            primitive.ObjectID `json:"id" validate:"required"`
	Description   string             `json:"description" validate:"max=150"`
	Duration      uint32             `json:"duration"`
	EndDateTime   time.Time          `json:"endDateTime"`
	StartDateTime *time.Time         `json:"startDateTime"`
	BranchId      primitive.ObjectID `bson:"branchId"`
}

type TimeUpdateDescriptionRequestModel struct {
	Id          primitive.ObjectID `json:"id" validate:"required"`
	Description string             `json:"description" validate:"max=150"`
}

type TimeUpdateDescriptionResponseModel struct {
	Meta MetaBaseModel `json:"meta"`
}

type TimeDeleteRequestModel struct {
	Id primitive.ObjectID `json:"id" validate:"required"`
}

type TimeDeleteResponseModel struct {
	Meta MetaBaseModel `json:"meta"`
}

type TimeDeleteAllRequestModel struct {
	BranchId primitive.ObjectID `bson:"branchId"`
}

type TimeDeleteAllResponseModel struct {
	Meta MetaBaseModel `json:"meta"`
}

type TimeGetAllByBranchRequestModel struct {
	BranchId primitive.ObjectID `bson:"branchId"`
}

type TimeGetAllByBranchResponseModel struct {
	Meta MetaBaseModel `json:"meta"`
	Data []TimeModel   `json:"data"`
}

type TimeStartRequestModel struct {
	Id            primitive.ObjectID `json:"id" validate:"required"`
	StartDateTime time.Time          `json:"startDateTime"`
}

type TimeStartResponseModel struct {
	Meta MetaBaseModel `json:"meta"`
}

type TimeEndRequestModel struct {
	Id          primitive.ObjectID `json:"id" validate:"required"`
	EndDateTime time.Time          `json:"endDateTime"`
}

type TimeEndResponseModel struct {
	Meta MetaBaseModel `json:"meta"`
}
