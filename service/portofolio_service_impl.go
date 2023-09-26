package service

import (
	"github.com/MafuSora/portofolio_db/data/request"
	"github.com/MafuSora/portofolio_db/data/response"
	"github.com/MafuSora/portofolio_db/helper"
	"github.com/MafuSora/portofolio_db/model"
	"github.com/MafuSora/portofolio_db/repository"
	"github.com/go-playground/validator/v10"
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
func (t *PortofoliosServiceImpl) Create(portofolios request.CreatePortofoliosRequest) {
	err := t.validate.Struct(portofolios)
	helper.ErrorPanic(err)
	portofolioModel := model.Portofolios{

		NamaPortofolio: portofolios.NamaPortofolio,
		Deskripsi:      portofolios.Deskripsi,
		UrlGambar:      portofolios.UrlGambar,
		UrlLink:        portofolios.UrlLink,
		Kategori:       portofolios.Kategori,
	}
	t.PortofoliosRepository.Save(portofolioModel)
}

// Delete implements PortofoliosService.
func (t *PortofoliosServiceImpl) Delete(portofoliosId int) {
	t.PortofoliosRepository.Delete(portofoliosId)
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
func (t *PortofoliosServiceImpl) FindById(portofoliosId int) response.PortofoliosResponse {
	portofolioData, err := t.PortofoliosRepository.FindById(portofoliosId)
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
func (t *PortofoliosServiceImpl) Update(portofolios request.UpdatePortofoliosRequest) {
	portofolioData, err := t.PortofoliosRepository.FindById(portofolios.Id)
	helper.ErrorPanic(err)
	portofolioData.NamaPortofolio = portofolios.NamaPortofolio
	portofolioData.Deskripsi = portofolios.Deskripsi
	portofolioData.UrlGambar = portofolios.UrlGambar
	portofolioData.UrlLink = portofolios.UrlLink
	portofolioData.Kategori = portofolios.Kategori

	t.PortofoliosRepository.Update(portofolioData)
}

// FindCategory implements PortofoliosService.
func (t *PortofoliosServiceImpl) FindCategory() []response.PortofoliosResponse {
	result := t.PortofoliosRepository.FindCategory()

	var portofolios []response.PortofoliosResponse
	for _, value := range result {
		portofolio := response.PortofoliosResponse{
			// Id:             value.Id,
			// NamaPortofolio: value.NamaPortofolio,
			// Deskripsi:      value.Deskripsi,
			// UrlGambar:      value.UrlGambar,
			// UrlLink:        value.UrlLink,
			Kategori: value.Kategori,
		}
		portofolios = append(portofolios, portofolio)
	}

	return portofolios
}
