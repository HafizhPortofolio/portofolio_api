package repository

import (
	"errors"

	"github.com/m/data/request"
	"github.com/m/helper"
	"github.com/m/model"
	"gorm.io/gorm"
)

type CertificatesRepositoryImpl struct {
	Db *gorm.DB
}

func NewCertificatesRepositoryImpl(Db *gorm.DB) CertificatesRepository {
	return &CertificatesRepositoryImpl{Db: Db}
}

// FindCategory implements CertificatesRepository.
func (t *CertificatesRepositoryImpl) FindCategory() []model.Certificate {
	var certificates []model.Certificate
	result := t.Db.Distinct("kategori").Find(&certificates)
	// result := t.Db.Find(&portofolios, kategori)
	helper.ErrorPanic(result.Error)
	return certificates
}

// Delete implements CertificatesRepository.
func (t *CertificatesRepositoryImpl) Delete(certificateId int) {
	var certificates model.Certificate
	result := t.Db.Where("id=?", certificateId).Delete(&certificates)
	helper.ErrorPanic(result.Error)
}

// FindAll implements CertificatesRepository.
func (t *CertificatesRepositoryImpl) FindAll() []model.Certificate {
	var certificates []model.Certificate
	result := t.Db.Find(&certificates)
	helper.ErrorPanic(result.Error)
	return certificates
}

// FindById implements CertificatesRepository.
func (t *CertificatesRepositoryImpl) FindById(certificateId int) (certificate model.Certificate, err error) {
	var certificates model.Certificate
	result := t.Db.Find(&certificates, certificateId)
	if result != nil {
		return certificates, nil
	} else {
		return certificates, errors.New("certificates Not Found")
	}
}

// FindEvent implements CertificatesRepository.
func (t *CertificatesRepositoryImpl) FindEvent(kategori string) []model.Certificate {
	var certificates []model.Certificate
	result := t.Db.Where("kategori =?", kategori).Find(&certificates)
	// result := t.Db.Find(&portofolios, kategori)
	helper.ErrorPanic(result.Error)
	return certificates
}

// FindIndustrial implements CertificatesRepository.
func (t *CertificatesRepositoryImpl) FindIndustrial(kategori string) []model.Certificate {
	var certificates []model.Certificate
	result := t.Db.Where("kategori =?", kategori).Find(&certificates)
	// result := t.Db.Find(&portofolios, kategori)
	helper.ErrorPanic(result.Error)
	return certificates
}

// FindLanguage implements CertificatesRepository.
func (t *CertificatesRepositoryImpl) FindLanguage(kategori string) []model.Certificate {
	var certificates []model.Certificate
	result := t.Db.Where("kategori =?", kategori).Find(&certificates)
	// result := t.Db.Find(&portofolios, kategori)
	helper.ErrorPanic(result.Error)
	return certificates
}

// FindProgramming implements CertificatesRepository.
func (t *CertificatesRepositoryImpl) FindProgramming(kategori string) []model.Certificate {
	var certificates []model.Certificate
	result := t.Db.Where("kategori =?", kategori).Find(&certificates)
	// result := t.Db.Find(&portofolios, kategori)
	helper.ErrorPanic(result.Error)
	return certificates
}

// FindSoftSkill implements CertificatesRepository.
func (t *CertificatesRepositoryImpl) FindSoftSkill(kategori string) []model.Certificate {
	var certificates []model.Certificate
	result := t.Db.Where("kategori =?", kategori).Find(&certificates)
	// result := t.Db.Find(&portofolios, kategori)
	helper.ErrorPanic(result.Error)
	return certificates
}

// Save implements CertificatesRepository.
func (t *CertificatesRepositoryImpl) Save(certificate model.Certificate) {
	result := t.Db.Find(&certificate)
	helper.ErrorPanic(result.Error)
}

// Update implements CertificatesRepository.
func (t *CertificatesRepositoryImpl) Update(certificate model.Certificate) {
	var updateCertificate = request.UpdateCertificatesRequest{
		Id:             certificate.Id,
		NamaSertifikat: certificate.NamaSertifikat,
		Deskripsi:      certificate.Deskripsi,
		UrlGambar:      certificate.UrlGambar,
		UrlLink:        certificate.UrlLink,
		Kategori:       certificate.Kategori,
	}
	result := t.Db.Model(&certificate).Updates(updateCertificate)
	helper.ErrorPanic(result.Error)
}
