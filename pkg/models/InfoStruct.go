package models

type Info struct {
	Count int    `json:"count"`
	Pages int    `json:"pages"`
	Next  string `json:"next" gorm:"column:link"`
	Prev  string `json:"prev" gorm:"column:link"`
}
