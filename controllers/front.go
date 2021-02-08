package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

// RegisterControllers exported
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc) // handles both paths defined
	http.Handle("/users/", *uc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
