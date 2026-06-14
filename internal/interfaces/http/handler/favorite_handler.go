package handler

import (
	"fmt"
	favoriteApp "galhub/internal/app/favorite"
	"galhub/internal/app/favorite/command"
	"galhub/internal/app/favorite/query"
	dto "galhub/internal/interfaces/http/dto/favorite"
	"galhub/internal/pkg/response"
	"galhub/internal/pkg/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FavoriteHandler struct {
	service *favoriteApp.Service
}

func NewFavoriteHandler(
	service *favoriteApp.Service,
) *FavoriteHandler {
	return &FavoriteHandler{
		service: service,
	}
}

func (h *FavoriteHandler) Create(
	c *gin.Context,
) {
	userID := c.GetUint64("user_id")
	if userID == 0 {
		response.Unauthorized(c, "未授权")
		return
	}

	var req dto.CreateFavoriteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err := h.service.Create(
		c.Request.Context(),
		command.CreateFavorite{
			UserID: userID,
			GameID: req.GameID,
		},
	)
	if err != nil {
		respondFavoriteError(c, err)
		return
	}

	response.Success(c, nil)
}

func (h *FavoriteHandler) GetByID(
	c *gin.Context,
) {
	userID := c.GetUint64("user_id")
	if userID == 0 {
		response.Unauthorized(c, "未授权")
		return
	}

	id, err := util.ParseID(c)
	if err != nil {
		response.BadRequest(c, "无效的收藏id")
		return
	}

	favorite, err := h.service.GetByID(
		c.Request.Context(),
		id,
		userID,
	)
	if err != nil {
		respondFavoriteError(c, err)
		return
	}

	response.Success(c, favorite)
}

func (h *FavoriteHandler) GetByGameID(
	c *gin.Context,
) {
	userID := c.GetUint64("user_id")
	if userID == 0 {
		response.Unauthorized(c, "未授权")
		return
	}

	gameID, err := parseFavoriteGameID(c)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	result, err := h.service.GetByGameID(
		c.Request.Context(),
		userID,
		gameID,
	)
	if err != nil {
		respondFavoriteError(c, err)
		return
	}

	response.Success(c, result)
}

func (h *FavoriteHandler) List(
	c *gin.Context,
) {
	userID := c.GetUint64("user_id")
	if userID == 0 {
		response.Unauthorized(c, "未授权")
		return
	}

	listQuery, err := parseListFavoriteQuery(
		c,
		userID,
	)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	favorites, err := h.service.List(
		c.Request.Context(),
		listQuery,
	)
	if err != nil {
		respondFavoriteError(c, err)
		return
	}

	response.Success(c, favorites)
}

func (h *FavoriteHandler) ListByUserID(
	c *gin.Context,
) {
	userID, err := util.ParseID(c)
	if err != nil {
		response.BadRequest(c, "无效的用户id")
		return
	}

	listQuery, err := parseListFavoriteQuery(
		c,
		userID,
	)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	favorites, err := h.service.List(
		c.Request.Context(),
		listQuery,
	)
	if err != nil {
		respondFavoriteError(c, err)
		return
	}

	response.Success(c, favorites)
}

func (h *FavoriteHandler) Update(
	c *gin.Context,
) {
	userID := c.GetUint64("user_id")
	if userID == 0 {
		response.Unauthorized(c, "未授权")
		return
	}

	id, err := util.ParseID(c)
	if err != nil {
		response.BadRequest(c, "无效的收藏id")
		return
	}

	var req dto.UpdateFavoriteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err = h.service.Update(
		c.Request.Context(),
		command.UpdateFavorite{
			ID:     id,
			UserID: userID,
			GameID: req.GameID,
		},
	)
	if err != nil {
		respondFavoriteError(c, err)
		return
	}

	response.Success(
		c,
		gin.H{
			"id": id,
		},
	)
}

func (h *FavoriteHandler) Delete(
	c *gin.Context,
) {
	userID := c.GetUint64("user_id")
	if userID == 0 {
		response.Unauthorized(c, "未授权")
		return
	}

	id, err := util.ParseID(c)
	if err != nil {
		response.BadRequest(c, "无效的收藏id")
		return
	}

	err = h.service.Delete(
		c.Request.Context(),
		id,
		userID,
	)
	if err != nil {
		respondFavoriteError(c, err)
		return
	}

	response.Success(
		c,
		gin.H{
			"id": id,
		},
	)
}

const maxFavoritePageSize = 100

func parseListFavoriteQuery(
	c *gin.Context,
	userID uint64,
) (query.ListFavorite, error) {
	page, size, err := util.ParsePageAndSize(
		c,
		maxFavoritePageSize,
	)
	if err != nil {
		return query.ListFavorite{}, favoriteApp.ErrInvalidPagination
	}

	return query.ListFavorite{
		UserID:  userID,
		Page:    page,
		Size:    size,
		Keyword: c.Query("keyword"),
	}, nil
}

func parseFavoriteGameID(
	c *gin.Context,
) (uint64, error) {
	gameIDValue := c.Param("game_id")
	if gameIDValue == "" {
		return 0, fmt.Errorf("无效的game_id")
	}

	gameID, err := strconv.ParseUint(
		gameIDValue,
		10,
		64,
	)
	if err != nil || gameID == 0 {
		return 0, fmt.Errorf("无效的game_id")
	}

	return gameID, nil
}
