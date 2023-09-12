package model

type Kegunaan struct {
	Id   int    `gorm:"type:int;primary_key"`
	Nama string `gorm:"type:varchar(255)"`
}
