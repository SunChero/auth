package handler

import (
	"encoding/json"
	"net/http"
)

type payload struct {
	Event struct {
		Data struct {
			New struct {
				Email string
				ID    string
			}
		}
	}
}

//post request
func (h *handler) sendVerificationCode(w http.ResponseWriter, r *http.Request) {
	var in payload
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := h.SendVerificationCode(in.Event.Data.New.ID, in.Event.Data.New.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}
	w.WriteHeader(http.StatusCreated)
}
