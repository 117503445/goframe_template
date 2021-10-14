package service

import (
	"context"
	"errors"
	"fmt"
	"goframe_template/app/dao"
	"goframe_template/app/model"
	"golang.org/x/crypto/bcrypt"

	"github.com/gogf/gf/frame/g"
)

var User = new(userService)

type userService struct{}

func (s *userService) SignUp(ctx context.Context, r *model.UserServiceSignUpReq) error {
	// 账号唯一性数据检查
	if !s.CheckUsername(ctx, r.Username) {
		return errors.New(fmt.Sprintf("账号 %s 已经存在", r.Username))
	}
	if cipher, err := s.EncryptPassword(r.Password); err != nil {
		return err
	} else {
		r.Password = cipher
	}
	if _, err := dao.User.Ctx(ctx).Save(r); err != nil {
		return err
	}
	return nil
}

// CheckUsername 检查账号是否符合规范(目前仅检查唯一性),存在返回false,否则true
func (s *userService) CheckUsername(ctx context.Context, username string) bool {
	if i, err := dao.User.Ctx(ctx).FindCount("username", username); err != nil {
		return false
	} else {
		return i == 0
	}
}

func (s *userService) GetUserByUsernamePassword(ctx context.Context, serviceReq *model.UserServiceLoginReq) *model.User {
	user := &model.User{}

	if err := dao.User.Ctx(ctx).Where(g.Map{"username=": serviceReq.Username}).Scan(user); err != nil {
		return nil
	} else {
		if s.CheckPassword(serviceReq.Password, user.Password) {
			return user
		} else {
			return nil
		}
	}
}

func (s *userService) EncryptPassword(str string) (string, error) {
	const PassWordCost = 12 // PassWordCost 密码加密难度
	if bytes, err := bcrypt.GenerateFromPassword([]byte(str), PassWordCost); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

func (s *userService) CheckPassword(plain string, cipher string) bool {
	return bcrypt.CompareHashAndPassword([]byte(cipher), []byte(plain)) == nil
}
