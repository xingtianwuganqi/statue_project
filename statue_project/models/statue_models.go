package models

import (
	"time"
	"gorm.io/gorm"
)

type Statue_Info struct  {
	gorm.Model
	Image string `json:"image" gorm:"size:126"`
	Views int `json:"views"`
	Statue_name string `json:"statue_name" gorm:"size:64"`
	Statue_author string `json:"statue_author" gorm:"size:32"`
	Statue_desc string `json:"Statue_desc" gorm:"size:1000"`
	Statue_date time.Time
	Like_num int `json:"like_num"`
	Collec_num int `json:"collect_num"`
}

type Statue_tags struct {
	gorm.Model
	Tag_name string `json:"tag_name"`
	
}