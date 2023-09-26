package repository

import "github.com/MafuSora/portofolio_db/model"

type CategoriesRepository interface {
	Save(category model.Category)
	Update(category model.Category)
	Delete(categoryId int)
	FindById(categoryId int) (category model.Category, err error)
	FindAll() []model.Category
	FindCategory() []model.Category
	FindPortofolio(kategori string) []model.Category
	FindCertificate(kategori string) []model.Category
}
