package model

type UserReq struct {
	Username string `json:"username,omitempty"`
	Body     string `json:"body" json:"body,omitempty"`
}

type UserResp struct {
	Username string `json:"username,omitempty"`
	Body     string `json:"body" json:"body,omitempty"`
	Code     int    `json:"code" json:"code,omitempty"`
}
