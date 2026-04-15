package adminapi

type backend struct {
	Name       string `json:"name"`
	Url        string `json:"url"`
	MaxRequest int    `json:"max_request"`
}
