package service

import (
	"fmt"

	"github.com/MafuSora/portofolio_db/data/request"
	"github.com/MafuSora/portofolio_db/data/response"
	"github.com/MafuSora/portofolio_db/helper"
	"github.com/MafuSora/portofolio_db/model"
	"github.com/MafuSora/portofolio_db/repository"
	"github.com/go-playground/validator/v10"
)

type CategoriesServiceImpl struct {
	CategoriesRepository repository.CategoriesRepository
	validate             *validator.Validate
}

func NewCategoriesServiceImpl(categoryRepository repository.CategoriesRepository, validate *validator.Validate) CategoriesService {
	return &CategoriesServiceImpl{
		CategoriesRepository: categoryRepository,
		validate:             validate,
	}
}

// Create implements CategoriesService.
func (t *CategoriesServiceImpl) Create(category request.CreateCategoriesRequest) {
	err := t.validate.Struct(category)
	helper.ErrorPanic(err)
	categoryModel := model.Category{
		Kategori: category.Kategori,
		Kegunaan: category.Kegunaan,
	}
	t.CategoriesRepository.Save(categoryModel)
	// fmt.Println(tagModel)
}

// Delete implements CategoriesService.
func (t *CategoriesServiceImpl) Delete(categoryId int) {
	t.CategoriesRepository.Delete(categoryId)
}

// FindAll implements CategoriesService.
func (t *CategoriesServiceImpl) FindAll() []response.CategoriesResponse {
	result := t.CategoriesRepository.FindAll()

	var categories []response.CategoriesResponse
	for _, value := range result {
		category := response.CategoriesResponse{
			Id:       value.Id,
			Kategori: value.Kategori,
			Kegunaan: value.Kegunaan,
		}
		categories = append(categories, category)
	}

	return categories
}

// FindCertificate implements CategoriesService.
func (t *CategoriesServiceImpl) FindCertificate(kategori string) []response.CategoriesResponse {
	result := t.CategoriesRepository.FindCertificate(kategori)
	fmt.Println(kategori)
	var categories []response.CategoriesResponse
	for _, value := range result {
		category := response.CategoriesResponse{

			Kategori: value.Kategori,
			Kegunaan: value.Kegunaan,
		}
		categories = append(categories, category)
	}
	return categories
}

// FindPortofolio implements CategoriesService.
func (t *CategoriesServiceImpl) FindPortofolio(kategori string) []response.CategoriesResponse {
	result := t.CategoriesRepository.FindPortofolio(kategori)

	var categories []response.CategoriesResponse
	for _, value := range result {
		category := response.CategoriesResponse{

			Kategori: value.Kategori,
			Kegunaan: value.Kegunaan,
		}
		categories = append(categories, category)
	}
	return categories
}

// FindById implements CategoriesService.
func (t *CategoriesServiceImpl) FindById(categoryId int) response.CategoriesResponse {
	categoryData, err := t.CategoriesRepository.FindById(categoryId)
	helper.ErrorPanic(err)

	categoryResponse := response.CategoriesResponse{
		Id:       categoryData.Id,
		Kategori: categoryData.Kategori,
		Kegunaan: categoryData.Kegunaan,
	}
	return categoryResponse
}

// FindCategory implements CategoriesService.
func (t *CategoriesServiceImpl) FindCategory() []response.CategoriesResponse {
	result := t.CategoriesRepository.FindCategory()

	var categories []response.CategoriesResponse
	for _, value := range result {
		category := response.CategoriesResponse{

			Kategori: value.Kategori,
			Kegunaan: value.Kegunaan,
		}
		categories = append(categories, category)
	}

	return categories
}

// Update implements CategoriesService.
func (t *CategoriesServiceImpl) Update(category request.UpdateCategoriesRequest) {
	categoryData, err := t.CategoriesRepository.FindById(category.Id)
	helper.ErrorPanic(err)
	categoryData.Kategori = category.Kategori
	categoryData.Kegunaan = category.Kegunaan
	t.CategoriesRepository.Update(categoryData)
}
