package main

import (
	"fmt"
)

type Gopher struct {
	Gopher string `json:"gopher"`
}

func main() {
	const gopher = "GOPHER"
	gogopher := GOPHER()
	gogopher.Gopher = gopher
	fmt.Println(gogopher)
}

func GOPHER() (gopher *Gopher) {
	gopher = &Gopher{Gopher: "gopher"}
	return
}
