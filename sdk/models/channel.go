package models

type Channel struct {
	CType Channeltype `json:"type,omitempty"`
	Name  string      `json:"name,omitempty"`
}

type Channeltype int

const (
	Text Channeltype = 1 << iota
	Voice
)
