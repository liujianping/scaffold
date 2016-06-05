package main

import (
	"log"
	"os"

	"github.com/liujianping/scaffold/commands"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println(os.Args[0], " ", r)
		}
	}()

	commands.App().Run(os.Args)
}
