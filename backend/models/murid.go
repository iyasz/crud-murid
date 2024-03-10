package models

type Murid struct {
	Id      int64  `gorm:"primaryKey" json:"id"`
	Name    string `gorm:"type:varchar(255)" json:"name"`
	Email   string `gorm:"type:varchar(255)" json:"email"`
	Nis     int64  `json:"nis"`
	ClassID int64  `json:"class_id"`
	Class   Class  `gorm:"foreignKey:ClassID"`
}
