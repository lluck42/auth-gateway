package admin

import "gq_volunteer_platform/model"

type AdminRecord struct {
	ID        uint `gorm:"primarykey;->"`
	AdminID   uint
	AdminName string
	Method    string
	ApiPath   string
	Ip        string
	model.BaseTime
}
