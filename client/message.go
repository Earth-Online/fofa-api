package client


type Message struct {
	Time string `json:"time"`
	Msg string `json:"msg"`
}

func NewMessage(time string, msg string)(*Message){
	return &Message{
		Time:time,
		Msg:msg,
	}
}
