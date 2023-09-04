package model

type Certificate struct {
	Id             int    `gorm:"type:int;primary_key"`
	NamaSertifikat string `gorm:"type:varchar(255)"`
	Deskripsi      string `gorm:"type:varchar(255)"`
	UrlGambar      string `gorm:"type:varchar(255)"`
	UrlLink        string `gorm:"type:varchar(255)"`
	Kategori       string `gorm:"type:varchar(255)"`
}
