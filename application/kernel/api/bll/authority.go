package bll

import (
	"magiccenter/common/model"
	"magiccenter/system"
)

// GetModuleAuthGroup 查询指定Module的授权分组信息
func GetModuleAuthGroup(id string) ([]model.Group, bool) {
	groups := []model.Group{}
	modulehub := system.GetModuleHub()
	sysMod, found := modulehub.FindModule(id)
	if found {
		authGroups := sysMod.AuthGroups()
		for _, ag := range authGroups {
			g := model.Group{}
			g.ID = ag.ID()
			g.Name = ag.Name()
			g.Description = ag.Description()
			g.Type = ag.Type()

			groups = append(groups, g)
		}

		return groups, true
	}

	return groups, false
}
