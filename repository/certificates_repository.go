package repository

import "github.com/MafuSora/portofolio_db/model"

type CertificatesRepository interface {
	Save(certificate model.Certificate)
	Update(certificate model.Certificate)
	Delete(certificateId int)
	FindById(certificateId int) (certificate model.Certificate, err error)
	FindAll() []model.Certificate
	FindSoftSkill(kategori string) []model.Certificate
	FindIndustrial(kategori string) []model.Certificate
	FindProgramming(kategori string) []model.Certificate
	FindLanguage(kategori string) []model.Certificate
	FindEvent(kategori string) []model.Certificate
	FindCategory() []model.Certificate
}
