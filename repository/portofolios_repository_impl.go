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
func (t *PortofoliosRepositoryImpl) FindCategory() []model.Portofolios {
	var portofolios []model.Portofolios
	result := t.Db.Distinct("kategori").Find(&portofolios)
	// result := t.Db.Find(&portofolios, kategori)
	helper.ErrorPanic(result.Error)
	return portofolios
}

// Delete implements PortofoliosRepository.
func (t *PortofoliosRepositoryImpl) Delete(portofoliosId int) {
	var portofolios model.Portofolios
	result := t.Db.Where("id=?", portofoliosId).Delete(&portofolios)
	helper.ErrorPanic(result.Error)
}

// FindAll implements PortofoliosRepository.
func (t *PortofoliosRepositoryImpl) FindAll() []model.Portofolios {
	var portofolios []model.Portofolios
	result := t.Db.Find(&portofolios)
	helper.ErrorPanic(result.Error)
	return portofolios
}

// FindBack implements PortofoliosRepository.
func (t *PortofoliosRepositoryImpl) FindBack(kategori string) []model.Portofolios {
	var portofolios []model.Portofolios
	result := t.Db.Where("kategori =?", kategori).Find(&portofolios)
	// result := t.Db.Find(&portofolios, kategori)
	helper.ErrorPanic(result.Error)
	return portofolios
}

// FindById implements PortofoliosRepository.
func (t *PortofoliosRepositoryImpl) FindById(portofoliosId int) (portofolios model.Portofolios, err error) {
	var portofolio model.Portofolios
	result := t.Db.Find(&portofolio, portofoliosId)
	if result != nil {
		return portofolio, nil
	} else {
		return portofolio, errors.New("portofolio Not Found")
	}
}

// FindData implements PortofoliosRepository.
func (t *PortofoliosRepositoryImpl) FindData(kategori string) []model.Portofolios {
	var portofolios []model.Portofolios
	result := t.Db.Where("kategori =?", kategori).Find(&portofolios)
	// result := t.Db.Find(&portofolios, kategori)
	helper.ErrorPanic(result.Error)
	return portofolios
}

// FindDesign implements PortofoliosRepository.
func (t *PortofoliosRepositoryImpl) FindDesign(kategori string) []model.Portofolios {
	var portofolios []model.Portofolios
	result := t.Db.Where("kategori =?", kategori).Find(&portofolios)
	// result := t.Db.Find(&portofolios, kategori)
	helper.ErrorPanic(result.Error)
	return portofolios
}

// FindFront implements PortofoliosRepository.
func (t *PortofoliosRepositoryImpl) FindFront(kategori string) []model.Portofolios {
	var portofolios []model.Portofolios
	result := t.Db.Where("kategori =?", kategori).Find(&portofolios)
	// result := t.Db.Find(&portofolios, kategori)
	helper.ErrorPanic(result.Error)
	return portofolios
}

// FindIndustrial implements PortofoliosRepository.
func (t *PortofoliosRepositoryImpl) FindIndustrial(kategori string) []model.Portofolios {
	var portofolios []model.Portofolios
	result := t.Db.Where("kategori =?", kategori).Find(&portofolios)
	// result := t.Db.Find(&portofolios, kategori)
	helper.ErrorPanic(result.Error)
	return portofolios
}

// Save implements PortofoliosRepository.
func (t *PortofoliosRepositoryImpl) Save(portofolios model.Portofolios) {
	result := t.Db.Create(&portofolios)
	helper.ErrorPanic(result.Error)
}

// Update implements PortofoliosRepository.
func (t *PortofoliosRepositoryImpl) Update(portofolios model.Portofolios) {
	var updatePortofolio = request.UpdatePortofoliosRequest{
		Id:             portofolios.Id,
		NamaPortofolio: portofolios.NamaPortofolio,
		Deskripsi:      portofolios.Deskripsi,
		UrlGambar:      portofolios.UrlGambar,
		UrlLink:        portofolios.UrlLink,
		Kategori:       portofolios.Kategori,
	}
	result := t.Db.Model(&portofolios).Updates(updatePortofolio)
	helper.ErrorPanic(result.Error)
}
