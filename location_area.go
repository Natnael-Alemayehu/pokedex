package main

type locationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []name `json:"results"`
}

type name struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
