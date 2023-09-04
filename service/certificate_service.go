package service

import (
	"github.com/m/data/request"
	"github.com/m/data/response"
)

type CertificatesService interface {
	Create(certificate request.CreateCertificatesRequest)
	Update(certificate request.UpdateCertificatesRequest)
	Delete(certificateId int)
	FindById(certificateId int) response.CertificatesResponse
	FindAll() []response.CertificatesResponse
	FindSoftSkill(kategori string) []response.CertificatesResponse
	FindIndustrial(kategori string) []response.CertificatesResponse
	FindProgramming(kategori string) []response.CertificatesResponse
	FindLanguage(kategori string) []response.CertificatesResponse
	FindEvent(kategori string) []response.CertificatesResponse
	FindCategory() []response.CertificatesResponse
}
