package controller

import (
	"net/http"
	"strconv"

	"Mociw/golang-api/helper"
	"Mociw/golang-api/model/web"
	"Mociw/golang-api/service"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (control *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryCreateReq := web.CategoryCreateRequest{}

	helper.ReadFromReq(r, &categoryCreateReq)

	CategoryRes := control.CategoryService.Create(r.Context(), categoryCreateReq)
	webRespones := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   CategoryRes,
	}
	helper.WriteToResponeBody(w, webRespones)
}

func (control *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryUpdateReq := web.CategoryUpdateRequest{}
	helper.ReadFromReq(r, &categoryUpdateReq)

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateReq.Id = id

	CategoryRes := control.CategoryService.Update(r.Context(), categoryUpdateReq)
	webRespones := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   CategoryRes,
	}
	helper.WriteToResponeBody(w, webRespones)
}

func (control *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	control.CategoryService.Delete(r.Context(), id)
	webRespones := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteToResponeBody(w, webRespones)
}

func (control *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryRes := control.CategoryService.FindById(r.Context(), id)
	webRespones := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryRes,
	}
	helper.WriteToResponeBody(w, webRespones)
}

func (control *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryResponses := control.CategoryService.FindAll(r.Context())
	webRespones := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}
	helper.WriteToResponeBody(w, webRespones)
}
