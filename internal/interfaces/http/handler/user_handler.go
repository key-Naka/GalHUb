package handler

import (
	app "galhub/internal/app/user"
	usercmd "galhub/internal/app/user/command"
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
		&usercmd.RegisterCommand{
			Username: req.Username,
			Email:    req.Email,
			Password: req.Password,
		},
	)
	if err != nil {
		respondUserError(c, err)
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
		&usercmd.LoginCommand{
			Email:    req.Email,
			Password: req.Password,
		},
	)

	if err != nil {
		respondUserAuthError(c, err)
		return
	}

	response.Success(c, user)
}
func (h *UserHandler) Profile(
	c *gin.Context,
) {
	userID, ok := currentUserID(c)
	if !ok {
		response.Unauthorized(c, "未授权")
		return
	}

	user, err := h.service.GetProfile(
		c.Request.Context(),
		userID,
	)

	if err != nil {
		respondUserError(c, err)
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
	userID, ok := currentUserID(c)
	if !ok {
		response.Unauthorized(c, "未授权")
		return
	}

	var req dto.UpdateProfileRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, err.Error())
		return
	}

	err := h.service.UpdateProfile(
		c.Request.Context(),
		userID,
		&usercmd.UpdateProfileCommand{
			Nickname: req.Nickname,
			Avatar:   req.Avatar,
		},
	)

	if err != nil {
		respondUserError(c, err)
		return
	}

	response.Success(c, nil)
}
func (h *UserHandler) ChangePassword(
	c *gin.Context,
) {
	userID, ok := currentUserID(c)
	if !ok {
		response.Unauthorized(c, "未授权")
		return
	}

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
		&usercmd.ChangePasswordCommand{
			OldPassword: req.OldPassword,
			NewPassword: req.NewPassword,
		},
	)

	if err != nil {
		respondUserError(c, err)
		return
	}

	response.Success(
		c,
		nil,
	)
}

func currentUserID(
	c *gin.Context,
) (uint64, bool) {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0, false
	}

	id, ok := userID.(uint64)
	if !ok || id == 0 {
		return 0, false
	}

	return id, true
}
