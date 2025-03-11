package api

type LocationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []LocationListItem
}

type LocationListItem struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func MapRequest(path string) (LocationAreaResponse, error) {
	return MakeRequest[LocationAreaResponse](path)
}
