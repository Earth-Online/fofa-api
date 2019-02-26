package client

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

type PublishResponse struct {
	ErrorMsg
}

//  PublishPoc will publish you poc
func (u *User) PublishPoc(exp Exploit, code []byte) (err error) {
	b64 := base64.StdEncoding.EncodeToString(code)
	codeMd5 := md5.Sum([]byte("StFofart" + b64 + "EPrond"))
	reqUrl := GetApiUrl(ApiPublish)
	queryString := reqUrl.Query()
	queryString.Add("title", exp.Name)
	queryString.Add("filename", exp.Filename)
	queryString.Add("product", exp.Product)
	queryString.Add("homepage", exp.Homapage)
	queryString.Add("level", strconv.Itoa(exp.Level))
	queryString.Add("0day", strconv.FormatBool(exp.IsZeroDay))
	queryString.Add("codemd5", string(codeMd5[:]))

	reqUrl.RawQuery = u.AddQuery(queryString).Encode()
	resp, err := u.Client.Post(reqUrl.String(), "application/x-www-form-urlencoded", strings.NewReader(b64))
	if err != nil {
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	ret := &PublishResponse{}
	err = json.Unmarshal(data, &ret)
	_ = resp.Body.Close()
	if err != nil {
		return err
	}
	if ret.Error {
		return errors.New(ret.Errmsg)
	}
	return
}
