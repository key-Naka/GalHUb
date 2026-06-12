package handler

import (
	gameCompanyApp "galhub/internal/app/game_company"
	"galhub/internal/app/game_company/command"
	dto "galhub/internal/interfaces/http/dto/game_company"
	"galhub/internal/pkg/response"
	"galhub/internal/pkg/util"

	"github.com/gin-gonic/gin"
)

type GameCompanyHandler struct {
	service *gameCompanyApp.Service
}

func NewGameCompanyHandler(
	service *gameCompanyApp.Service,
) *GameCompanyHandler {
	return &GameCompanyHandler{
		service: service,
	}
}

func (h *GameCompanyHandler) ReplaceCompanies(
	c *gin.Context,
) {
	gameID, err := util.ParseID(c)
	if err != nil {
		response.BadRequest(c, "鏃犳晥鐨?game_id")
		return
	}

	var req dto.ReplaceGameCompaniesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err = h.service.ReplaceGameCompanies(
		c.Request.Context(),
		command.ReplaceGameCompanies{
			GameID:     gameID,
			CompanyIDs: req.CompanyIDs,
		},
	)

	if err != nil {
		respondGameCompanyError(c, err)
		return
	}

	response.Success(c, nil)
}

func (h *GameCompanyHandler) GetCompanies(
	c *gin.Context,
) {
	gameID, err := util.ParseID(c)
	if err != nil {
		response.BadRequest(c, "鏃犳晥鐨?game_id")
		return
	}

	companyIDs, err := h.service.GetCompanyIDsByGameID(
		c.Request.Context(),
		gameID,
	)

	if err != nil {
		respondGameCompanyError(c, err)
		return
	}

	if companyIDs == nil {
		companyIDs = []uint64{}
	}

	response.Success(c, companyIDs)
}

