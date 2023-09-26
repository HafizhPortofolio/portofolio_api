package service

import (
	"github.com/MafuSora/portofolio_db/data/request"
	"github.com/MafuSora/portofolio_db/data/response"
)

type ExperiencesService interface {
	Create(experience request.CreateExperienceRequest)
	Update(experience request.UpdateExperienceRequest)
	Delete(experienceId int)
	FindById(experienceId int) response.ExperienceResponse
	FindAll() []response.ExperienceResponse
}
