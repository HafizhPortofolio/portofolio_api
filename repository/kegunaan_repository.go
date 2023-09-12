package repository

import "github.com/m/model"

type KegunaansRepository interface {
	Save(kegunaan model.Kegunaan)
	Update(kegunaan model.Kegunaan)
	Delete(KegunaanId int)
	FindById(KegunaanId int) (kegunaan model.Kegunaan, err error)
	FindAll() []model.Kegunaan
}
