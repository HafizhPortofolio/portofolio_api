package repository

import (
	"errors"

	"github.com/m/data/request"
	"github.com/m/helper"
	"github.com/m/model"
	"gorm.io/gorm"
)

type KegunaansRepositoryImpl struct {
	Db *gorm.DB
}

func NewKegunaansRepositoryImpl(Db *gorm.DB) KegunaansRepository {
	return &KegunaansRepositoryImpl{Db: Db}
}

// Delete implements KegunaansRepository.
func (t *KegunaansRepositoryImpl) Delete(KegunaanId int) {
	var kegunaans model.Kegunaan
	result := t.Db.Where("id=?", KegunaanId).Delete(&kegunaans)
	helper.ErrorPanic(result.Error)
}

// FindAll implements KegunaansRepository.
func (t *KegunaansRepositoryImpl) FindAll() []model.Kegunaan {
	var kegunaans []model.Kegunaan
	result := t.Db.Find(&kegunaans)
	helper.ErrorPanic(result.Error)
	return kegunaans
}

// FindById implements KegunaansRepository.
func (t *KegunaansRepositoryImpl) FindById(KegunaanId int) (kegunaan model.Kegunaan, err error) {
	var kegunaans model.Kegunaan
	result := t.Db.Find(&kegunaans, KegunaanId)

	if result != nil {
		return kegunaans, nil
	} else {
		return kegunaans, errors.New("kegunaan Not Found")
	}
}

// Save implements KegunaansRepository.
func (t *KegunaansRepositoryImpl) Save(kegunaan model.Kegunaan) {
	result := t.Db.Create(&kegunaan)
	helper.ErrorPanic(result.Error)
}

// Update implements KegunaansRepository.
func (t *KegunaansRepositoryImpl) Update(kegunaan model.Kegunaan) {
	var updateKegunaan = request.UpdateKegunaansRequest{
		Id:   kegunaan.Id,
		Nama: kegunaan.Nama,
	}
	result := t.Db.Model(&kegunaan).Updates(updateKegunaan)
	helper.ErrorPanic(result.Error)
}
