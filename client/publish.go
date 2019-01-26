package client

type PublishResponse struct {
	Error  bool   `json:"error"`
	Errmsg string `json:"errmsg"`
}

func NewPublishResponse() *PublishResponse {
	return &PublishResponse{}
}
