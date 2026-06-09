package handler

import (
	gameApp "galhub/internal/app/game"
	"galhub/internal/app/game/command"

	dto "galhub/internal/interfaces/http/dto/game"
	"time"

	"github.com/gin-gonic/gin"

	"galhub/internal/pkg/response"
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

	if req.ReleaseDate != "" {

		t, err := time.Parse(
			"2006-01-02",
			req.ReleaseDate,
		)

		if err != nil {

			response.Fail(
				c,
				"release_date format error",
			)

			return
		}

		releaseDate = &t
	}

	err := h.service.Create(
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
