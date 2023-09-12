package service

import (
	"github.com/m/data/request"
	"github.com/m/data/response"
)

type FooterService interface {
	Create(footer request.CreateFootersRequest)
	Update(footer request.UpdateFooterRequest)
	Delete(footerId int)
	FindById(footerId int) response.FooterResponse
	FindAll() []response.FooterResponse
	FindFirst() []response.FooterResponse
}
