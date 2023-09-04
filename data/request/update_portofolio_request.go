package request

type UpdatePortofoliosRequest struct {
	Id             int    `validate:"required"`
	NamaPortofolio string `validate:"required,min=1,max=200" json:"namaPortofolio"`
	Deskripsi      string `json:"deskripsi"`
	UrlGambar      string `json:"urlGambar"`
	UrlLink        string `json:"urlLink"`
	Kategori       string `json:"kategori"`
}
