package service

import (
	"github.com/MafuSora/portofolio_db/data/request"
	"github.com/MafuSora/portofolio_db/data/response"
)

type PortofoliosService interface {
	Create(portofolios request.CreatePortofoliosRequest)
	Update(portofolios request.UpdatePortofoliosRequest)
	Delete(portofoliosId int)
	FindById(portofoliosId int) response.PortofoliosResponse
	FindAll() []response.PortofoliosResponse
	FindData(kategori string) []response.PortofoliosResponse
	FindBack(kategori string) []response.PortofoliosResponse
	FindFront(kategori string) []response.PortofoliosResponse
	FindDesign(kategori string) []response.PortofoliosResponse
	FindIndustrial(kategori string) []response.PortofoliosResponse
	FindCategory() []response.PortofoliosResponse
}
