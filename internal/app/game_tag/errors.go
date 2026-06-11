package game_tag

import "errors"

var (
	ErrGameTagUpdateFailed = errors.New("更新游戏标签失败")
	ErrGameTagNotFound     = errors.New("游戏标签不存在")
)
