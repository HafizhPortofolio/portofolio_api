package request

type UpdateExperienceRequest struct {
	Id           int    `validate:"required"`
	Instansi     string `validate:"required,min=1,max=200" json:"instansi"`
	Posisi       string `validate:"required,min=1,max=200" json:"posisi"`
	TanggalAwal  string `json:"tanggalAwal"`
	TanggalAkhir string `json:"tanggalAkhir"`
	Deskripsi    string `json:"deskripsi"`
	Skill        string `json:"skill"`
	UrlFoto      string `json:"urlFoto"`
	UrlLink      string `json:"urlLink"`
}
