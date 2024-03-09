package models

type Class struct {
	Id      int64  `gorm:"primaryKey" json:"id"`
	Name    string `gorm:"type:varchar(255)" json:"name"`
	Jurusan string `gorm:"type:varchar(255)" json:"jurusan"`
}
