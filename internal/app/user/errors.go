package user

import "errors"

var (
	ErrUserExists      = errors.New("用户存在")
	ErrEmailExists     = errors.New("邮箱已存在")
	ErrInvalidPassword = errors.New("密码错误")
)
