package controllers

import "net/http"

// RegisterControllers exported
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc) // handles both paths defined
	http.Handle("/users/", *uc)
}
