package service

import (
	"github.com/MafuSora/portofolio_db/data/request"
	"github.com/MafuSora/portofolio_db/data/response"
	"github.com/MafuSora/portofolio_db/helper"
	"github.com/MafuSora/portofolio_db/model"
	"github.com/MafuSora/portofolio_db/repository"
	"github.com/go-playground/validator/v10"
)

type FootersServiceImpl struct {
	FooterRepository repository.FooterRepository
	validate         *validator.Validate
}

func NewFootersServiceImpl(footersRepository repository.FooterRepository, validate *validator.Validate) FooterService {
	return &FootersServiceImpl{
		FooterRepository: footersRepository,
		validate:         validate,
	}
}

// FindFirst implements FooterService.
func (t *FootersServiceImpl) FindFirst() []response.FooterResponse {
	result := t.FooterRepository.FindFirst()

	var footers []response.FooterResponse
	for _, value := range result {
		footer := response.FooterResponse{
			Id:        value.Id,
			NoHP:      value.NoHP,
			Email:     value.Email,
			Facebook:  value.Facebook,
			Twitter:   value.Twitter,
			LinkedIn:  value.LinkedIn,
			Github:    value.Github,
			Instagram: value.Instagram,
			Youtube:   value.Youtube,
			UrlFoto:   value.UrlFoto,
		}
		footers = append(footers, footer)
	}

	return footers
}

// Create implements FooterService.
func (t *FootersServiceImpl) Create(footer request.CreateFootersRequest) {
	err := t.validate.Struct(footer)
	helper.ErrorPanic(err)
	footerModel := model.Footer{
		NoHP:      footer.NoHP,
		Email:     footer.Email,
		Facebook:  footer.Facebook,
		Twitter:   footer.Twitter,
		LinkedIn:  footer.LinkedIn,
		Github:    footer.Github,
		Instagram: footer.Instagram,
		Youtube:   footer.Youtube,
		UrlFoto:   footer.UrlFoto,
	}
	t.FooterRepository.Save(footerModel)
	// fmt.Println(tagModel)
}

// Delete implements FooterService.
func (t *FootersServiceImpl) Delete(footerId int) {
	t.FooterRepository.Delete(footerId)
}

// FindAll implements FooterService.
func (t *FootersServiceImpl) FindAll() []response.FooterResponse {
	result := t.FooterRepository.FindAll()

	var footers []response.FooterResponse
	for _, value := range result {
		footer := response.FooterResponse{
			Id:        value.Id,
			NoHP:      value.NoHP,
			Email:     value.Email,
			Facebook:  value.Facebook,
			Twitter:   value.Twitter,
			LinkedIn:  value.LinkedIn,
			Github:    value.Github,
			Instagram: value.Instagram,
			Youtube:   value.Youtube,
			UrlFoto:   value.UrlFoto,
		}
		footers = append(footers, footer)
	}

	return footers
}

// FindById implements FooterService.
func (t *FootersServiceImpl) FindById(footerId int) response.FooterResponse {
	footerData, err := t.FooterRepository.FindById(footerId)
	helper.ErrorPanic(err)

	footerResponse := response.FooterResponse{
		Id:        footerData.Id,
		NoHP:      footerData.NoHP,
		Email:     footerData.Email,
		Facebook:  footerData.Facebook,
		Twitter:   footerData.Twitter,
		LinkedIn:  footerData.LinkedIn,
		Github:    footerData.Github,
		Instagram: footerData.Instagram,
		Youtube:   footerData.Youtube,
		UrlFoto:   footerData.UrlFoto,
	}
	return footerResponse
}

// footerData
// Update implements FooterService.
func (t *FootersServiceImpl) Update(footer request.UpdateFooterRequest) {
	footerData, err := t.FooterRepository.FindById(footer.Id)
	helper.ErrorPanic(err)

	footerData.NoHP = footer.NoHP
	footerData.Email = footer.Email
	footerData.Facebook = footer.Facebook
	footerData.Twitter = footer.Twitter
	footerData.LinkedIn = footer.LinkedIn
	footerData.Github = footer.Github
	footerData.Instagram = footer.Instagram
	footerData.Youtube = footer.Youtube
	footerData.UrlFoto = footer.UrlFoto
	t.FooterRepository.Update(footerData)
}
