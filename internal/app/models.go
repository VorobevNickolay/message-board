package app

type ErrorModel struct {
	Error string `json:"error"`
}
type TokenModel struct {
	Token string `json:"token"`
}

var UnknownError = ErrorModel{"Unknown error"}
