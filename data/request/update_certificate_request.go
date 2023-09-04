package request

type UpdateCertificatesRequest struct {
	Id             int    `validate:"required"`
	NamaSertifikat string `validate:"required,min=1,max=200" json:"namaSertifikat"`
	Deskripsi      string `json:"deskripsi"`
	UrlGambar      string `json:"urlGambar"`
	UrlLink        string `json:"urlLink"`
	Kategori       string `json:"kategori"`
}
