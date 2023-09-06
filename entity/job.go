package entity

import (
	"time"

	"gorm.io/gorm"
)

type Job struct{
	gorm.Model
	StartTime time.Time `json:"startTime" gorm:"type:datetime(3)"`
	EndTime time.Time `json:"endTime" gorm:"type:datetime(3)"`
	Price int `json:"price" gorm:"type:int(10)"`
	Remark string `json:"remark" gorm:"type:text"`
}

func (g *Job) TableName() string {
	return "job"
}