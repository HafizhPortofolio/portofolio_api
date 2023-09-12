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
func (t *ProfilesServiceImpl) Create(profiles request.CreateProfilesRequest) {
	err := t.validate.Struct(profiles)
	helper.ErrorPanic(err)
	profileModel := model.Profiles{
		Nama:               profiles.Nama,
		TempatLahir:        profiles.TempatLahir,
		TanggalLahir:       profiles.TanggalLahir,
		Alamat:             profiles.Alamat,
		Email:              profiles.Email,
		NoHandphone:        profiles.NoHandphone,
		PendidikanTerakhir: profiles.PendidikanTerakhir,
		Jurusan:            profiles.Jurusan,
		Universitas:        profiles.Universitas,
		UrlFotoProfil:      profiles.UrlFotoProfil,
		Skill:              profiles.Skill,
		Header:             profiles.Header,
		DeskripsiDiri:      profiles.DeskripsiDiri,
	}
	t.ProfilesRepository.Save(profileModel)
	// fmt.Println(profileModel)
}

// Delete implements ProfilesService.
func (t *ProfilesServiceImpl) Delete(profilesId int) {
	t.ProfilesRepository.Delete(profilesId)
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
			Skill:              value.Skill,
			Header:             value.Header,
			DeskripsiDiri:      value.DeskripsiDiri,
		}
		profiles = append(profiles, profile)
	}

	return profiles
}

// FindFirst implements ProfilesService.
func (t *ProfilesServiceImpl) FindFirst() []response.ProfilesResponse {
	result := t.ProfilesRepository.FindFirst()

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
			Skill:              value.Skill,
			Header:             value.Header,
			DeskripsiDiri:      value.DeskripsiDiri,
		}
		profiles = append(profiles, profile)
	}

	return profiles
}

// FindById implements ProfilesService.
func (t *ProfilesServiceImpl) FindById(profilesId int) response.ProfilesResponse {
	profileData, err := t.ProfilesRepository.FindById(profilesId)
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
		Skill:              profileData.Skill,
		Header:             profileData.Header,
		DeskripsiDiri:      profileData.DeskripsiDiri,
	}
	return profileResponse
}

// Update implements ProfilesService.
func (t *ProfilesServiceImpl) Update(profiles request.UpdateProfilesRequest) {
	profileData, err := t.ProfilesRepository.FindById(profiles.Id)
	helper.ErrorPanic(err)
	profileData.Nama = profiles.Nama
	profileData.TempatLahir = profiles.TempatLahir
	profileData.TanggalLahir = profiles.TanggalLahir
	profileData.Alamat = profiles.Alamat
	profileData.Email = profiles.Email
	profileData.NoHandphone = profiles.NoHandphone
	profileData.PendidikanTerakhir = profiles.PendidikanTerakhir
	profileData.Jurusan = profiles.Jurusan
	profileData.Universitas = profiles.Universitas
	profileData.UrlFotoProfil = profiles.UrlFotoProfil
	profileData.Skill = profiles.Skill
	profileData.Header = profiles.Header
	profileData.DeskripsiDiri = profiles.DeskripsiDiri
	t.ProfilesRepository.Update(profileData)
}
