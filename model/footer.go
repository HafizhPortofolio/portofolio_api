package model

type Footer struct {
	Id        int    `gorm:"type:int;primary_key"`
	NoHP      string `gorm:"type:varchar(255)"`
	Email     string `gorm:"type:varchar(255)"`
	Facebook  string `gorm:"type:varchar(255)"`
	Twitter   string `gorm:"type:varchar(255)"`
	LinkedIn  string `gorm:"type:varchar(255)"`
	Github    string `gorm:"type:varchar(255)"`
	Instagram string `gorm:"type:varchar(255)"`
	Youtube   string `gorm:"type:varchar(255)"`
	UrlFoto   string `gorm:"type:varchar(255)"`
}
