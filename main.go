package main

var Url string = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
var CurUrl *string = &Url
var PrevUrl *string = nil

func main() {
	startRepl()
}
