package handler

import (
	app "galhub/internal/app/user"
	dto "galhub/internal/interfaces/http/dto/user"
	"galhub/internal/pkg/response"

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
		"token":    user.Token,
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	})
}
func (h *UserHandler) Profile(
	c *gin.Context,
) {

	userID := c.GetUint64(
		"user_id",
	)

	user, err := h.service.GetProfile(
		c.Request.Context(),
		userID,
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
		user,
	)
}
func (h *UserHandler) UpdateProfile(
	c *gin.Context,
) {

	userID := c.GetUint64("user_id")

	var req dto.UpdateProfileRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, err.Error())
		return
	}

	err := h.service.UpdateProfile(
		c.Request.Context(),
		userID,
		app.UpdateProfileCommand{
			Nickname: req.Nickname,
			Avatar:   req.Avatar,
		},
	)

	if err != nil {
		response.Fail(c, err.Error())
		return
	}

	response.Success(c, nil)
}
func (h *UserHandler) ChangePassword(
	c *gin.Context,
) {

	userID := c.GetUint64(
		"user_id",
	)

	var req dto.ChangePasswordRequest

	if err := c.ShouldBindJSON(
		&req,
	); err != nil {

		response.Fail(
			c,
			err.Error(),
		)

		return
	}

	err := h.service.ChangePassword(
		c.Request.Context(),
		userID,
		app.ChangePasswordCommand{
			OldPassword: req.OldPassword,
			NewPassword: req.NewPassword,
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
