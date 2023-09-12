package response

type CategoriesResponse struct {
	Id       int    `json:"id"`
	Kategori string `json:"kategori"`
	Kegunaan string `json:"kegunaan"`
}
