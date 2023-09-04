package response

type CertificatesResponse struct {
	Id             int    `json:"id"`
	NamaSertifikat string `json:"namaSertifikat"`
	Deskripsi      string `json:"deskripsi"`
	UrlGambar      string `json:"urlGambar"`
	UrlLink        string `json:"urlLink"`
	Kategori       string `json:"kategori"`
}
