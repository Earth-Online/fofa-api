package client

import "errors"

//  ErrorMsg 是错误信息类型 仅用于嵌入.
type ErrorMsg struct {
	// 请求是否错误
	Error bool `json:"error"`
	// 错误信息
	Errmsg string `json:"errmsg"`
}

func NewErrorMsg(errmsg string) *ErrorMsg {
	return &ErrorMsg{Errmsg: errmsg, Error: true}
}

// GetErrorMsg 确认是否发生错误.
func (e *ErrorMsg) GetErrorMsg() error {
	if e.Error {
		return errors.New(e.Errmsg)
	}
	return nil
}
