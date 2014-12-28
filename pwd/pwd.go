package main

import (
	"flag"
	"fmt"
	"github.com/ter0/goreutils/common"
	"os"
)

func main() {

	versionPtr := flag.Bool("version", false, "print the version")
	flag.Parse()

	// only print the version if it's the only option
	if flag.NFlag() == 1 && *versionPtr == true {
		common.PrintVersion()
		os.Exit(0)
	}

	dir, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(dir)
}
