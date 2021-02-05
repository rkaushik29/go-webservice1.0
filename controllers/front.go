package controllers

import "net/http"

func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc) // handles both paths defined
	http.Handle("/users/", *uc)
}

func hello() {}
