package models

type Server struct {
	UUID        string `json:"uuid,omitempty"`
	Name        string `json:"nm,omitempty"`
	ProfilePic  string `json:"pp,omitempty"`
	Description string `json:"ds,omitempty"`
	InsertDate  int64  `json:"ins,omitempty"`
}
