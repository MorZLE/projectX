package model

import "encoding/json"

type Message struct {
	Topic  string `json:"topic" json:"topic,omitempty"`
	Body   string `json:"body" json:"body,omitempty"`
	Group  string `json:"group" json:"group,omitempty"`
	Status string `json:"status" json:"status,omitempty"`
	Error  error
}

func (u *Message) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func (u *Message) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &u)
}
