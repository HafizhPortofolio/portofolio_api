package repository

import (
	"errors"

	"github.com/m/data/request"
	"github.com/m/helper"
	"github.com/m/model"
	"gorm.io/gorm"
)

type ExperiencesRepositoryImpl struct {
	Db *gorm.DB
}

func NewExperiencesRepositoryImpl(Db *gorm.DB) ExperienceRepository {
	return &ExperiencesRepositoryImpl{Db: Db}
}

// Delete implements ExperienceRepository.
func (t *ExperiencesRepositoryImpl) Delete(experienceId int) {
	var experiences model.Experience
	result := t.Db.Where("id=?", experienceId).Delete(&experiences)
	helper.ErrorPanic(result.Error)
}

// FindAll implements ExperienceRepository.
func (t *ExperiencesRepositoryImpl) FindAll() []model.Experience {
	var experiences []model.Experience
	result := t.Db.Find(&experiences)
	helper.ErrorPanic(result.Error)
	return experiences
}

// FindById implements ExperienceRepository.
func (t *ExperiencesRepositoryImpl) FindById(experienceId int) (experience model.Experience, err error) {
	var experiences model.Experience
	result := t.Db.Find(&experiences, experienceId)

	if result != nil {
		return experiences, nil
	} else {
		return experiences, errors.New("kegunaan Not Found")
	}
}

// Save implements ExperienceRepository.
func (t *ExperiencesRepositoryImpl) Save(experience model.Experience) {
	result := t.Db.Create(&experience)
	helper.ErrorPanic(result.Error)
}

// Update implements ExperienceRepository.
func (t *ExperiencesRepositoryImpl) Update(experience model.Experience) {
	var updateExperience = request.UpdateExperienceRequest{
		Id:           experience.Id,
		Instansi:     experience.Instansi,
		Posisi:       experience.Posisi,
		TanggalAwal:  experience.TanggalAwal,
		TanggalAkhir: experience.TanggalAkhir,
		Deskripsi:    experience.Deskripsi,
		Skill:        experience.Skill,
		UrlFoto:      experience.UrlFoto,
		UrlLink:      experience.UrlLink,
	}
	result := t.Db.Model(&experience).Updates(updateExperience)
	helper.ErrorPanic(result.Error)
}
