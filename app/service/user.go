package service

import (
	"errors"
	"fmt"
	"goframe_learn/app/dao"
	"goframe_learn/app/model"
)


var User = new(userService)

type userService struct{}

// @summary 用户注册
// @tags    user
// @accept json
// @produce json
// @Param b body model.UserServiceSignUpReq true "UserServiceSignUpReq"
// @router  /api/user/sign-up [POST]
// @success 200 {object} response.JsonResponse
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
