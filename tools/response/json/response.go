package json

import (
	"encoding/json"
	"net/http"
)

const (
	Success = 1
	Fail    = -1
	Error   = -2
)

var response Response

func init() {
	response = Response{
		Code: Success,
		Data: json.Encoder{},
	}
}

// 响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// 基础json响应
func Json(w http.ResponseWriter, status int, data interface{}) {

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)

}

// list 响应
func ResponseJson(w http.ResponseWriter, data interface{})  {

	response.Data=data

	response.Message="请求成功!"

	Json(w, http.StatusOK, response)
}

// http status 201
func WithCreated(w http.ResponseWriter) {

	response.Message = "添加成功!"

	Json(w, http.StatusCreated, response)

}
