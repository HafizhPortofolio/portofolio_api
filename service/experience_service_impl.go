package service

import (
	"github.com/MafuSora/portofolio_db/data/request"
	"github.com/MafuSora/portofolio_db/data/response"
	"github.com/MafuSora/portofolio_db/helper"
	"github.com/MafuSora/portofolio_db/model"
	"github.com/MafuSora/portofolio_db/repository"
	"github.com/go-playground/validator/v10"
)

type ExperiencesServiceImpl struct {
	ExperienceRepository repository.ExperienceRepository
	validate             *validator.Validate
}

func NewExperiencesServiceImpl(experiencesRepository repository.ExperienceRepository, validate *validator.Validate) ExperiencesService {
	return &ExperiencesServiceImpl{
		ExperienceRepository: experiencesRepository,
		validate:             validate,
	}
}

// Create implements ExperiencesService.
func (t *ExperiencesServiceImpl) Create(experience request.CreateExperienceRequest) {
	err := t.validate.Struct(experience)
	helper.ErrorPanic(err)
	experienceModel := model.Experience{

		Instansi:     experience.Instansi,
		Posisi:       experience.Posisi,
		TanggalAwal:  experience.TanggalAwal,
		TanggalAkhir: experience.TanggalAkhir,
		Deskripsi:    experience.Deskripsi,
		UrlFoto:      experience.UrlFoto,
		UrlLink:      experience.UrlLink,
	}
	t.ExperienceRepository.Save(experienceModel)
}

// Delete implements ExperiencesService.
func (t *ExperiencesServiceImpl) Delete(experienceId int) {
	t.ExperienceRepository.Delete(experienceId)
}

// FindAll implements ExperiencesService.
func (t *ExperiencesServiceImpl) FindAll() []response.ExperienceResponse {
	result := t.ExperienceRepository.FindAll()

	var experiences []response.ExperienceResponse
	for _, value := range result {
		experience := response.ExperienceResponse{
			Id:           value.Id,
			Instansi:     value.Instansi,
			Posisi:       value.Posisi,
			TanggalAwal:  value.TanggalAwal,
			TanggalAkhir: value.TanggalAkhir,
			Deskripsi:    value.Deskripsi,
			UrlFoto:      value.UrlFoto,
			UrlLink:      value.UrlLink,
		}
		experiences = append(experiences, experience)
	}

	return experiences
}

// FindById implements ExperiencesService.
func (t *ExperiencesServiceImpl) FindById(experienceId int) response.ExperienceResponse {
	experienceData, err := t.ExperienceRepository.FindById(experienceId)
	helper.ErrorPanic(err)

	experienceResponse := response.ExperienceResponse{
		Id:           experienceData.Id,
		Instansi:     experienceData.Instansi,
		Posisi:       experienceData.Posisi,
		TanggalAwal:  experienceData.TanggalAwal,
		TanggalAkhir: experienceData.TanggalAkhir,
		Deskripsi:    experienceData.Deskripsi,
		UrlFoto:      experienceData.UrlFoto,
		UrlLink:      experienceData.UrlLink,
	}
	return experienceResponse
}

// Update implements ExperiencesService.
func (t *ExperiencesServiceImpl) Update(experience request.UpdateExperienceRequest) {
	experienceData, err := t.ExperienceRepository.FindById(experience.Id)
	helper.ErrorPanic(err)

	experienceData.Instansi = experience.Instansi
	experienceData.Posisi = experience.Posisi
	experienceData.TanggalAkhir = experience.TanggalAkhir
	experienceData.TanggalAwal = experience.TanggalAwal
	experienceData.Deskripsi = experience.Deskripsi
	experienceData.UrlFoto = experience.UrlFoto
	experienceData.UrlLink = experience.UrlLink
	t.ExperienceRepository.Update(experienceData)

}
