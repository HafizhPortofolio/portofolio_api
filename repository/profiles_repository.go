package repository

import "github.com/m/model"

type ProfileRepository interface {
	Save(profile model.Profile)
	Update(profile model.Profile)
	Delete(profileId int)
	FindById(profileId int) (profile model.Profile, err error)
	FindAll() []model.Profile
}
