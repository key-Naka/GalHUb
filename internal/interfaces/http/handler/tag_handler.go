package handler

import (
	tagApp "galhub/internal/app/tag"
	"galhub/internal/app/tag/command"
	dto "galhub/internal/interfaces/http/dto/tag"
	"galhub/internal/pkg/response"
	"galhub/internal/pkg/util"

	"github.com/gin-gonic/gin"
)

type TagHandler struct {
	service *tagApp.Service
}

func NewTagHandler(
	service *tagApp.Service,
) *TagHandler {
	return &TagHandler{
		service: service,
	}
}

func (h *TagHandler) Create(
	c *gin.Context,
) {
	var req dto.CreateTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err := h.service.Create(
		c.Request.Context(),
		command.CreateTag{
			Name: req.Name,
		},
	)

	if err != nil {
		respondTagError(c, err)
		return
	}

	response.Success(c, nil)
}

func (h *TagHandler) GetByID(
	c *gin.Context,
) {
	id, err := util.ParseID(c)
	if err != nil {
		response.BadRequest(c, "鏃犳晥鐨?id")
		return
	}

	tag, err := h.service.GetByID(
		c.Request.Context(),
		id,
	)

	if err != nil {
		respondTagError(c, err)
		return
	}

	response.Success(c, tag)
}

func (h *TagHandler) List(
	c *gin.Context,
) {
	tags, err := h.service.List(
		c.Request.Context(),
	)

	if err != nil {
		respondTagError(c, err)
		return
	}

	response.Success(c, tags)
}

func (h *TagHandler) Update(
	c *gin.Context,
) {
	id, err := util.ParseID(c)
	if err != nil {
		response.BadRequest(c, "鏃犳晥鐨?id")
		return
	}

	var req dto.UpdateTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	err = h.service.Update(
		c.Request.Context(),
		command.UpdateTag{
			ID:   id,
			Name: req.Name,
		},
	)

	if err != nil {
		respondTagError(c, err)
		return
	}

	response.Success(
		c,
		gin.H{
			"id": id,
		},
	)
}

func (h *TagHandler) Delete(
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
		respondTagError(c, err)
		return
	}

	response.Success(
		c,
		gin.H{
			"id": id,
		},
	)
}

