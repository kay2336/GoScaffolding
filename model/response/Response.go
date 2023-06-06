package resp

import (
	"awesomeProject/pkg/e"
	"net/http"
)

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

// RespList
// 带有总数的列表构建器 ???
func RespList(items interface{}, total int64) Response {
	return Response{
		Status: http.StatusOK,
		//Data: DataList{
		//	Item:  items,
		//	Total: total,
		//},
		Msg: "ok",
	}
}

// RespSuccess
// 成功返回
func RespSuccess() *Response {
	status := e.SUCCESS
	return &Response{
		Status: status,
		Data:   "操作成功",
		Msg:    e.GetMsg(status),
	}

}

// RespError
// 错误返回
func RespError(err error, data string, status int) *TrackedErrorResponse {
	status = e.ERROR
	return &TrackedErrorResponse{
		Response: Response{
			Status: status,
			Msg:    e.GetMsg(status),
			Data:   data,
			Error:  err.Error(),
		},
		// TrackId:  // TODO 加上track id
	}

}

// RespSuccessWithData 带data成功返回
func RespSuccessWithData(data interface{}, code ...int) *Response {
	status := e.SUCCESS
	if code != nil {
		status = code[0]
	}

	r := &Response{
		Status: status,
		Data:   data,
		Msg:    e.GetMsg(status),
	}

	return r
}
