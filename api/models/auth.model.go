package models

import "time"

type LoginRequestModel struct {
	Email    string `json:"email" validate:"required, min=5, max=50"`
	Password string `json:"password" validate:"required, min=6, max=25"`
}

type LoginResponseDataModel struct {
	Token     string    `json:"token"`
	ExpireAt  time.Time `json:"expire_at"`
	ReturnUrl string    `json:"return_url"`
}

type LoginResponseModel struct {
	Meta MetaBaseModel          `json:"meta"`
	Data LoginResponseDataModel `json:"data"`
}

type SignupRequestModel struct {
	Email           string `json:"email" validate:"required, min=5, max=50"`
	Password        string `json:"password" validate:"required, min=6, max=25"`
	ConfirmPassword string `json:"confirmpassword" validate:"required, min=6, max=25"`
}

type SignupResponseModel struct {
	Meta MetaBaseModel `json:"meta"`
}
