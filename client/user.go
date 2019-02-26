package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// Fofa User
type User struct {
	// account email
	Email string `json:"email"`
	// account api token
	Token  string `json:"token, omitempty"`
	Client *http.Client
	UserInfo
}

// user info
type UserInfo struct {
	// user avatar
	Avatar string `json:"avatar"`
	// f coin num
	Fcoin int `json:"fcoin"`
	// fofa domain
	FofaServer bool `json:"fofa_server"`
	// cli version
	CliVersion string `json:"fofacli_ver"`
	// is verify user?
	IsVerified bool `json:"is_verified"`
	// is vip user?
	IsVip bool `json:"isvip"`
	// unread message num
	MessageNum int `json:"message"`
	// username
	UserName string `json:"user_name"`
	// viplevel
	VipLever int `json:"vip_lever"`
}

// NewUser Get an a instance from email and token
func NewUser(email string, token string) *User {
	u := &User{Email: email, Token: token, Client: &http.Client{}}
	return u
}

// AddQuery will add authentication parameter
func (u *User) AddQuery(query url.Values) (value url.Values) {
	query.Add("email", u.Email)
	query.Add("key", u.Token)
	query.Add("fofacliversion", CliVersion)
	return query
}

// Req make a request from the url
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

// Req make a request from the url, then Convert to a goquery document
func (u *User) ReqHtml(reqUrl url.URL) (doc *goquery.Document, err error) {
	resp, err := u.Client.Get(reqUrl.String())
	if err != nil {
		return
	}
	doc, err = goquery.NewDocumentFromReader(resp.Body)
	return
}

// Me will get user info
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

// get shop poc total
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

// get number of pages exploit info
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
