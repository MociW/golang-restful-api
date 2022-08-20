package exception

import (
	"net/http"

	"Mociw/golang-api/helper"
	"Mociw/golang-api/model/web"

	"github.com/go-playground/validator"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if notFoundError(w, r, err) {
		return
	}
	if validationError(w, r, err) {
		return
	}
	internalServerError(w, r, err)
}

func validationError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		webRespone := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}

		helper.WriteToResponeBody(w, webRespone)
		return true
	} else {
		return false
	}
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		webRespone := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   exception.Error,
		}

		helper.WriteToResponeBody(w, webRespone)
		return true
	} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	webRespone := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "internal Server Error",
		Data:   err,
	}

	helper.WriteToResponeBody(w, webRespone)
}
