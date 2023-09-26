package service

import (
	"github.com/MafuSora/portofolio_db/data/request"
	"github.com/MafuSora/portofolio_db/data/response"
)

type ProfilesService interface {
	Create(profiles request.CreateProfilesRequest)
	Update(profiles request.UpdateProfilesRequest)
	Delete(profilesId int)
	FindById(profilesId int) response.ProfilesResponse
	FindAll() []response.ProfilesResponse
	FindFirst() []response.ProfilesResponse
}
