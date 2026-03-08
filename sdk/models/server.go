package models

type Server struct {
	UUID        string `json:"uuid,omitempty"`
	Name        string `json:"name,omitempty"`
	ProfilePic  string `json:"profile_pic,omitempty"`
	Description string `json:"description,omitempty"`
	InsertDate  int64  `json:"ins,omitempty"`
}
