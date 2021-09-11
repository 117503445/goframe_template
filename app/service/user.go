package service

import (
	"errors"
	"fmt"
	"goframe_template/app/dao"
	"goframe_template/app/model"

	"github.com/gogf/gf/frame/g"
)

var User = new(userService)

type userService struct{}

func (s *userService) SignUp(r *model.UserServiceSignUpReq) error {
	// 账号唯一性数据检查
	if !s.CheckUsername(r.Username) {
		return errors.New(fmt.Sprintf("账号 %s 已经存在", r.Username))
	}
	if err := r.EncryptPassword(); err != nil {
		return err
	}
	if _, err := dao.User.Save(r); err != nil {
		return err
	}
	return nil
}

// 检查账号是否符合规范(目前仅检查唯一性),存在返回false,否则true
func (s *userService) CheckUsername(username string) bool {
	if i, err := dao.User.FindCount("username", username); err != nil {
		return false
	} else {
		return i == 0
	}
}

func (s *userService) GetUserByUsernamePassword(serviceReq *model.UserServiceLoginReq) *model.User {
	user := &model.User{}

	if err := dao.User.Where(g.Map{"username=": serviceReq.Username}).Struct(user); err != nil {
		// g.Log().Line().Error(err)
		return nil
	} else {
		if model.CheckPassword(serviceReq.Password, user.Password) {
			return user
		} else {
			return nil
		}
	}
}
