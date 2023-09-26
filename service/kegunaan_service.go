package service

import (
	"github.com/MafuSora/portofolio_db/data/request"
	"github.com/MafuSora/portofolio_db/data/response"
)

type KegunaansService interface {
	Create(kegunaan request.CreateKegunaansRequest)
	Update(kegunaan request.UpdateKegunaansRequest)
	Delete(kegunaanId int)
	FindById(kegunaanId int) response.KegunaansResponse
	FindAll() []response.KegunaansResponse
}
