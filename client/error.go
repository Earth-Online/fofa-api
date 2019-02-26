package client

import "errors"

//  error info
type ErrorMsg struct {
	// is error?
	Error bool `json:"error"`
	// error msg
	Errmsg string `json:"errmsg"`
}

func NewErrorMsg(errmsg string) *ErrorMsg {
	return &ErrorMsg{Errmsg: errmsg, Error: true}
}

// GetErrorMsg get error msg
func (e *ErrorMsg) GetErrorMsg() error {
	if e.Error {
		return errors.New(e.Errmsg)
	}
	return nil
}
