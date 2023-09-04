package response

type PortofoliosResponse struct {
	Id             int    `json:"id"`
	NamaPortofolio string `json:"namaPortofolio"`
	Deskripsi      string `json:"deskripsi"`
	UrlGambar      string `json:"urlGambar"`
	UrlLink        string `json:"urlLink"`
	Kategori       string `json:"kategori"`
}
