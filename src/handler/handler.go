package handler

import (
	"net/http"
	"service"

	"github.com/gorilla/pat"
)

type handler struct {
	*service.Service
}

//New create a new http handler
func New(s *service.Service) http.Handler {

	h := &handler{s}
	router := pat.New()
	router.Get("/auth/email/{email}", h.sendMagicLink)
	router.Get("/auth/{provider}", h.initAuth)
	router.Get("/auth/{provider}/callback", h.completeAuth)

	//router.Post("/event/new_user", h.eNewUser)

	return router

}
