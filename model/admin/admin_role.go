package admin

import "gq_volunteer_platform/model"

type AdminRole struct {
	ID                uint `gorm:"primarykey;->"`
	Name              string
	AdminMenuIds      string // 菜单id json
	AdminMenuApiPaths string // 权限路径 json
	model.BaseTime
}
