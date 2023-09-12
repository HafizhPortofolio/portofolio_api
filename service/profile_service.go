package service

import (
	"github.com/m/data/request"
	"github.com/m/data/response"
)

type ProfilesService interface {
	Create(profiles request.CreateProfilesRequest)
	Update(profiles request.UpdateProfilesRequest)
	Delete(profilesId int)
	FindById(profilesId int) response.ProfilesResponse
	FindAll() []response.ProfilesResponse
	FindFirst() []response.ProfilesResponse
}
