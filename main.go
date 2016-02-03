package main

import (
	"github.com/liujianping/scaffold/cmd"
	"log"
	"os"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println(os.Args[0], " ", r)
		}
	}()

	cmd.App().Run(os.Args)
}
