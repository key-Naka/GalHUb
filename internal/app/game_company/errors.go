package game_company

import "errors"

var (
	ErrGameCompanyUpdateFailed = errors.New("更新游戏公司失败")
	ErrGameCompanyNotFound     = errors.New("游戏公司不存在")
)
