package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	// read regex to figure out what to do
	userIDPattern *regexp.Regexp
	// resource request on user collection

	// manipulate users
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Controller connected..."))
}

// Constructor function
func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
