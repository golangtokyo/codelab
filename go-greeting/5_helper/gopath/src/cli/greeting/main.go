package main

import (
	"greeting"
	"os"
)

func main() {
	var g greeting.Greeting
	g.Do(os.Stdout)
}
