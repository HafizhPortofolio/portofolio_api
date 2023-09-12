package model

type Profiles struct {
	Id                 int    `gorm:"type:int;primary_key"`
	Nama               string `gorm:"type:varchar(255)"`
	TempatLahir        string `gorm:"type:varchar(255)"`
	TanggalLahir       string `gorm:"type:varchar(255)"`
	Alamat             string `gorm:"type:varchar(255)"`
	Email              string `gorm:"type:varchar(255)"`
	NoHandphone        string `gorm:"type:varchar(255)"`
	PendidikanTerakhir string `gorm:"type:varchar(255)"`
	Jurusan            string `gorm:"type:varchar(255)"`
	Universitas        string `gorm:"type:varchar(255)"`
	UrlFotoProfil      string `gorm:"type:varchar(255)"`
	Skill              string `gorm:"type:varchar(255)"`
	Header             string `gorm:"type:varchar(255)"`
	DeskripsiDiri      string `gorm:"type:varchar(2000)"`
}
