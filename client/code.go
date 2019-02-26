package client

import (
	"encoding/base64"
	"encoding/json"
	"errors"
)

// CodeResponse is install poc code request response
type CodeResponse struct {
	Code string `json:"code"`
	ErrorMsg
}

// GetError get error and msg
func (c *CodeResponse) GetError() error {
	if c.Error {
		return errors.New(c.Errmsg)
	}
	return nil
}

// download you poc code
func (u *User) GetPocCode(filename string) (code string, err error) {
	reqUrl := GetApiUrl(ApiCode)
	queryString := reqUrl.Query()
	queryString.Add("filename", filename)
	reqUrl.RawQuery = u.AddQuery(queryString).Encode()
	body, err := u.Req(reqUrl)
	if err != nil {
		return
	}
	expCode := CodeResponse{}
	err = json.Unmarshal(body, &expCode)
	if err != nil {
		return
	}
	if expCode.Error {
		err = expCode.GetError()
		return
	}
	exp, err := base64.StdEncoding.DecodeString(expCode.Code)
	code = string(exp)
	return
}
