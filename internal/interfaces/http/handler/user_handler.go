package handler

import (
	app "galhub/internal/app/user"
	"galhub/internal/interfaces/http/dto"
	"galhub/internal/interfaces/http/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *app.Service
}

func NewUserHandler(service *app.Service) *UserHandler {
	return &UserHandler{service: service}
}
func (h *UserHandler) Register(
	c *gin.Context,
) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, err.Error())
		return
	}
	err := h.service.Register(
		c.Request.Context(),
		&app.RegisterCommand{
			Username: req.Username,
			Email:    req.Email,
			Password: req.Password,
		},
	)
	if err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *UserHandler) Login(
	c *gin.Context,
) {

	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, err.Error())
		return
	}

	user, err := h.service.Login(
		c.Request.Context(),
		app.LoginCommand{
			Email:    req.Email,
			Password: req.Password,
		},
	)

	if err != nil {
		response.Fail(c, err.Error())
		return
	}

	response.Success(c, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	})
}
