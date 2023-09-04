package repository

import (
	"errors"

	"github.com/m/data/request"
	"github.com/m/helper"
	"github.com/m/model"
	"gorm.io/gorm"
)

type ProfilesRepositoryImpl struct {
	Db *gorm.DB
}

func NewProfilesRepositoryImpl(Db *gorm.DB) ProfileRepository {
	return &ProfilesRepositoryImpl{Db: Db}
}

// Delete implements ProfileRepository.
func (t *ProfilesRepositoryImpl) Delete(profileId int) {
	var profiles model.Profile
	result := t.Db.Where("id=?", profileId).Delete(&profiles)
	helper.ErrorPanic(result.Error)
}

// FindAll implements ProfileRepository.
func (t *ProfilesRepositoryImpl) FindAll() []model.Profile {
	var profiles []model.Profile
	result := t.Db.Find(&profiles)
	helper.ErrorPanic(result.Error)
	return profiles
}

// FindById implements ProfileRepository.
func (t *ProfilesRepositoryImpl) FindById(profileId int) (profile model.Profile, err error) {
	var profiles model.Profile
	result := t.Db.Find(&profiles, profileId)
	if result != nil {
		return profiles, nil
	} else {
		return profiles, errors.New("profile Not Found")
	}
}

// Save implements ProfileRepository.
func (t *ProfilesRepositoryImpl) Save(profile model.Profile) {
	result := t.Db.Find(&profile)
	helper.ErrorPanic(result.Error)
}

// Update implements ProfileRepository.
func (t *ProfilesRepositoryImpl) Update(profile model.Profile) {
	var updateProfile = request.UpdateProfilesRequest{
		Id:                 profile.Id,
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
	result := t.Db.Model(&profile).Updates(updateProfile)
	helper.ErrorPanic(result.Error)
}
