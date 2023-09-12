package request

type UpdateCategoriesRequest struct {
	Id       int    `validate:"required"`
	Kategori string `validate:"required,min=1,max=200" json:"kategori"`
	Kegunaan string `validate:"required,min=1,max=200" json:"kegunaan"`
}
