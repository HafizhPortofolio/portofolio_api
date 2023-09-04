package response

type ProfilesResponse struct {
	Id                 int    `json:"id"`
	Nama               string `json:"nama"`
	TempatLahir        string `json:"tempatLahir"`
	TanggalLahir       string `json:"tanggalLahir"`
	Alamat             string `json:"alamat"`
	Email              string `json:"email"`
	NoHandphone        string `json:"noHandphone"`
	PendidikanTerakhir string `json:"pendidikanTerakhir"`
	Jurusan            string `json:"jurusan"`
	Universitas        string `json:"universitas"`
	UrlFotoProfil      string `json:"urlFotoProfil"`
}
