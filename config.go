package main

type config struct {
	Next     string
	Previous string
}

func NewConfig(next, previous string) config {
	return config{
		Next:     next,
		Previous: previous,
	}
}
