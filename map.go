package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getArea(url string) ([]NamedURL, error) {

	res, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("Error fetching map: %w", err)
	}

	decoder := json.NewDecoder(res.Body)
	var areas AreaResponse
	if err := decoder.Decode(&areas); err != nil {
		return nil, fmt.Errorf("Error while decoding the body %w", err)
	}
	PrevUrl = areas.Previous
	CurUrl = areas.Next

	return areas.Results, nil

}
func getMap() ([]NamedURL, error) {
	if CurUrl == nil {
		return nil, fmt.Errorf("Your'e on the last page")
	}
	areas, err := getArea(*CurUrl)
	if err != nil {
		return nil, err
	}
	return areas, nil
}
func getBMap() ([]NamedURL, error) {
	if PrevUrl == nil {
		return nil, fmt.Errorf("Your'e on the first page")
	}
	areas, err := getArea(*PrevUrl)

	if err != nil {
		return nil, err
	}
	return areas, nil
}

func commandMap() error {
	locations, err := getMap()
	if err != nil {
		return err
	}
	for _, location := range locations {
		fmt.Println(location.Name)
	}
	return nil
}

func commandBMap() error {
	locations, err := getBMap()
	if err != nil {
		return err
	}
	println()
	for _, location := range locations {
		fmt.Println(location.Name)
	}
	return nil
}
