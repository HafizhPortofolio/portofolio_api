package service

import (
	"github.com/m/data/request"
	"github.com/m/data/response"
)

type ProfilesService interface {
	Create(profile request.CreateProfilesRequest)
	Update(profile request.UpdateProfilesRequest)
	Delete(profileId int)
	FindById(profileId int) response.ProfilesResponse
	FindAll() []response.ProfilesResponse
}
