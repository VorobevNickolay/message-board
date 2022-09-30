package app

import "encoding/json"

type ValidationErrors struct {
	Errors map[string]string `json:"errors"`
}

func NewValidationErrors() ValidationErrors {
	return ValidationErrors{
		Errors: make(map[string]string, 0),
	}
}

func (ve ValidationErrors) Error() string {
	res, _ := json.Marshal(ve)
	return string(res)
}
