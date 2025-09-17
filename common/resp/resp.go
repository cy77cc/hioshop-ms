package resp

import (
	"github.com/cy77cc/hioshop_ms/common/utils"
	"github.com/cy77cc/hioshop_ms/common/xcode"
)

// 定义统一的响应结构体

type Resp struct {
	Code      xcode.Xcode `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

func Error(code xcode.Xcode, msg string) *Resp {
	return &Resp{
		Code:      code,
		Msg:       msg,
		Data:      nil,
		Timestamp: utils.GetTimestamp(),
	}
}

func Success(data interface{}) *Resp {
	return &Resp{
		Code:      xcode.Success,
		Msg:       "success",
		Data:      data,
		Timestamp: utils.GetTimestamp(),
	}
}
