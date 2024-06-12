package admin

import "auth-gateway/model"

type AdminRecord struct {
	ID        uint `gorm:"primarykey;->"`
	AdminID   uint
	AdminName string
	Method    string
	ApiPath   string
	Ip        string
	model.BaseTime
}
