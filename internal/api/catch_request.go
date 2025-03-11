package api

type CatchResponse struct {
	BaseExperience int        `json:"base_experience"`
	Height         int        `json:"height"`
	Weight         int        `json:"weight"`
	Name           string     `json:"name"`
	Stats          []StatItem `json:"stats"`
	Types          []TypeItem `json:"types"`
}

type StatItem struct {
	BaseStat int `json:"base_stat"`
	Effort   int `json:"effort"`
	Stat     struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"stat"`
}

type TypeItem struct {
	Slot int `json:"slot"`
	Type struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"type"`
}

func CatchRequest(path string) (CatchResponse, error) {
	return MakeRequest[CatchResponse](path)
}
