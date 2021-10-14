package service

import (
	"github.com/gogf/gf/frame/g"
	"goframe_template/app/model"
	"reflect"
)

var UserRole = new(userRoleService)

type userRoleService struct{}

func (s *userRoleService) GetRolesByUser(user *model.User) []model.Role {
	// todo orm
	var roles []model.Role
	g.Log().Line().Debug(reflect.TypeOf(roles))

	if roleIds, err := g.DB().Model("user_role").Array("role_id", "user_id = ", user.Id); err != nil {
		g.Log().Line().Error(err)
		return nil
	} else {
		if err := g.DB().Model("role").Where("id IN (?)", roleIds).Scan(&roles); err != nil {
			g.Log().Line().Error(err)
			return nil
		}
		// g.Log().Line().Debug(roles)
		return roles
	}
}

func (s *userRoleService) HasRole(user *model.User, roleName string) bool {
	roles := s.GetRolesByUser(user)
	for _, r := range roles {
		if roleName == r.Name {
			return true
		}
	}
	return false
}
