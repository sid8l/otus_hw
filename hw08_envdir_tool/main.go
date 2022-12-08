package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		fmt.Println("Example: ./envdir /path/to/dir command arg1 arg2")
		os.Exit(1)
	}
	env, err := ReadDir(flag.Arg(0))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(RunCmd(flag.Args()[1:], env))
}
