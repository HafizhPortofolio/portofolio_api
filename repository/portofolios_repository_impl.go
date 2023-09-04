package repository

import (
	"errors"

	"github.com/m/data/request"
	"github.com/m/helper"
	"github.com/m/model"
	"gorm.io/gorm"
)

type PortofoliosRepositoryImpl struct {
	Db *gorm.DB
}

func NewPortofoliosRepositoryImpl(Db *gorm.DB) PortofoliosRepository {
	return &PortofoliosRepositoryImpl{Db: Db}
}

// FindCategory implements PortofoliosRepository.
func (t *PortofoliosRepositoryImpl) FindCategory() []model.Portofolio {
	var portofolios []model.Portofolio
	result := t.Db.Distinct("kategori").Find(&portofolios)
	// result := t.Db.Find(&portofolios, kategori)
	helper.ErrorPanic(result.Error)
	return portofolios
}

// Delete implements PortofoliosRepository.
func (t *PortofoliosRepositoryImpl) Delete(portofolioId int) {
	var portofolios model.Portofolio
	result := t.Db.Where("id=?", portofolioId).Delete(&portofolios)
	helper.ErrorPanic(result.Error)
}

// FindAll implements PortofoliosRepository.
func (t *PortofoliosRepositoryImpl) FindAll() []model.Portofolio {
	var portofolios []model.Portofolio
	result := t.Db.Find(&portofolios)
	helper.ErrorPanic(result.Error)
	return portofolios
}

// FindBack implements PortofoliosRepository.
func (t *PortofoliosRepositoryImpl) FindBack(kategori string) []model.Portofolio {
	var portofolios []model.Portofolio
	result := t.Db.Where("kategori =?", kategori).Find(&portofolios)
	// result := t.Db.Find(&portofolios, kategori)
	helper.ErrorPanic(result.Error)
	return portofolios
}

// FindById implements PortofoliosRepository.
func (t *PortofoliosRepositoryImpl) FindById(portofolioId int) (portofolio model.Portofolio, err error) {
	var portofolios model.Portofolio
	result := t.Db.Find(&portofolios, portofolioId)
	if result != nil {
		return portofolios, nil
	} else {
		return portofolios, errors.New("portofolio Not Found")
	}
}

// FindData implements PortofoliosRepository.
func (t *PortofoliosRepositoryImpl) FindData(kategori string) []model.Portofolio {
	var portofolios []model.Portofolio
	result := t.Db.Where("kategori =?", kategori).Find(&portofolios)
	// result := t.Db.Find(&portofolios, kategori)
	helper.ErrorPanic(result.Error)
	return portofolios
}

// FindDesign implements PortofoliosRepository.
func (t *PortofoliosRepositoryImpl) FindDesign(kategori string) []model.Portofolio {
	var portofolios []model.Portofolio
	result := t.Db.Where("kategori =?", kategori).Find(&portofolios)
	// result := t.Db.Find(&portofolios, kategori)
	helper.ErrorPanic(result.Error)
	return portofolios
}

// FindFront implements PortofoliosRepository.
func (t *PortofoliosRepositoryImpl) FindFront(kategori string) []model.Portofolio {
	var portofolios []model.Portofolio
	result := t.Db.Where("kategori =?", kategori).Find(&portofolios)
	// result := t.Db.Find(&portofolios, kategori)
	helper.ErrorPanic(result.Error)
	return portofolios
}

// FindIndustrial implements PortofoliosRepository.
func (t *PortofoliosRepositoryImpl) FindIndustrial(kategori string) []model.Portofolio {
	var portofolios []model.Portofolio
	result := t.Db.Where("kategori =?", kategori).Find(&portofolios)
	// result := t.Db.Find(&portofolios, kategori)
	helper.ErrorPanic(result.Error)
	return portofolios
}

// Save implements PortofoliosRepository.
func (t *PortofoliosRepositoryImpl) Save(portofolio model.Portofolio) {
	result := t.Db.Find(&portofolio)
	helper.ErrorPanic(result.Error)
}

// Update implements PortofoliosRepository.
func (t *PortofoliosRepositoryImpl) Update(portofolio model.Portofolio) {
	var updatePortofolio = request.UpdatePortofoliosRequest{
		Id:             portofolio.Id,
		NamaPortofolio: portofolio.NamaPortofolio,
		Deskripsi:      portofolio.Deskripsi,
		UrlGambar:      portofolio.UrlGambar,
		UrlLink:        portofolio.UrlLink,
		Kategori:       portofolio.Kategori,
	}
	result := t.Db.Model(&portofolio).Updates(updatePortofolio)
	helper.ErrorPanic(result.Error)
}
