package main

import (
	"log"
	"os"

	"github.com/liujianping/scaffold/cmd"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println(os.Args[0], " ", r)
		}
	}()

	cmd.App().Run(os.Args)
}
