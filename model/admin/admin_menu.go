package admin

import (
	"gq_volunteer_platform/model"
)

type AdminMenu struct {
	ID       uint `gorm:"primarykey;->"`
	ParentID uint
	Name     string
	Label    string
	Path     string
	Icon     string
	Type     string
	Show     string
	Sort     uint

	model.BaseTime
}

type AdminMenuTree struct {
	ID       uint
	ParentID uint
	Name     string
	Label    string
	Path     string
	Icon     string
	Type     string
	Show     string
	Sort     uint
	Children []AdminMenuTree
}

// 生成tree
func (me AdminMenu) Tree(ParentID uint, list []AdminMenu) (tree []AdminMenuTree) {

	for _, v := range list {
		if v.ParentID == ParentID {
			var one AdminMenuTree
			one.ID = v.ID
			one.ParentID = v.ParentID
			one.Name = v.Name
			one.Label = v.Label
			one.Path = v.Path
			one.Icon = v.Icon
			one.Type = v.Type
			one.Show = v.Show
			one.Sort = v.Sort
			one.Children = me.Tree(v.ID, list)

			tree = append(tree, one)
		}
	}

	return tree
}
