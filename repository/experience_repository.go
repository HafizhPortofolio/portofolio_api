package repository

import "github.com/MafuSora/portofolio_db/model"

type ExperienceRepository interface {
	Save(experience model.Experience)
	Update(experience model.Experience)
	Delete(experienceId int)
	FindById(experienceId int) (experience model.Experience, err error)
	FindAll() []model.Experience
}
