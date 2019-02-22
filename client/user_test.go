package client

import "testing"

// Please don't submit it
var email = "test@qq.com"
var token = "123456"

func TestNewUser(t *testing.T) {
	fakeEmail := "fake"
	fakeToken := "fake"
	user := NewUser(fakeEmail, fakeToken)
	if user.Email != fakeEmail && user.Token != fakeToken {
		t.Error("error create user")
	}
}

func TestUser_Me(t *testing.T) {
	user := NewUser(email, token)
	err := user.Me()
	if err != nil {
		t.Error(err)
	}
}

func TestGetApiUrl(t *testing.T) {
	//user := NewUser(email, token)
	reqUrl := GetApiUrl(ApiMy)
	if reqUrl.Scheme != "https" && reqUrl.Host != FofaServer && reqUrl.Path != ApiMy {
		t.Error("error create api url")
	}
}

func TestUser_GetPoc(t *testing.T) {
	user := NewUser(email, token)
	resp, err := user.GetPoc()
	if err != nil {
		t.Error(err)
	}
	if resp.Error {
		t.Error("error get poc list")
	}
	t.Log(resp)
}

func TestUser_GetMessages(t *testing.T) {
	user := NewUser(email, token)
	resp, err := user.GetMessages()
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}

func TestUser_GetShopPocNum(t *testing.T) {
	user := NewUser(email, token)
	resp, err := user.GetShopPocNum()
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}

func TestUser_GetShopPoc(t *testing.T) {
	user := NewUser(email, token)
	resp, err := user.GetShopPoc(1)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}

func TestUser_Search(t *testing.T) {
	user := NewUser(email, token)
	resp, err := user.Search("domain=\"nosec.org\"", "domain,host,ip,port,title,country,city", 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}

func TestUser_GetPocCode(t *testing.T) {
	user := NewUser(email, token)
	resp, err := user.GetPocCode("phpyun_v4_install_getshell.rb")
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}

func TestUser_SearchRule(t *testing.T) {
	user := NewUser(email, token)
	resp, err := user.SearchRule("i")
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}

func TestUser_GetRuleInfo(t *testing.T) {
	user := NewUser(email, token)
	resp, err := user.GetRuleInfo("IdeaCMS")
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
