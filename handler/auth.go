package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"bitbucket.org/sunchero/auth/service"
)

type input struct {
	Email string
}

//post request
func (h *handler) createVerificationCode(w http.ResponseWriter, r *http.Request) {
	var in input
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	log.Printf(`this is the value of captured email %v`, in.Email)
	_, err := h.CreateVerificationCode(ctx, in.Email)
	if err != nil {
		respondErr(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

//get request with params
func (h *handler) loginWithCode(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	token, err := h.LoginWithCode(r.Context(), q.Get("verification_code"))
	if err == service.ErrInvalidVerificationCode || err == service.ErrInvalidRedirectURI {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err == service.ErrVerificationCodeNotFound {
		http.Error(w, err.Error(), http.StatusGone)
		return
	}

	if err == service.ErrExpiredToken {
		http.Error(w, err.Error(), http.StatusGone)
		return
	}

	if err != nil {
		respondErr(w, err)
		return
	}

	w.Header().Set("Authorization", "Bearer "+token)

}
