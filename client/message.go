package client

import "github.com/PuerkitoBio/goquery"

type Message struct {
	Time string `json:"time"`
	Msg  string `json:"msg"`
}

func NewMessage(time string, msg string) *Message {
	return &Message{
		Time: time,
		Msg:  msg,
	}
}

func (u *User) GetMessages() (messages []Message, err error) {
	reqUrl := GetApiUrl(ApiMessage)
	queryString := reqUrl.Query()
	reqUrl.RawQuery = u.AddQuery(queryString).Encode()
	body, err := u.ReqHtml(reqUrl)
	if err != nil {
		return
	}

	body.Find(".xx_c").Each(func(i int, selection *goquery.Selection) {
		time := selection.Find(".xx_time").Text()
		message := selection.Find("li").Text()
		messages = append(messages, *NewMessage(time, message))
	})
	return
}
