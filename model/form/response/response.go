package resp

import "awesomeProject/pkg/e"

// Response
// 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// TrackedErrorResponse
// 有追踪信息的错误反应
type TrackedErrorResponse struct {
	Response
	TrackId string `json:"track_id"`
}

// RespSuccess
// 成功返回
func RespSuccess() *Response {
	status := e.SUCCESS
	r := &Response{
		Status: status,
		Data:   "操作成功",
		Msg:    e.GetMsg(status),
	}
	return r
}

// RespError
// 错误返回
func RespError(err error, data string) *TrackedErrorResponse {
	status := e.ERROR

	r := &TrackedErrorResponse{
		Response: Response{
			Status: status,
			Msg:    e.GetMsg(status),
			Data:   data,
			Error:  err.Error(),
		},
		// TrackId:  // TODO 加上track id
	}

	return r
}
