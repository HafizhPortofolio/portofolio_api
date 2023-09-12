package request

type UpdateProfilesRequest struct {
	Id                 int    `validate:"required"`
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
	Skill              string `json:"skill"`
	Header             string `json:"header"`
	DeskripsiDiri      string `json:"deskripsiDiri"`
}
