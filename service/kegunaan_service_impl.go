package service

import (
	"github.com/MafuSora/portofolio_db/data/request"
	"github.com/MafuSora/portofolio_db/data/response"
	"github.com/MafuSora/portofolio_db/helper"
	"github.com/MafuSora/portofolio_db/model"
	"github.com/MafuSora/portofolio_db/repository"
	"github.com/go-playground/validator/v10"
)

type KegunaansServiceImpl struct {
	KegunaansRepository repository.KegunaansRepository
	validate            *validator.Validate
}

func NewKegunaansServiceImpl(kegunaanRepository repository.KegunaansRepository, validate *validator.Validate) KegunaansService {
	return &KegunaansServiceImpl{
		KegunaansRepository: kegunaanRepository,
		validate:            validate,
	}
}

// Create implements KegunaansService.
func (t *KegunaansServiceImpl) Create(kegunaan request.CreateKegunaansRequest) {
	err := t.validate.Struct(kegunaan)
	helper.ErrorPanic(err)
	kegunaanModel := model.Kegunaan{
		Nama: kegunaan.Nama,
	}
	t.KegunaansRepository.Save(kegunaanModel)
}

// Delete implements KegunaansService.
func (t *KegunaansServiceImpl) Delete(kegunaanId int) {
	t.KegunaansRepository.Delete(kegunaanId)
}

// FindAll implements KegunaansService.
func (t *KegunaansServiceImpl) FindAll() []response.KegunaansResponse {
	result := t.KegunaansRepository.FindAll()

	var kegunaans []response.KegunaansResponse
	for _, value := range result {
		category := response.KegunaansResponse{
			Id:   value.Id,
			Nama: value.Nama,
		}
		kegunaans = append(kegunaans, category)
	}

	return kegunaans
}

// FindById implements KegunaansService.
func (t *KegunaansServiceImpl) FindById(kegunaanId int) response.KegunaansResponse {
	kegunaanData, err := t.KegunaansRepository.FindById(kegunaanId)
	helper.ErrorPanic(err)

	kegunaanResponse := response.KegunaansResponse{
		Id:   kegunaanData.Id,
		Nama: kegunaanData.Nama,
	}
	return kegunaanResponse
}

// Update implements KegunaansService.
func (t *KegunaansServiceImpl) Update(kegunaan request.UpdateKegunaansRequest) {
	kegunaanData, err := t.KegunaansRepository.FindById(kegunaan.Id)
	helper.ErrorPanic(err)
	kegunaanData.Nama = kegunaan.Nama
	t.KegunaansRepository.Update(kegunaanData)
}
