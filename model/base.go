package model

type BaseTime struct {
	CreatedAt string `gorm:"->"`
	UpdatedAt string `gorm:"->"`
}
