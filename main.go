package main

import (
	"github.com/kenju/go-minruby/repl"
	"log"
	"os"
)

func main() {
	log.Println("Hello! This is go-minruby!")

	repl.Start(os.Stdin, os.Stdout)
}