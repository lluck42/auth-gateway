package config

import (
	"gq_volunteer_platform/bean"
	"gq_volunteer_platform/model"
)

// 支付配置
type Config struct {
	ID      uint `gorm:"primarykey;->"`
	Project string
	Mobile  string

	model.BaseTime
}

// 获取 confg
func (conf *Config) Get() error {

	if err := bean.MysqlDb.Where("id=?", 1).Take(conf).Error; err != nil {
		return err
	}

	return nil
}
