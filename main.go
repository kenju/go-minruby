package main

import (
	"log"
	"os"
	"os/user"

	"github.com/kenju/go-minruby/repl"
)

func main() {
	usr, _ := user.Current()

	log.Printf("Hello %s! Welcome to go-minruby!", usr.Username)

	repl.Start(os.Stdin, os.Stdout)
}
