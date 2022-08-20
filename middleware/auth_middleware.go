package middleware

import (
	"net/http"

	"Mociw/golang-api/helper"
	"Mociw/golang-api/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMidlleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if "Secret" == r.Header.Get("X-API-Key") {
		middleware.Handler.ServeHTTP(w, r)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		webRespone := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}

		helper.WriteToResponeBody(w, webRespone)
	}
}
