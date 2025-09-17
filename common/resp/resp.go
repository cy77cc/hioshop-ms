package resp

import "github.com/cy77cc/hioshop_ms/common/xcode"

// 定义统一的响应结构体

type Resp struct {
	Code xcode.Xcode `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Error(code xcode.Xcode, msg string) *Resp {
	return &Resp{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

func Success(data interface{}) *Resp {
	return &Resp{
		Code: xcode.Success,
		Msg:  "success",
		Data: data,
	}
}
