package models

type Server struct {
	ID          int64  `json:"id,omitempty"`
	Owner       int64  `json:"owner,omitempty"`
	Name        string `json:"name,omitempty"`
	ProfilePic  string `json:"profile_pic,omitempty"`
	Description string `json:"description,omitempty"`
	InsertDate  int64  `json:"ins,omitempty"`
}
