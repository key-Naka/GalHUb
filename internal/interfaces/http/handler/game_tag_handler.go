package handler

import (
	gameTagApp "galhub/internal/app/game_tag"
	"galhub/internal/app/game_tag/command"
	dto "galhub/internal/interfaces/http/dto/game_tag"
	"galhub/internal/pkg/response"
	"galhub/internal/pkg/util"

	"github.com/gin-gonic/gin"
)

type GameTagHandler struct {
	service *gameTagApp.Service
}

func NewGameTagHandler(
	service *gameTagApp.Service,
) *GameTagHandler {
	return &GameTagHandler{
		service: service,
	}
}

func (h *GameTagHandler) ReplaceTags(
	c *gin.Context,
) {
	gameID, err := util.ParseID(c)
	if err != nil {
		response.BadRequest(c, "鏃犳晥鐨?game_id")
		return
	}

	var req dto.ReplaceGameTagsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err = h.service.ReplaceGameTags(
		c.Request.Context(),
		command.ReplaceGameTags{
			GameID: gameID,
			TagIDs: req.TagIDs,
		},
	)

	if err != nil {
		respondGameTagError(c, err)
		return
	}

	response.Success(c, nil)
}

func (h *GameTagHandler) GetTags(
	c *gin.Context,
) {
	gameID, err := util.ParseID(c)
	if err != nil {
		response.BadRequest(c, "鏃犳晥鐨?game_id")
		return
	}

	tagIDs, err := h.service.GetTagIDsByGameID(
		c.Request.Context(),
		gameID,
	)

	if err != nil {
		respondGameTagError(c, err)
		return
	}

	if tagIDs == nil {
		tagIDs = []uint64{}
	}

	response.Success(c, tagIDs)
}

