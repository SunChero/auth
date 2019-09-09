package handler

import (
	"log"
	"net/http"
)

func (h *handler) sendMagicLink(res http.ResponseWriter, req *http.Request) {
	email := req.URL.Query().Get(":email")
	log.Println(`received this email on param`, email)
	//send email
}
