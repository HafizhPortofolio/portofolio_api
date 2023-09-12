package repository

import (
	"errors"

	"github.com/m/data/request"
	"github.com/m/helper"
	"github.com/m/model"
	"gorm.io/gorm"
)

type CategoriesRepositoryImpl struct {
	Db *gorm.DB
}

func NewCategoriesRepositoryImpl(Db *gorm.DB) CategoriesRepository {
	return &CategoriesRepositoryImpl{Db: Db}
}

// Delete implements CategoriesRepository.
func (t *CategoriesRepositoryImpl) Delete(categoryId int) {
	var categories model.Category
	result := t.Db.Where("id=?", categoryId).Delete(&categories)
	helper.ErrorPanic(result.Error)
}

// FindCertificate implements CategoriesRepository.
func (t *CategoriesRepositoryImpl) FindCertificate(kategori string) []model.Category {
	var categories []model.Category
	result := t.Db.Where("kategori =?", kategori).Find(&categories)
	// result := t.Db.Find(&portofolios, kategori)
	helper.ErrorPanic(result.Error)
	return categories
}

// FindPortofolio implements CategoriesRepository.
func (t *CategoriesRepositoryImpl) FindPortofolio(kategori string) []model.Category {
	var categories []model.Category
	result := t.Db.Where("kategori =?", kategori).Find(&categories)
	// result := t.Db.Find(&portofolios, kategori)
	helper.ErrorPanic(result.Error)
	return categories
}

// FindAll implements CategoriesRepository.
func (t *CategoriesRepositoryImpl) FindAll() []model.Category {
	var categories []model.Category
	result := t.Db.Find(&categories)
	helper.ErrorPanic(result.Error)
	return categories
}

// FindById implements CategoriesRepository.
func (t *CategoriesRepositoryImpl) FindById(categoryId int) (category model.Category, err error) {
	var categories model.Category
	result := t.Db.Find(&categories, categoryId)
	if result != nil {
		return categories, nil
	} else {
		return categories, errors.New("tag Not Found")
	}
}

// FindCategory implements CategoriesRepository.
func (t *CategoriesRepositoryImpl) FindCategory() []model.Category {
	var categories []model.Category
	result := t.Db.Distinct("kategori").Find(&categories)
	// result := t.Db.Find(&portofolios, kategori)
	helper.ErrorPanic(result.Error)
	return categories
}

// Save implements CategoriesRepository.
func (t *CategoriesRepositoryImpl) Save(category model.Category) {
	result := t.Db.Create(&category)
	helper.ErrorPanic(result.Error)
}

// Update implements CategoriesRepository.
func (t *CategoriesRepositoryImpl) Update(category model.Category) {
	var updateCategory = request.UpdateCategoriesRequest{
		Id:       category.Id,
		Kategori: category.Kategori,
		Kegunaan: category.Kegunaan,
	}
	result := t.Db.Model(&category).Updates(updateCategory)
	helper.ErrorPanic(result.Error)
}
