package handler

import (
	gameApp "galhub/internal/app/game"
	"galhub/internal/app/game/command"
	"strconv"

	dto "galhub/internal/interfaces/http/dto/game"
	"time"

	"github.com/gin-gonic/gin"

	"galhub/internal/pkg/response"
	"galhub/internal/pkg/utill"
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

		response.Fail(
			c,
			err.Error(),
		)

		return
	}

	var releaseDate *time.Time

	releaseDate, err := utill.ParseReleaseDate(
		req.ReleaseDate,
	)

	if err != nil {
		response.Fail(
			c,
			"release_date format error",
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

		response.Fail(
			c,
			err.Error(),
		)

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

	id, err := utill.ParseID(c)

	if err != nil {
		response.Fail(c, "invalid id")
		return
	}

	game, err := h.service.GetByID(
		c.Request.Context(),
		id,
	)

	if err != nil {
		response.Fail(c, err.Error())
		return
	}

	response.Success(
		c,
		game,
	)
}

func (h *GameHandler) List(
	c *gin.Context,
) {

	page, _ := strconv.Atoi(
		c.DefaultQuery("page", "1"),
	)

	size, _ := strconv.Atoi(
		c.DefaultQuery("size", "10"),
	)

	games, err := h.service.List(
		c.Request.Context(),
		page,
		size,
	)

	if err != nil {
		response.Fail(c, err.Error())
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

	id, err := utill.ParseID(c)

	if err != nil {
		response.Fail(c, "invalid id")
		return
	}

	var req dto.UpdateGameRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, err.Error())
		return
	}

	releaseDate, err := utill.ParseReleaseDate(
		req.ReleaseDate,
	)

	if err != nil {
		response.Fail(
			c,
			"release_date format error",
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
		response.Fail(c, err.Error())
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

	id, err := utill.ParseID(c)

	if err != nil {
		response.Fail(c, "invalid id")
		return
	}

	err = h.service.Delete(
		c.Request.Context(),
		id,
	)

	if err != nil {
		response.Fail(c, err.Error())
		return
	}

	response.Success(
		c,
		gin.H{
			"id": id,
		},
	)
}
