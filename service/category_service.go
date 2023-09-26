package service

import (
	"github.com/MafuSora/portofolio_db/data/request"
	"github.com/MafuSora/portofolio_db/data/response"
)

type CategoriesService interface {
	Create(category request.CreateCategoriesRequest)
	Update(category request.UpdateCategoriesRequest)
	Delete(categoryId int)
	FindById(categoryId int) response.CategoriesResponse
	FindAll() []response.CategoriesResponse
	FindCategory() []response.CategoriesResponse
	FindPortofolio(kategori string) []response.CategoriesResponse
	FindCertificate(kategori string) []response.CategoriesResponse
}
