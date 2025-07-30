package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type FragmentCountModel struct {
	Branch int `json:"branch"`
}

type FragmentModel struct {
	Id          primitive.ObjectID `bson:"_id"`
	Name        string             `json:"name" validate:"min=2, max=50"`
	Description string             `json:"description" validate:"max=150"`
	Color       string             `json:"color"`
	Counts      FragmentCountModel `json:"counts"`
	Status      bool               `json:"status"`
}

type FragmentCreateRequestModel struct {
	Name        string             `json:"name" validate:"required, min=2, max=50"`
	Description string             `json:"description" validate:"max=150"`
	CassetteId  primitive.ObjectID `bson:"cassetteId"`
}

type FragmentCreateResponseModel struct {
	Meta MetaBaseModel `json:"meta"`
	Data IdBaseModel   `json:"data"`
}

type FragmentUpdateRequestModel struct {
	Id          string             `json:"id" validate:"required"`
	Name        string             `json:"name" validate:"required, min=2, max=50"`
	Description string             `json:"description" validate:"max=150"`
	CassetteId  primitive.ObjectID `json:"cassetteId"`
}

type FragmentUpdateResponseModel struct {
	Meta MetaBaseModel `json:"meta"`
}

type FragmentDeleteRequestModel struct {
	Id string `json:"id" validate:"required"`
}

type FragmentDeleteResponseModel struct {
	Meta MetaBaseModel `json:"meta"`
}

type FragmentGetAllCassetteRequestModel struct {
	CassetteId primitive.ObjectID `json:"cassetteId" query:"cassetteId" binding:"required"`
}

type FragmentGetAllCassetteResponseModel struct {
	Meta MetaBaseModel   `json:"meta"`
	Data []FragmentModel `json:"data"`
}
