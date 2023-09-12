package repository

import "github.com/m/model"

type FooterRepository interface {
	Save(footer model.Footer)
	Update(footer model.Footer)
	Delete(footerId int)
	FindById(footerId int) (footer model.Footer, err error)
	FindAll() []model.Footer
	FindFirst() []model.Footer
}
