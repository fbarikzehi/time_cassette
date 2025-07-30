package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type CountCassetteModel struct {
	Fragment int `json:"fragment"`
}
type PlayStatusModel struct {
	TimeId     primitive.ObjectID `json:"time_id"`
	BranchName string             `json:"branch_name"`
}

type StatusModel struct {
	Plays []PlayStatusModel `json:"plays"`
}

type CassetteModel struct {
	Id           primitive.ObjectID `bson:"_id"`
	Name         string             `json:"name" validate:"min=2, max=50"`
	Color        string             `json:"color"`
	IsPrivate    bool               `json:"is_private"`
	TotalOfTimes int                `json:"total_of_times"`
	Counts       CountCassetteModel `json:"counts"`
	Status       StatusModel        `json:"status"`
}

type CassetteCreateRequestModel struct {
	Name        string `json:"name" validate:"required, min=2, max=50"`
	Description string `json:"description" validate:"max=150"`
	IsPrivate   bool   `json:"isPrivate" default:"true"`
}

type CassetteCreateResponseModel struct {
	Meta MetaBaseModel `json:"meta"`
	Data IdBaseModel   `json:"data"`
}

type CassetteUpdateRequestModel struct {
	Id          string `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required, min=2, max=50"`
	Description string `json:"description" validate:"max=150"`
}

type CassetteUpdateResponseModel struct {
	Meta MetaBaseModel `json:"meta"`
}

type CassetteDeleteRequestModel struct {
	Id string `json:"id" validate:"required"`
}

type CassetteDeleteResponseModel struct {
	Meta MetaBaseModel `json:"meta"`
}

type CassetteGetAllResponseModel struct {
	Meta MetaBaseModel   `json:"meta"`
	Data []CassetteModel `json:"data"`
}
