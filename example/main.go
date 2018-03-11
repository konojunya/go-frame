package main

import (
	frame "github.com/konojunya/go-frame"
)

type User struct {
	Name string `frame:"Customer Name"`
	Age  int    `frame:"Age"`
	URL  string `frame:"HomePage URL"`
}

func main() {
	user := User{
		Name: "konojunya",
		Age:  21,
		URL:  "https://konojunya.com",
	}

	frame.Print(user)
}
