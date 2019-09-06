package handler

import (
	"net/http"

	"github.com/markbates/goth/gothic"
)

func (h *handler) initAuth(res http.ResponseWriter, req *http.Request) {
	guser, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		gothic.BeginAuthHandler(res, req)
	} else {
		token, _ := h.CompleteAuth(req.Context(), &guser)
		res.Header().Set("Authorization", "Bearer "+token)
	}
}

func (h *handler) completeAuth(res http.ResponseWriter, req *http.Request) {
	guser, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		res.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	token, err := h.CompleteAuth(req.Context(), &guser)
	res.Header().Set("Authorization", "Bearer "+token)
}
