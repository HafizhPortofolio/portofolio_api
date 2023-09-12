package service

import (
	"github.com/m/data/request"
	"github.com/m/data/response"
)

type KegunaansService interface {
	Create(kegunaan request.CreateKegunaansRequest)
	Update(kegunaan request.UpdateKegunaansRequest)
	Delete(kegunaanId int)
	FindById(kegunaanId int) response.KegunaansResponse
	FindAll() []response.KegunaansResponse
}
