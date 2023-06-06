package api

import (
	resp "awesomeProject/model/response"
	"awesomeProject/pkg/e"
)

// ErrorResponse 返回错误信息
func ErrorResponse(err error) *resp.TrackedErrorResponse {
	return resp.RespError(err, "参数错误", e.InvalidParams)

}
