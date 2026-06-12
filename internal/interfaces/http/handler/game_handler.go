package handler

import (
	"fmt"
	gameApp "galhub/internal/app/game"
	"galhub/internal/app/game/command"
	"galhub/internal/app/game/query"

	dto "galhub/internal/interfaces/http/dto/game"
	"time"

	"github.com/gin-gonic/gin"

	"galhub/internal/pkg/response"
	"galhub/internal/pkg/util"
)

type GameHandler struct {
	service *gameApp.Service
}

func NewGameHandler(
	service *gameApp.Service,
) *GameHandler {
	return &GameHandler{
		service: service,
	}
}
func (h *GameHandler) Create(
	c *gin.Context,
) {

	var req dto.CreateGameRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		response.BadRequest(
			c,
			err.Error(),
		)

		return
	}

	var releaseDate *time.Time

	releaseDate, err := util.ParseReleaseDate(
		req.ReleaseDate,
	)

	if err != nil {
		response.BadRequest(
			c,
			"release_date 鏍煎紡閿欒锛屽簲涓?YYYY-MM-DD",
		)
		return
	}

	err = h.service.Create(
		c.Request.Context(),
		command.CreateGame{
			Title:         req.Title,
			OriginalTitle: req.OriginalTitle,

			Cover: req.Cover,

			Description: req.Description,

			ReleaseDate: releaseDate,
		},
	)

	if err != nil {
		respondGameError(c, err)
		return
	}

	response.Success(
		c,
		nil,
	)
}
func (h *GameHandler) GetByID(
	c *gin.Context,
) {

	id, err := util.ParseID(c)

	if err != nil {
		response.BadRequest(c, "鏃犳晥鐨?id")
		return
	}

	game, err := h.service.GetByID(
		c.Request.Context(),
		id,
	)

	if err != nil {
		respondGameError(c, err)
		return
	}

	response.Success(
		c,
		game,
	)
}
func (h *GameHandler) GetDetail(
	c *gin.Context,
) {

	id, err := util.ParseID(c)

	if err != nil {
		response.BadRequest(c, "鏃犳晥鐨?id")
		return
	}

	detail, err := h.service.GetDetail(
		c.Request.Context(),
		id,
	)

	if err != nil {
		respondGameError(c, err)
		return
	}

	response.Success(
		c,
		detail,
	)
}

func (h *GameHandler) List(
	c *gin.Context,
) {
	listQuery, err := parseListGameQuery(c)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	games, err := h.service.List(
		c.Request.Context(),
		listQuery,
	)

	if err != nil {
		respondGameError(c, err)
		return
	}

	response.Success(
		c,
		games,
	)
}
func (h *GameHandler) Update(
	c *gin.Context,
) {

	id, err := util.ParseID(c)

	if err != nil {
		response.BadRequest(c, "鏃犳晥鐨?id")
		return
	}

	var req dto.UpdateGameRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	releaseDate, err := util.ParseReleaseDate(
		req.ReleaseDate,
	)

	if err != nil {
		response.BadRequest(
			c,
			"release_date 鏍煎紡閿欒锛屽簲涓?YYYY-MM-DD",
		)
		return
	}

	err = h.service.Update(
		c.Request.Context(),
		command.UpdateGame{
			ID:            id,
			Title:         req.Title,
			OriginalTitle: req.OriginalTitle,
			Cover:         req.Cover,
			Description:   req.Description,
			ReleaseDate:   releaseDate,
			Status:        req.Status,
		},
	)

	if err != nil {
		respondGameError(c, err)
		return
	}

	response.Success(
		c,
		gin.H{
			"id": id,
		},
	)
}
func (h *GameHandler) Delete(
	c *gin.Context,
) {

	id, err := util.ParseID(c)

	if err != nil {
		response.BadRequest(c, "鏃犳晥鐨?id")
		return
	}

	err = h.service.Delete(
		c.Request.Context(),
		id,
	)

	if err != nil {
		respondGameError(c, err)
		return
	}

	response.Success(
		c,
		gin.H{
			"id": id,
		},
	)
}

const maxPageSize = 100

func parsePageAndSize(
	c *gin.Context,
) (int, int, error) {
	page, size, err := util.ParsePageAndSize(
		c,
		maxPageSize,
	)
	if err != nil {
		return 0, 0, gameApp.ErrInvalidPagination
	}

	return page, size, nil
}

func parseListGameQuery(
	c *gin.Context,
) (query.ListGame, error) {
	page, size, err := parsePageAndSize(c)
	if err != nil {
		return query.ListGame{}, err
	}

	result := query.ListGame{
		Page:    page,
		Size:    size,
		Keyword: c.Query("keyword"),
	}

	tagID, err := util.ParsePositiveUintQuery(
		c,
		"tag_id",
	)
	if err != nil {
		return query.ListGame{}, fmt.Errorf("鏃犳晥鐨?tag_id")
	}
	result.TagID = tagID

	companyID, err := util.ParsePositiveUintQuery(
		c,
		"company_id",
	)
	if err != nil {
		return query.ListGame{}, fmt.Errorf("鏃犳晥鐨?company_id")
	}
	result.CompanyID = companyID

	statusValue := c.Query("status")
	if statusValue != "" {
		status, err := util.ParseGameStatus(statusValue)
		if err != nil {
			return query.ListGame{}, fmt.Errorf("鏃犳晥鐨?status")
		}

		result.Status = &status
	}

	return result, nil
}
