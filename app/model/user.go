package model

import (
	"golang.org/x/crypto/bcrypt"
)

type UserApiSignUpReq struct {
	Username string `v:"required|length:5,16#账号不能为空|账号长度应当在:min到:max之间"`
	Password string `v:"required|length:5,16#请输入确认密码|密码长度应当在:min到:max之间"`
}

type UserServiceSignUpReq struct {
	Username string
	Password string
}

type UserApiLoginReq struct {
	Username string `v:"required|length:5,16#账号不能为空|账号长度应当在:min到:max之间"`
	Password string `v:"required|length:5,16#请输入确认密码|密码长度应当在:min到:max之间"`
}

type UserServiceLoginReq struct {
	Username string
	Password string
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
)

func (user *UserServiceSignUpReq) EncryptPassword() error {
	if cipher, err := EncryptPassword(user.Password); err != nil {
		return err
	} else {
		user.Password = cipher
		return nil
	}
}

func EncryptPassword(str string) (string, error) {
	if bytes, err := bcrypt.GenerateFromPassword([]byte(str), PassWordCost); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

func CheckPassword(plain string, cipher string) bool {
	return bcrypt.CompareHashAndPassword([]byte(cipher), []byte(plain)) == nil
}

type UserApiRequest struct {
	Username string `orm:"username"   json:"username"` //
	Password string `orm:"password"   json:"password"` //
	Avatar   string `orm:"avatar"     json:"avatar"`   //
}
type UserApiResponse struct {
	Id       uint64 `orm:"id,primary" json:"id"`       //
	Username string `orm:"username"   json:"username"` //
	Password string `orm:"password"   json:"password"` //
	Avatar   string `orm:"avatar"     json:"avatar"`   //
}
