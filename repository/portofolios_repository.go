package repository

import "github.com/m/model"

type PortofoliosRepository interface {
	Save(portofolio model.Portofolio)
	Update(portofolio model.Portofolio)
	Delete(portofolioId int)
	FindById(portofolioId int) (portofolio model.Portofolio, err error)
	FindAll() []model.Portofolio
	FindData(kategori string) []model.Portofolio
	FindBack(kategori string) []model.Portofolio
	FindFront(kategori string) []model.Portofolio
	FindDesign(kategori string) []model.Portofolio
	FindIndustrial(kategori string) []model.Portofolio
	FindCategory() []model.Portofolio
}
