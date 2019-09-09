package handler

import (
	"net/http"

	"bitbucket.org/sunchero/auth/service"
	"github.com/gorilla/pat"
)

type handler struct {
	*service.Service
}

//New create a new http handler
func New(s *service.Service) http.Handler {

	h := &handler{s}
	router := pat.New()
	//auth
	router.Post("/auth/create_verification_code", h.createVerificationCode)
	router.Post("/auth/send_verification_code", h.sendVerificationCode)
	router.Get("/auth/login_with_code", h.loginWithCode)
	//oauth
	router.Get("/auth/{provider}", h.initAuth)
	router.Get("/auth/{provider}/callback", h.completeAuth)

	//router.Post("/event/new_user", h.eNewUser)

	return router

}
