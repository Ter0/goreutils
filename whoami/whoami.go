package main

import (
    "flag"
    "fmt"
    "github.com/ter0/goreutils/common"
    "os"
    "os/user"
)

func main() {
    versionPtr := flag.Bool("version", false, "print the version")
    flag.Parse()

    if *versionPtr == true {
        common.PrintVersion()
        os.Exit(0)
    }

    user, err := user.Current()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Println(user.Username)
}
