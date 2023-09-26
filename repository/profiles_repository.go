package repository

import "github.com/MafuSora/portofolio_db/model"

type ProfileRepository interface {
	Save(profiles model.Profiles)
	Update(profiles model.Profiles)
	Delete(profilesId int)
	FindById(profilesId int) (profiles model.Profiles, err error)
	FindAll() []model.Profiles
	FindFirst() []model.Profiles
}
