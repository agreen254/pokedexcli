package api

type LocationAreaResponse struct {
	Count    int
	Next     *string
	Previous *string
	Results  []LocationListItem
}

type LocationListItem struct {
	Name string
	Url  string
}

func MakeRequestMap(path string) (LocationAreaResponse, error) {
	var res LocationAreaResponse
	err := MakeRequest(path, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}
