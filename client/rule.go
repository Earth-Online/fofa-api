package client

import (
	"encoding/json"
	"fmt"
	"time"
)

// search rule result
type SearchRuleResponse struct {
	ErrorMsg
	// rule name
	Data []string `json:"data"`
}

// rule info result
type RuleInfoResponse struct {
	ErrorMsg
	// 查询规则细节
	Rule string `json:"rule"`
	// 规则名
	Url string `json:"url"`
}

// SearchRule  receive search key return search response
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

// SearchRule  receive rulename return rule info
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
