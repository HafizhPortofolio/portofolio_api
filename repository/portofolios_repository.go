package repository

import "github.com/MafuSora/portofolio_db/model"

type PortofoliosRepository interface {
	Save(portofolios model.Portofolios)
	Update(portofolios model.Portofolios)
	Delete(portofoliosId int)
	FindById(portofoliosId int) (portofolio model.Portofolios, err error)
	FindAll() []model.Portofolios
	FindData(kategori string) []model.Portofolios
	FindBack(kategori string) []model.Portofolios
	FindFront(kategori string) []model.Portofolios
	FindDesign(kategori string) []model.Portofolios
	FindIndustrial(kategori string) []model.Portofolios
	FindCategory() []model.Portofolios
}
