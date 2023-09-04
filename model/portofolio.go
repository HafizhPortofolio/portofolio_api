package model

type Portofolio struct {
	Id             int    `gorm:"type:int;primary_key"`
	NamaPortofolio string `gorm:"type:varchar(255)"`
	Deskripsi      string `gorm:"type:varchar(255)"`
	UrlGambar      string `gorm:"type:varchar(255)"`
	UrlLink        string `gorm:"type:varchar(255)"`
	Kategori       string `gorm:"type:varchar(255)"`
}
