package client


type Search []string

type SearchResponse struct {
	Mode string `json:"mode"`
	Error string `json:"error"`
	Query string `json:"query"`
	Page int `json:"page"`
	Size int `json:"size"`
	Results []Search `json:"results"`
}

