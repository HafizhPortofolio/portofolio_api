package request

type CreateCategoriesRequest struct {
	Kategori string `validate:"required,min=1,max=200" json:"kategori"`
	Kegunaan string `validate:"required,min=1,max=200" json:"kegunaan"`
}
