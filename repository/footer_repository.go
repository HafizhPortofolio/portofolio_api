package repository

import "github.com/MafuSora/portofolio_db/model"

type FooterRepository interface {
	Save(footer model.Footer)
	Update(footer model.Footer)
	Delete(footerId int)
	FindById(footerId int) (footer model.Footer, err error)
	FindAll() []model.Footer
	FindFirst() []model.Footer
}
