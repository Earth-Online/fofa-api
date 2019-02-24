package client

import "net/url"

const (
	CliVersion     = "4.10.2"
	FofaServer     = "fofa.so"
	ApiMy          = "/api/v1/info/my"
	ApiExploitList = "/api/v1/user/exploits"
	ApiMessage     = "/messages"
	ApiSearchRule  = "/api/v1/exploit/rule_list"
	ApiRuleInfo    = "/api/v1/exploit/rule_info"
	ApiShop        = "/exploits/shop"
	ApiShopSum     = "/exploits/all_exploits_count"
	ApiSearch      = "/api/v1/search/all"
	ApiCode        = "/api/v1/exploit/code"
	ApiPublish     = "/api/v1/exploit/publish"
)

//  GetApiUrl receive a api path,  get Full url
func GetApiUrl(path string) (reqUrl url.URL) {
	reqUrl = url.URL{
		Scheme: "https",
		Host:   FofaServer,
		Path:   path,
	}
	return
}
