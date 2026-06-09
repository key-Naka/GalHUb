package response

import "github.com/gin-gonic/gin"

func json(
	c *gin.Context,
	httpStatus int,
	code int,
	msg string,
	data any,
) {
	body := gin.H{
		"code": code,
		"msg":  msg,
	}

	if data != nil {
		body["data"] = data
	}

	c.JSON(httpStatus, body)
}

func Success(
	c *gin.Context,
	data any,
) {
	json(c, 200, 0, "success", data)
}

func Fail(
	c *gin.Context,
	msg string,
) {
	BadRequest(c, msg)
}

func BadRequest(
	c *gin.Context,
	msg string,
) {
	json(c, 400, 1, msg, nil)
}

func Unauthorized(
	c *gin.Context,
	msg string,
) {
	json(c, 401, 1, msg, nil)
}

func NotFound(
	c *gin.Context,
	msg string,
) {
	json(c, 404, 1, msg, nil)
}

func Conflict(
	c *gin.Context,
	msg string,
) {
	json(c, 409, 1, msg, nil)
}

func Internal(
	c *gin.Context,
	msg string,
) {
	json(c, 500, 1, msg, nil)
}
