package client

import "errors"

type CodeResponse struct {
	Error bool `json:"error"`
	Errmsg string `json:"errmsg"`
	Code string `json:"code"`
}

func (c *CodeResponse) GetError() error {
	if c.Error {
		return errors.New(c.Errmsg)
	}
	return nil
}
