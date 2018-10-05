package main

import (
	"log"
	"os"

	"github.com/kenju/go-minruby/repl"
)

func main() {
	log.Println("Hello! This is go-minruby!")

	repl.Start(os.Stdin, os.Stdout)
}
