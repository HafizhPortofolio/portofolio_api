package model

type Experience struct {
	Id           int    `gorm:"type:int;primary_key"`
	Instansi     string `gorm:"type:varchar(255)"`
	Posisi       string `gorm:"type:varchar(255)"`
	TanggalAwal  string `gorm:"type:varchar(255)"`
	TanggalAkhir string `gorm:"type:varchar(255)"`
	Deskripsi    string `gorm:"type:varchar(20000)"`
	Skill        string `gorm:"type:varchar(2000)"`
	UrlFoto      string `gorm:"type:varchar(255)"`
	UrlLink      string `gorm:"type:varchar(255)"`
}
