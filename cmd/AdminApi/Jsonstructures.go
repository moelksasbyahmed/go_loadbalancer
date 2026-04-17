package adminapi

type backend struct {
	Name       string `json:"name"`
	Url        string `json:"url"`
	MaxRequest int    `json:"max_request"`
}
type statusResponse struct {
	Name           string `json:"name"`
	Alive          bool   `json:"alive"`
	Url            string `json:"url"`
	CurrentTraffic int    `json:"current_traffic"`
	OverallTraffic int    `json:"overall_traffic"`
}
