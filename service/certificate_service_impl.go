package service

import (
	"github.com/MafuSora/portofolio_db/data/request"
	"github.com/MafuSora/portofolio_db/data/response"
	"github.com/MafuSora/portofolio_db/helper"
	"github.com/MafuSora/portofolio_db/model"
	"github.com/MafuSora/portofolio_db/repository"
	"github.com/go-playground/validator/v10"
)

type CertificatesServiceImpl struct {
	CertificatesRepository repository.CertificatesRepository
	validate               *validator.Validate
}

func NewCertificatesServiceImpl(certificatesRepository repository.CertificatesRepository, validate *validator.Validate) CertificatesService {
	return &CertificatesServiceImpl{
		CertificatesRepository: certificatesRepository,
		validate:               validate,
	}
}

// Create implements CertificatesService.
func (t *CertificatesServiceImpl) Create(certificate request.CreateCertificatesRequest) {
	err := t.validate.Struct(certificate)
	helper.ErrorPanic(err)
	certificateModel := model.Certificate{

		NamaSertifikat: certificate.NamaSertifikat,
		Deskripsi:      certificate.Deskripsi,
		UrlGambar:      certificate.UrlGambar,
		UrlLink:        certificate.UrlLink,
		Kategori:       certificate.Kategori,
	}
	t.CertificatesRepository.Save(certificateModel)
}

// Delete implements CertificatesService.
func (t *CertificatesServiceImpl) Delete(certificateId int) {
	t.CertificatesRepository.Delete(certificateId)
}

// FindAll implements CertificatesService.
func (t *CertificatesServiceImpl) FindAll() []response.CertificatesResponse {
	result := t.CertificatesRepository.FindAll()

	var certificates []response.CertificatesResponse
	for _, value := range result {
		certificate := response.CertificatesResponse{
			Id:             value.Id,
			NamaSertifikat: value.NamaSertifikat,
			Deskripsi:      value.Deskripsi,
			UrlGambar:      value.UrlGambar,
			UrlLink:        value.UrlLink,
			Kategori:       value.Kategori,
		}
		certificates = append(certificates, certificate)
	}

	return certificates
}

// FindById implements CertificatesService.
func (t *CertificatesServiceImpl) FindById(certificateId int) response.CertificatesResponse {
	certificateData, err := t.CertificatesRepository.FindById(certificateId)
	helper.ErrorPanic(err)

	certificateResponse := response.CertificatesResponse{
		Id:             certificateData.Id,
		NamaSertifikat: certificateData.NamaSertifikat,
		Deskripsi:      certificateData.Deskripsi,
		UrlGambar:      certificateData.UrlGambar,
		UrlLink:        certificateData.UrlLink,
		Kategori:       certificateData.Kategori,
	}
	return certificateResponse
}

// FindEvent implements CertificatesService.
func (t *CertificatesServiceImpl) FindEvent(kategori string) []response.CertificatesResponse {
	result := t.CertificatesRepository.FindEvent("Event")

	var certificates []response.CertificatesResponse
	for _, value := range result {
		certificate := response.CertificatesResponse{
			Id:             value.Id,
			NamaSertifikat: value.NamaSertifikat,
			Deskripsi:      value.Deskripsi,
			UrlGambar:      value.UrlGambar,
			UrlLink:        value.UrlLink,
			Kategori:       value.Kategori,
		}
		certificates = append(certificates, certificate)
	}

	return certificates
}

// FindIndustrial implements CertificatesService.
func (t *CertificatesServiceImpl) FindIndustrial(kategori string) []response.CertificatesResponse {
	result := t.CertificatesRepository.FindIndustrial("Industrial Engineering")

	var certificates []response.CertificatesResponse
	for _, value := range result {
		certificate := response.CertificatesResponse{
			Id:             value.Id,
			NamaSertifikat: value.NamaSertifikat,
			Deskripsi:      value.Deskripsi,
			UrlGambar:      value.UrlGambar,
			UrlLink:        value.UrlLink,
			Kategori:       value.Kategori,
		}
		certificates = append(certificates, certificate)
	}

	return certificates
}

// FindLanguage implements CertificatesService.
func (t *CertificatesServiceImpl) FindLanguage(kategori string) []response.CertificatesResponse {
	result := t.CertificatesRepository.FindLanguage("Language")

	var certificates []response.CertificatesResponse
	for _, value := range result {
		certificate := response.CertificatesResponse{
			Id:             value.Id,
			NamaSertifikat: value.NamaSertifikat,
			Deskripsi:      value.Deskripsi,
			UrlGambar:      value.UrlGambar,
			UrlLink:        value.UrlLink,
			Kategori:       value.Kategori,
		}
		certificates = append(certificates, certificate)
	}

	return certificates
}

// FindProgramming implements CertificatesService.
func (t *CertificatesServiceImpl) FindProgramming(kategori string) []response.CertificatesResponse {
	result := t.CertificatesRepository.FindProgramming("Programming")

	var certificates []response.CertificatesResponse
	for _, value := range result {
		certificate := response.CertificatesResponse{
			Id:             value.Id,
			NamaSertifikat: value.NamaSertifikat,
			Deskripsi:      value.Deskripsi,
			UrlGambar:      value.UrlGambar,
			UrlLink:        value.UrlLink,
			Kategori:       value.Kategori,
		}
		certificates = append(certificates, certificate)
	}

	return certificates
}

// FindSoftSkill implements CertificatesService.
func (t *CertificatesServiceImpl) FindSoftSkill(kategori string) []response.CertificatesResponse {
	result := t.CertificatesRepository.FindSoftSkill("Soft Skill")

	var certificates []response.CertificatesResponse
	for _, value := range result {
		certificate := response.CertificatesResponse{
			Id:             value.Id,
			NamaSertifikat: value.NamaSertifikat,
			Deskripsi:      value.Deskripsi,
			UrlGambar:      value.UrlGambar,
			UrlLink:        value.UrlLink,
			Kategori:       value.Kategori,
		}
		certificates = append(certificates, certificate)
	}

	return certificates
}

// Update implements CertificatesService.
func (t *CertificatesServiceImpl) Update(certificate request.UpdateCertificatesRequest) {
	certificateData, err := t.CertificatesRepository.FindById(certificate.Id)
	helper.ErrorPanic(err)

	certificateData.NamaSertifikat = certificate.NamaSertifikat
	certificateData.Deskripsi = certificate.Deskripsi
	certificateData.UrlGambar = certificate.UrlGambar
	certificateData.UrlLink = certificate.UrlLink
	certificateData.Kategori = certificate.Kategori

	t.CertificatesRepository.Update(certificateData)
}

// FindCategory implements CertificatesService.
func (t *CertificatesServiceImpl) FindCategory() []response.CertificatesResponse {
	result := t.CertificatesRepository.FindCategory()

	var certificates []response.CertificatesResponse
	for _, value := range result {
		certificate := response.CertificatesResponse{
			Id:             value.Id,
			NamaSertifikat: value.NamaSertifikat,
			Deskripsi:      value.Deskripsi,
			UrlGambar:      value.UrlGambar,
			UrlLink:        value.UrlLink,
			Kategori:       value.Kategori,
		}
		certificates = append(certificates, certificate)
	}

	return certificates
}
