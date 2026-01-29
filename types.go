package main

type NamedURL struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
type AreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []NamedURL
}

type CLICommand struct {
	name        string
	description string
	callback    func() error
}
