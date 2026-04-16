package main

type Backend struct {
	Name        string `json:"name"`
	Url         string `json:"url"`
	Max_request int    `json:"max_request"`
}
