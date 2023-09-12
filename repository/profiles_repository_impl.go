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

// FindFirst implements ProfileRepository.
func (t *ProfilesRepositoryImpl) FindFirst() []model.Profiles {
	var profiles []model.Profiles
	result := t.Db.First(&profiles)
	helper.ErrorPanic(result.Error)
	return profiles
}

// Delete implements ProfileRepository.
func (t *ProfilesRepositoryImpl) Delete(profilesId int) {
	var profiles model.Profiles
	result := t.Db.Where("id=?", profilesId).Delete(&profiles)
	helper.ErrorPanic(result.Error)
}

// FindAll implements ProfileRepository.
func (t *ProfilesRepositoryImpl) FindAll() []model.Profiles {
	var profiles []model.Profiles
	result := t.Db.Find(&profiles)
	helper.ErrorPanic(result.Error)
	return profiles
}

// FindById implements ProfileRepository.
func (t *ProfilesRepositoryImpl) FindById(profilesId int) (profiles model.Profiles, err error) {
	var profile model.Profiles
	result := t.Db.Find(&profile, profilesId)
	if result != nil {
		return profile, nil
	} else {
		return profile, errors.New("profile Not Found")
	}
}

// Save implements ProfileRepository.
func (t *ProfilesRepositoryImpl) Save(profiles model.Profiles) {
	result := t.Db.Create(&profiles)
	helper.ErrorPanic(result.Error)
}

// Update implements ProfileRepository.
func (t *ProfilesRepositoryImpl) Update(profiles model.Profiles) {
	var updateProfile = request.UpdateProfilesRequest{
		Id:                 profiles.Id,
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
	result := t.Db.Model(&profiles).Updates(updateProfile)
	helper.ErrorPanic(result.Error)
}
