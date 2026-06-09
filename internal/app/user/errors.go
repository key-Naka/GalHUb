package user

import "errors"

var (
	ErrInvalidCommand  = errors.New("无效请求")
	ErrUserExists      = errors.New("用户存在")
	ErrUserNotFound    = errors.New("用户不存在")
	ErrEmailExists     = errors.New("邮箱已存在")
	ErrInvalidPassword = errors.New("密码错误")
)
