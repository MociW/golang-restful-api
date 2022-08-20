package helper

import (
	"Mociw/golang-api/model/domain"
	"Mociw/golang-api/model/web"
)

func ToCategoryRes(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryRes []web.CategoryResponse
	for _, category := range categories {
		categoryRes = append(categoryRes, ToCategoryRes(category))
	}
	return categoryRes
}
