package models

// type ResponseBaseModel struct {
// 	Meta MetaBaseModel `json:"meta"`
// 	Data any           `json:"data,omitempty"`
// }

type IdBaseModel struct {
	Id string `json:"id"`
}

type MetaBaseModel struct {
	Result   bool     `json:"result"`
	Messages []string `json:"messages"`
}
