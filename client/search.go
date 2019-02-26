package client

import (
	"encoding/base64"
	"encoding/json"
	"strconv"
)

type Search []string

// fofa search  result
type SearchResponse struct {
	ErrorMsg
	// search mode
	Mode string `json:"mode"`
	// query keyword
	Query string `json:"query"`
	// query page num
	Page int `json:"page"`
	// query total num
	Size int `json:"size"`
	// query result
	Results []Search `json:"results"`
}

// Search receive a query, fields, pagenum, return search results
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
