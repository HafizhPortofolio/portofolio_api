package model

type Category struct {
	Id       int    `gorm:"type:int;primary_key"`
	Kategori string `gorm:"type:varchar(255)"`
	Kegunaan string `gorm:"type:varchar(255)"`
}
