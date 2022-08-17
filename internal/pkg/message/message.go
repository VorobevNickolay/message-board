package message

type Message struct {
	ID     string `json:"id"`
	UserId string `json:"userId"`
	Text   string `json:"text"`
}
