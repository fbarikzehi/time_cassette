package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type BranchCountModel struct {
	Time int `json:"time"`
}

type BranchModel struct {
	Id    primitive.ObjectID `bson:"_id"`
	Name  string             `json:"name" validate:"min=2, max=50"`
	Color string             `json:"color"`
	// FragmentName string             `json:"fragmentName"`
	Counts   BranchCountModel `json:"counts"`
	UserName string           `json:"userName"`
	Status   bool             `json:"status"`
}

type BranchCreateRequestModel struct {
	Name             string             `json:"name" validate:"required, min=2, max=50"`
	FragmentId       primitive.ObjectID `bson:"fragmentId"`
	HandlerUserEmail string             `bson:"handlerUserEmail"`
}

type BranchCreateResponseModel struct {
	Meta MetaBaseModel `json:"meta"`
	Data IdBaseModel   `json:"data"`
}

type BranchUpdateRequestModel struct {
	Id         string             `json:"id" validate:"required"`
	Name       string             `json:"name" validate:"required, min=2, max=50"`
	FragmentId primitive.ObjectID `bson:"fragmentId"`
}

type BranchUpdateResponseModel struct {
	Meta MetaBaseModel `json:"meta"`
}

type BranchConfirmRequestModel struct {
	Id      string `json:"id" validate:"required"`
	Confirm bool   `json:"confirm"`
}

type BranchConfirmResponseModel struct {
	Meta MetaBaseModel `json:"meta"`
}

type BranchDeleteConfirmationRequestModel struct {
	Id         string `json:"id" validate:"required"`
	SecretCode string `json:"secretCode" validate:"required"`
}

type BranchDeleteConfirmationResponseModel struct {
	Meta MetaBaseModel `json:"meta"`
}

type BranchDeleteRequestModel struct {
	Id          string `json:"id" validate:"required"`
	Description string `json:"description"`
}

type BranchDeleteResponseModel struct {
	Meta MetaBaseModel `json:"meta"`
}

type BranchGetAllByFragmentRequestModel struct {
	FragmentId primitive.ObjectID `bson:"fragmentId"`
}

type BranchGetAllResponseModel struct {
	Meta MetaBaseModel `json:"meta"`
	Data []BranchModel `json:"data"`
}
