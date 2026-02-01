package models

type Results struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type AreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []Results
}
