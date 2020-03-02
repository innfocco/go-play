package model

type Message struct {
	From   string `json:"from"`
	Text   string `json:"text"`
	Readed bool   `json:"readed"`
}
