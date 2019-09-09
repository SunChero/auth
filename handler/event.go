package handler

import (
	"log"
	"net/http"
)

func (h *handler) eNewUser(res http.ResponseWriter, req *http.Request) {

	log.Printf(`got request : %v`, req)
}
