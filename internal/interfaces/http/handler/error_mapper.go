package handler

import (
	"errors"
	companyApp "galhub/internal/app/company"
	favoriteApp "galhub/internal/app/favorite"
	gameApp "galhub/internal/app/game"
	gameCompanyApp "galhub/internal/app/game_company"
	gameTagApp "galhub/internal/app/game_tag"
	tagApp "galhub/internal/app/tag"
	userApp "galhub/internal/app/user"
	"galhub/internal/pkg/response"
	"log"

	"github.com/gin-gonic/gin"
)

func respondInternalError(
	c *gin.Context,
	err error,
) {
	log.Printf("http handler internal error: %v", err)
	response.Internal(c, "服务器内部错误")
}

func respondCompanyError(
	c *gin.Context,
	err error,
) {
	switch {
	case errors.Is(err, companyApp.ErrCompanyNotFound):
		response.NotFound(c, err.Error())
	default:
		respondInternalError(c, err)
	}
}

func respondGameCompanyError(
	c *gin.Context,
	err error,
) {
	switch {
	case errors.Is(err, gameCompanyApp.ErrGameCompanyNotFound):
		response.NotFound(c, err.Error())
	default:
		respondInternalError(c, err)
	}
}

func respondGameTagError(
	c *gin.Context,
	err error,
) {
	switch {
	case errors.Is(err, gameTagApp.ErrGameTagNotFound):
		response.NotFound(c, err.Error())
	default:
		respondInternalError(c, err)
	}
}

func respondTagError(
	c *gin.Context,
	err error,
) {
	switch {
	case errors.Is(err, tagApp.ErrTagNotFound):
		response.NotFound(c, err.Error())
	default:
		respondInternalError(c, err)
	}
}

func respondGameError(
	c *gin.Context,
	err error,
) {
	switch {
	case errors.Is(err, gameApp.ErrGameNotFound):
		response.NotFound(c, err.Error())
	case errors.Is(err, gameApp.ErrInvalidPagination):
		response.BadRequest(c, err.Error())
	default:
		respondInternalError(c, err)
	}
}

func respondFavoriteError(
	c *gin.Context,
	err error,
) {
	switch {
	case errors.Is(err, favoriteApp.ErrInvalidPagination):
		response.BadRequest(c, err.Error())
	case errors.Is(err, favoriteApp.ErrFavoriteExists):
		response.Conflict(c, err.Error())
	case errors.Is(err, favoriteApp.ErrFavoriteNotFound), errors.Is(err, favoriteApp.ErrGameNotFound):
		response.NotFound(c, err.Error())
	default:
		respondInternalError(c, err)
	}
}

func respondUserAuthError(
	c *gin.Context,
	err error,
) {
	switch {
	case errors.Is(err, userApp.ErrUserNotFound), errors.Is(err, userApp.ErrInvalidPassword):
		response.Unauthorized(c, "邮箱或密码错误")
	default:
		respondInternalError(c, err)
	}
}

func respondUserError(
	c *gin.Context,
	err error,
) {
	switch {
	case errors.Is(err, userApp.ErrInvalidCommand):
		response.BadRequest(c, err.Error())
	case errors.Is(err, userApp.ErrUserExists), errors.Is(err, userApp.ErrEmailExists):
		response.Conflict(c, err.Error())
	case errors.Is(err, userApp.ErrUserNotFound):
		response.NotFound(c, err.Error())
	case errors.Is(err, userApp.ErrInvalidPassword):
		response.BadRequest(c, "旧密码错误")
	default:
		respondInternalError(c, err)
	}
}
