package handler

import (
	companyApp "galhub/internal/app/company"
	"galhub/internal/app/company/command"
	dto "galhub/internal/interfaces/http/dto/company"
	"galhub/internal/pkg/response"
	"galhub/internal/pkg/util"

	"github.com/gin-gonic/gin"
)

type CompanyHandler struct {
	service *companyApp.Service
}

func NewCompanyHandler(
	service *companyApp.Service,
) *CompanyHandler {
	return &CompanyHandler{
		service: service,
	}
}

func (h *CompanyHandler) Create(
	c *gin.Context,
) {
	var req dto.CreateCompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err := h.service.Create(
		c.Request.Context(),
		command.CreateCompany{
			Name: req.Name,
		},
	)

	if err != nil {
		respondCompanyError(c, err)
		return
	}

	response.Success(c, nil)
}

func (h *CompanyHandler) GetByID(
	c *gin.Context,
) {
	id, err := util.ParseID(c)
	if err != nil {
		response.BadRequest(c, "鏃犳晥鐨?id")
		return
	}

	company, err := h.service.GetByID(
		c.Request.Context(),
		id,
	)

	if err != nil {
		respondCompanyError(c, err)
		return
	}

	response.Success(c, company)
}

func (h *CompanyHandler) List(
	c *gin.Context,
) {
	companies, err := h.service.List(
		c.Request.Context(),
	)

	if err != nil {
		respondCompanyError(c, err)
		return
	}

	response.Success(c, companies)
}

func (h *CompanyHandler) Update(
	c *gin.Context,
) {
	id, err := util.ParseID(c)
	if err != nil {
		response.BadRequest(c, "鏃犳晥鐨?id")
		return
	}

	var req dto.UpdateCompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err = h.service.Update(
		c.Request.Context(),
		command.UpdateCompany{
			ID:   id,
			Name: req.Name,
		},
	)

	if err != nil {
		respondCompanyError(c, err)
		return
	}

	response.Success(
		c,
		gin.H{
			"id": id,
		},
	)
}

func (h *CompanyHandler) Delete(
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
		respondCompanyError(c, err)
		return
	}

	response.Success(
		c,
		gin.H{
			"id": id,
		},
	)
}
