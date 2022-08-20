package app

type UserModel struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}
type ErrorModel struct {
	Error string `json:"error"`
}
type TokenModel struct {
	Token string `json:"token"`
}

var unknownError = ErrorModel{"Unknown error"}
