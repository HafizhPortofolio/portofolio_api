package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/m/data/request"
	"github.com/m/data/response"
	"github.com/m/helper"
	"github.com/m/model"
	"github.com/m/repository"
)

type ProfilesServiceImpl struct {
	ProfilesRepository repository.ProfileRepository
	validate           *validator.Validate
}

func NewProfilesServiceImpl(profilesRepository repository.ProfileRepository, validate *validator.Validate) ProfilesService {
	return &ProfilesServiceImpl{
		ProfilesRepository: profilesRepository,
		validate:           validate,
	}
}

// Create implements ProfilesService.
func (t *ProfilesServiceImpl) Create(profile request.CreateProfilesRequest) {
	err := t.validate.Struct(profile)
	helper.ErrorPanic(err)
	profileModel := model.Profile{
		Nama:               profile.Nama,
		TempatLahir:        profile.TempatLahir,
		TanggalLahir:       profile.TanggalLahir,
		Alamat:             profile.Alamat,
		Email:              profile.Email,
		NoHandphone:        profile.NoHandphone,
		PendidikanTerakhir: profile.PendidikanTerakhir,
		Jurusan:            profile.Jurusan,
		Universitas:        profile.Universitas,
		UrlFotoProfil:      profile.UrlFotoProfil,
	}
	t.ProfilesRepository.Save(profileModel)
}

// Delete implements ProfilesService.
func (t *ProfilesServiceImpl) Delete(profileId int) {
	t.ProfilesRepository.Delete(profileId)
}

// FindAll implements ProfilesService.
func (t *ProfilesServiceImpl) FindAll() []response.ProfilesResponse {
	result := t.ProfilesRepository.FindAll()

	var profiles []response.ProfilesResponse
	for _, value := range result {
		profile := response.ProfilesResponse{
			Id:                 value.Id,
			Nama:               value.Nama,
			TempatLahir:        value.TempatLahir,
			TanggalLahir:       value.TanggalLahir,
			Alamat:             value.Alamat,
			Email:              value.Email,
			NoHandphone:        value.NoHandphone,
			PendidikanTerakhir: value.PendidikanTerakhir,
			Jurusan:            value.Jurusan,
			Universitas:        value.Universitas,
			UrlFotoProfil:      value.UrlFotoProfil,
		}
		profiles = append(profiles, profile)
	}

	return profiles
}

// FindById implements ProfilesService.
func (t *ProfilesServiceImpl) FindById(profileId int) response.ProfilesResponse {
	profileData, err := t.ProfilesRepository.FindById(profileId)
	helper.ErrorPanic(err)

	profileResponse := response.ProfilesResponse{
		Id:                 profileData.Id,
		Nama:               profileData.Nama,
		TempatLahir:        profileData.TempatLahir,
		TanggalLahir:       profileData.TanggalLahir,
		Alamat:             profileData.Alamat,
		Email:              profileData.Email,
		NoHandphone:        profileData.NoHandphone,
		PendidikanTerakhir: profileData.PendidikanTerakhir,
		Jurusan:            profileData.Jurusan,
		Universitas:        profileData.Universitas,
		UrlFotoProfil:      profileData.UrlFotoProfil,
	}
	return profileResponse
}

// Update implements ProfilesService.
func (t *ProfilesServiceImpl) Update(profile request.UpdateProfilesRequest) {
	profileData, err := t.ProfilesRepository.FindById(profile.Id)
	helper.ErrorPanic(err)
	profileData.Nama = profile.Nama
	profileData.TempatLahir = profile.TempatLahir
	profileData.TanggalLahir = profile.TanggalLahir
	profileData.Alamat = profile.Alamat
	profileData.Email = profile.Email
	profileData.NoHandphone = profile.NoHandphone
	profileData.PendidikanTerakhir = profile.PendidikanTerakhir
	profileData.Jurusan = profile.Jurusan
	profileData.Universitas = profile.Universitas
	profileData.UrlFotoProfil = profile.UrlFotoProfil
	t.ProfilesRepository.Update(profileData)
}
