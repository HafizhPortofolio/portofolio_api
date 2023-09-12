package repository

import "github.com/m/model"

type ExperienceRepository interface {
	Save(experience model.Experience)
	Update(experience model.Experience)
	Delete(experienceId int)
	FindById(experienceId int) (experience model.Experience, err error)
	FindAll() []model.Experience
}
