package models

type User struct {
	ID             int64  `json:"id,omitempty"`
	UUID           string `json:"uuid,omitempty"`
	Username       string `json:"username,omitempty"`
	Email          string `json:"em,omitempty"`
	ProfilePicture string `json:"pp,omitempty"`
	Password       string `json:"ps,omitempty"`
	InsertDate     int64  `json:"ins,omitempty"`
}
