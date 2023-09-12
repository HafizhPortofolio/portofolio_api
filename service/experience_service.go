package service

import (
	"github.com/m/data/request"
	"github.com/m/data/response"
)

type ExperiencesService interface {
	Create(experience request.CreateExperienceRequest)
	Update(experience request.UpdateExperienceRequest)
	Delete(experienceId int)
	FindById(experienceId int) response.ExperienceResponse
	FindAll() []response.ExperienceResponse
}
