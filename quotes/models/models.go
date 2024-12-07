package models

type Quotes struct {
	Quotes []Quote `json:"quotes"`
}

type Quote struct {
	ID     int    `json:"id"`
	Quote  string `json:"quote"`
	Author string `json:"author"`
}
