package models

type Channel struct {
	CType Channeltype `json:"ty,omitempty"`
	Name  string      `json:"nm,omitempty"`
}

type Channeltype int

const (
	Text Channeltype = 1 << iota
	Voice
)
