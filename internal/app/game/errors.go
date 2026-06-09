package game

import "errors"

var (
	ErrGameNotFound      = errors.New("游戏不存在")
	ErrInvalidPagination = errors.New("无效分页参数")
)
