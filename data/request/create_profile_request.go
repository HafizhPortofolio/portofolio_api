package request

type CreateProfilesRequest struct {
	Nama               string `validate:"required,min=1,max=200" json:"nama"`
	TempatLahir        string `validate:"required,min=1,max=200" json:"tempatLahir"`
	TanggalLahir       string `validate:"required,min=1,max=200" json:"tanggalLahir"`
	Alamat             string `validate:"required,min=1,max=200" json:"alamat"`
	Email              string `validate:"required,min=1,max=200" json:"email"`
	NoHandphone        string `validate:"required,min=1,max=200" json:"noHandphone"`
	PendidikanTerakhir string `validate:"required,min=1,max=200" json:"pendidikanTerakhir"`
	Jurusan            string `validate:"required,min=1,max=200" json:"jurusan"`
	Universitas        string `validate:"required,min=1,max=200" json:"universitas"`
	UrlFotoProfil      string `json:"urlFotoProfil"`
	Skill              string `json:"skill"`
	Header             string `json:"header"`
	DeskripsiDiri      string `json:"deskripsiDiri"`
}
