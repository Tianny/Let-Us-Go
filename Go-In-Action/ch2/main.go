package main

import (
	_ "LetUsGo/Go-In-Action/ch2/matchers"
	"LetUsGo/Go-In-Action/ch2/search"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
