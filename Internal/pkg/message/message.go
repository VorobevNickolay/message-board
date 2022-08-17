package message

type Message struct {
	ID     uint64 `json:"id"`
	UserId uint64 `json:"UserId"`
	Text   string `json:"Text"`
}
