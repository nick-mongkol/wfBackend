package models

type ChatMessage struct {
	User    string `json:"user"`
	Message string `json:"message"`
}
