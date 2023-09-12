package response

type ExperienceResponse struct {
	Id           int    `json:"id"`
	Instansi     string `json:"instansi"`
	Posisi       string `json:"posisi"`
	TanggalAwal  string `json:"tanggalAwal"`
	TanggalAkhir string `json:"tanggalAkhir"`
	Deskripsi    string `json:"deskripsi"`
	Skill        string `json:"skill"`
	UrlFoto      string `json:"urlFoto"`
	UrlLink      string `json:"urlLink"`
}
