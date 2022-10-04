package rest

type ErrorModel struct {
	Error string `json:"error"`
}
type TokenModel struct {
	Token string `json:"token"`
}

var UnknownError = ErrorModel{"Unknown error"}
var AccessHeader = "X-Access-Token"
