package favorite

import "errors"

var (
	ErrFavoriteNotFound  = errors.New("收藏不存在")
	ErrFavoriteExists    = errors.New("该游戏已收藏")
	ErrInvalidPagination = errors.New("无效分页参数")
	ErrGameNotFound      = errors.New("游戏不存在")
)
