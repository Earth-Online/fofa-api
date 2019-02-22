package client

import (
	"encoding/json"
	"fmt"
	"time"
)

type SearchRuleResponse struct {
	Error bool     `json:"error"`
	Data  []string `json:"data"`
}

type RuleInfoResponse struct {
	Error bool   `json:"error"`
	Rule  string `json:"rule"`
	Url   string `json:"url"`
}

func (u *User) SearchRule(key string) (resp SearchRuleResponse, err error) {
	reqUrl := GetApiUrl(ApiSearchRule)
	queryString := reqUrl.Query()
	queryString.Add("product", key)
	queryString.Add("_", fmt.Sprintf("%d", time.Now().Unix()))
	reqUrl.RawQuery = u.AddQuery(queryString).Encode()
	body, err := u.Req(reqUrl)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return
	}
	return
}

func (u *User) GetRuleInfo(name string) (rule RuleInfoResponse, err error) {
	reqUrl := GetApiUrl(ApiRuleInfo)
	queryString := reqUrl.Query()
	queryString.Add("product", name)
	reqUrl.RawQuery = u.AddQuery(queryString).Encode()
	body, err := u.Req(reqUrl)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &rule)
	if err != nil {
		return
	}
	return
}
