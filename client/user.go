package client

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Fofa User
type User struct {
	Email  string `json:"email"`
	Token  string `json:"token, omitempty"`
	Client *http.Client
	UserInfo
}

func NewUser(email string, token string) *User {
	u := &User{Email: email, Token: token, Client: &http.Client{}}
	return u
}

func (u *User) AddQuery(query url.Values) (value url.Values) {
	query.Add("email", u.Email)
	query.Add("key", u.Token)
	query.Add("fofacliversion", CliVersion)
	return query
}

func (u *User) Req(reqUrl url.URL) (data []byte, err error) {
	resp, err := u.Client.Get(reqUrl.String())
	if err != nil {
		return
	}
	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		err = errors.New(fmt.Sprintf("error http code %d", resp.StatusCode))
		return
	}
	data, err = ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()
	return
}

func (u *User) ReqHtml(reqUrl url.URL) (doc *goquery.Document, err error) {
	resp, err := u.Client.Get(reqUrl.String())
	if err != nil {
		return
	}
	doc, err = goquery.NewDocumentFromReader(resp.Body)
	return
}

func (u *User) Me() (err error) {
	reqUrl := GetApiUrl(ApiMy)
	queryString := reqUrl.Query()
	reqUrl.RawQuery = u.AddQuery(queryString).Encode()
	body, err := u.Req(reqUrl)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &u.UserInfo)
	return
}

// get me poc
func (u *User) GetPoc() (exploits ExploitResponse, err error) {
	reqUrl := GetApiUrl(ApiExploitList)
	queryString := reqUrl.Query()
	reqUrl.RawQuery = u.AddQuery(queryString).Encode()
	body, err := u.Req(reqUrl)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &exploits)
	return
}

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

func (u *User) GetShopPocNum() (num int, err error) {
	reqUrl := GetApiUrl(ApiShopSum)
	queryString := reqUrl.Query()
	reqUrl.RawQuery = u.AddQuery(queryString).Encode()
	body, err := u.Req(reqUrl)
	if err != nil {
		return
	}
	num, err = strconv.Atoi(string(body))
	return
}

func (u *User) GetShopPoc(page int) (exploits []Exploit, err error) {
	reqUrl := GetApiUrl(ApiShop)
	queryString := reqUrl.Query()
	queryString.Add("page", strconv.Itoa(page))
	reqUrl.RawQuery = u.AddQuery(queryString).Encode()
	body, err := u.ReqHtml(reqUrl)
	if err != nil {
		return
	}
	body.Find(".datashow").Each(func(i int, selection *goquery.Selection) {
		updateTime, _ := selection.Find("time").Attr("datetime")
		exp := Exploit{
			Name:     selection.Find(".sc_s1").Text(),
			UpdateAt: updateTime,
			// todo
			//	Author:selection.Find(".fa fa-user").Next().Text(),
			Product: selection.Find("a[style=\"text-decoration:underline\"]").Text(),
			Status:  PUBLISH,
		}
		exploits = append(exploits, exp)
	})
	return
}

func (u *User) Search(query string, fields string, page int) (searchs SearchResponse, err error) {
	reqUrl := GetApiUrl(ApiSearch)
	queryString := reqUrl.Query()
	queryString.Add("fields", fields)
	queryString.Add("page", strconv.Itoa(page))
	queryString.Add("qbase64", base64.StdEncoding.EncodeToString([]byte(query)))
	reqUrl.RawQuery = u.AddQuery(queryString).Encode()
	body, err := u.Req(reqUrl)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &searchs)
	return
}

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
	ret := NewPublishResponse()
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

// user info.. from  api/v1/info/my
type UserInfo struct {
	Avatar     string `json:"avatar"`
	Fcoin      int    `json:"fcoin"`
	FofaServer bool   `json:"fofa_server"`
	CliVersion string `json:"fofacli_ver"`
	IsVerified bool   `json:"is_verified"`
	IsVip      bool   `json:"isvip"`
	MessageNum int    `json:"message"`
	UserName   string `json:"user_name"`
	VipLever   int    `json:"vip_lever"`
}
