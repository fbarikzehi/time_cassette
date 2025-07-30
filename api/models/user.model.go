package models

type EmailModel struct {
	Email string `json:"email"`
}

type SearchEmailRequestModel struct {
	Email string `json:"email"`
}

type SearchEmailResonseModel struct {
	Meta MetaBaseModel `json:"meta"`
	Data []EmailModel  `json:"data"`
}
