package model

type Message struct {
	Topic  string `json:"topic" json:"topic,omitempty"`
	Body   string `json:"body" json:"body,omitempty"`
	Group  string `json:"group" json:"group,omitempty"`
	Status string `json:"status" json:"status,omitempty"`
	Error  error
}
