package repository

import (
	"errors"

	"github.com/MafuSora/portofolio_db/data/request"
	"github.com/MafuSora/portofolio_db/helper"
	"github.com/MafuSora/portofolio_db/model"
	"gorm.io/gorm"
)

type FootersRepositoryImpl struct {
	Db *gorm.DB
}

func NewFootersRepositoryImpl(Db *gorm.DB) FooterRepository {
	return &FootersRepositoryImpl{Db: Db}
}

// FindFirst implements FooterRepository.
func (t *FootersRepositoryImpl) FindFirst() []model.Footer {
	var footers []model.Footer
	result := t.Db.First(&footers)
	helper.ErrorPanic(result.Error)
	return footers
}

// Delete implements FooterRepository.
func (t *FootersRepositoryImpl) Delete(footerId int) {
	var footers model.Footer
	result := t.Db.Where("id=?", footerId).Delete(&footers)
	helper.ErrorPanic(result.Error)
}

// FindAll implements FooterRepository.
func (t *FootersRepositoryImpl) FindAll() []model.Footer {
	var footers []model.Footer
	result := t.Db.Find(&footers)
	helper.ErrorPanic(result.Error)
	return footers
}

// FindById implements FooterRepository.
func (t *FootersRepositoryImpl) FindById(footerId int) (footer model.Footer, err error) {
	var footers model.Footer
	result := t.Db.Find(&footers, footerId)

	if result != nil {
		return footers, nil
	} else {
		return footers, errors.New("kegunaan Not Found")
	}
}

// Save implements FooterRepository.
func (t *FootersRepositoryImpl) Save(footer model.Footer) {
	result := t.Db.Create(&footer)
	helper.ErrorPanic(result.Error)
}

// Update implements FooterRepository.
func (t *FootersRepositoryImpl) Update(footer model.Footer) {
	var updateFooter = request.UpdateFooterRequest{
		Id:        footer.Id,
		NoHP:      footer.NoHP,
		Email:     footer.Email,
		Facebook:  footer.Facebook,
		Twitter:   footer.Twitter,
		LinkedIn:  footer.LinkedIn,
		Github:    footer.Github,
		Instagram: footer.Instagram,
		Youtube:   footer.Youtube,
		UrlFoto:   footer.UrlFoto,
	}
	result := t.Db.Model(&footer).Updates(updateFooter)
	helper.ErrorPanic(result.Error)
}
