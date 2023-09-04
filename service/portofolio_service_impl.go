package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/m/data/request"
	"github.com/m/data/response"
	"github.com/m/helper"
	"github.com/m/model"
	"github.com/m/repository"
)

type PortofoliosServiceImpl struct {
	PortofoliosRepository repository.PortofoliosRepository
	validate              *validator.Validate
}

func NewPortofoliosServiceImpl(portofoliosRepository repository.PortofoliosRepository, validate *validator.Validate) PortofoliosService {
	return &PortofoliosServiceImpl{
		PortofoliosRepository: portofoliosRepository,
		validate:              validate,
	}
}

// Create implements PortofoliosService.
func (t *PortofoliosServiceImpl) Create(portofolio request.CreatePortofoliosRequest) {
	err := t.validate.Struct(portofolio)
	helper.ErrorPanic(err)
	portofolioModel := model.Portofolio{

		NamaPortofolio: portofolio.NamaPortofolio,
		Deskripsi:      portofolio.Deskripsi,
		UrlGambar:      portofolio.UrlGambar,
		UrlLink:        portofolio.UrlLink,
		Kategori:       portofolio.Kategori,
	}
	t.PortofoliosRepository.Save(portofolioModel)
}

// Delete implements PortofoliosService.
func (t *PortofoliosServiceImpl) Delete(portofolioId int) {
	t.PortofoliosRepository.Delete(portofolioId)
}

// FindAll implements PortofoliosService.
func (t *PortofoliosServiceImpl) FindAll() []response.PortofoliosResponse {
	result := t.PortofoliosRepository.FindAll()

	var portofolios []response.PortofoliosResponse
	for _, value := range result {
		portofolio := response.PortofoliosResponse{
			Id:             value.Id,
			NamaPortofolio: value.NamaPortofolio,
			Deskripsi:      value.Deskripsi,
			UrlGambar:      value.UrlGambar,
			UrlLink:        value.UrlLink,
			Kategori:       value.Kategori,
		}
		portofolios = append(portofolios, portofolio)
	}

	return portofolios
}

// FindBack implements PortofoliosService.
func (t *PortofoliosServiceImpl) FindBack(kategori string) []response.PortofoliosResponse {
	result := t.PortofoliosRepository.FindBack(kategori)

	var portofolios []response.PortofoliosResponse
	for _, value := range result {
		portofolio := response.PortofoliosResponse{
			Id:             value.Id,
			NamaPortofolio: value.NamaPortofolio,
			Deskripsi:      value.Deskripsi,
			UrlGambar:      value.UrlGambar,
			UrlLink:        value.UrlLink,
			Kategori:       value.Kategori,
		}
		portofolios = append(portofolios, portofolio)
	}

	return portofolios
}

// FindById implements PortofoliosService.
func (t *PortofoliosServiceImpl) FindById(portofolioId int) response.PortofoliosResponse {
	portofolioData, err := t.PortofoliosRepository.FindById(portofolioId)
	helper.ErrorPanic(err)

	portofolioResponse := response.PortofoliosResponse{
		Id:             portofolioData.Id,
		NamaPortofolio: portofolioData.NamaPortofolio,
		Deskripsi:      portofolioData.Deskripsi,
		UrlGambar:      portofolioData.UrlGambar,
		UrlLink:        portofolioData.UrlLink,
		Kategori:       portofolioData.Kategori,
	}
	return portofolioResponse
}

// FindData implements PortofoliosService.
func (t *PortofoliosServiceImpl) FindData(kategori string) []response.PortofoliosResponse {
	result := t.PortofoliosRepository.FindData(kategori)

	var portofolios []response.PortofoliosResponse
	for _, value := range result {
		portofolio := response.PortofoliosResponse{
			Id:             value.Id,
			NamaPortofolio: value.NamaPortofolio,
			Deskripsi:      value.Deskripsi,
			UrlGambar:      value.UrlGambar,
			UrlLink:        value.UrlLink,
			Kategori:       value.Kategori,
		}
		portofolios = append(portofolios, portofolio)
	}

	return portofolios
}

// FindDesign implements PortofoliosService.
func (t *PortofoliosServiceImpl) FindDesign(kategori string) []response.PortofoliosResponse {
	result := t.PortofoliosRepository.FindDesign(kategori)

	var portofolios []response.PortofoliosResponse
	for _, value := range result {
		portofolio := response.PortofoliosResponse{
			Id:             value.Id,
			NamaPortofolio: value.NamaPortofolio,
			Deskripsi:      value.Deskripsi,
			UrlGambar:      value.UrlGambar,
			UrlLink:        value.UrlLink,
			Kategori:       value.Kategori,
		}
		portofolios = append(portofolios, portofolio)
	}

	return portofolios
}

// FindFront implements PortofoliosService.
func (t *PortofoliosServiceImpl) FindFront(kategori string) []response.PortofoliosResponse {
	result := t.PortofoliosRepository.FindFront(kategori)

	var portofolios []response.PortofoliosResponse
	for _, value := range result {
		portofolio := response.PortofoliosResponse{
			Id:             value.Id,
			NamaPortofolio: value.NamaPortofolio,
			Deskripsi:      value.Deskripsi,
			UrlGambar:      value.UrlGambar,
			UrlLink:        value.UrlLink,
			Kategori:       value.Kategori,
		}
		portofolios = append(portofolios, portofolio)
	}

	return portofolios
}

// FindIndustrial implements PortofoliosService.
func (t *PortofoliosServiceImpl) FindIndustrial(kategori string) []response.PortofoliosResponse {
	result := t.PortofoliosRepository.FindIndustrial(kategori)

	var portofolios []response.PortofoliosResponse
	for _, value := range result {
		portofolio := response.PortofoliosResponse{
			Id:             value.Id,
			NamaPortofolio: value.NamaPortofolio,
			Deskripsi:      value.Deskripsi,
			UrlGambar:      value.UrlGambar,
			UrlLink:        value.UrlLink,
			Kategori:       value.Kategori,
		}
		portofolios = append(portofolios, portofolio)
	}

	return portofolios
}

// Update implements PortofoliosService.
func (t *PortofoliosServiceImpl) Update(portofolio request.UpdatePortofoliosRequest) {
	portofolioData, err := t.PortofoliosRepository.FindById(portofolio.Id)
	helper.ErrorPanic(err)
	portofolioData.NamaPortofolio = portofolio.NamaPortofolio
	portofolioData.Deskripsi = portofolio.Deskripsi
	portofolioData.UrlGambar = portofolio.UrlGambar
	portofolioData.UrlLink = portofolio.UrlLink
	portofolioData.Kategori = portofolio.Kategori

	t.PortofoliosRepository.Update(portofolioData)
}

// FindCategory implements PortofoliosService.
func (t *PortofoliosServiceImpl) FindCategory() []response.PortofoliosResponse {
	result := t.PortofoliosRepository.FindCategory()

	var portofolios []response.PortofoliosResponse
	for _, value := range result {
		portofolio := response.PortofoliosResponse{
			Id:             value.Id,
			NamaPortofolio: value.NamaPortofolio,
			Deskripsi:      value.Deskripsi,
			UrlGambar:      value.UrlGambar,
			UrlLink:        value.UrlLink,
			Kategori:       value.Kategori,
		}
		portofolios = append(portofolios, portofolio)
	}

	return portofolios
}
