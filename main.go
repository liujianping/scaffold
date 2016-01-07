package main

import (
	"log"
	"os"
	"runtime"

	"github.com/liujianping/scaffold/cmd"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println(os.Args[0], " ", r)
		}
	}()

	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.App().Run(os.Args)
}
