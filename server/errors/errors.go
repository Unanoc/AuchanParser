package errors

import (
	"encoding/json"
	"log"
)

//easyjson:json
type Error struct {
	Message string `json:"message,omitempty"`
}

func (r *Error) Error() string {
	errorBytes, err := json.Marshal(r)
	if err != nil {
		log.Println(err)
	}
	return string(errorBytes)
}

func New(msg string) error {
	return &Error{Message: msg}
}

var ProductNotFound = New("Product not found")
var ProductsNotFound = New("Products not found")
var ProdcutIsExist = New("Products exist")