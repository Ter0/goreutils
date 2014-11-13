package main

import (
    "fmt"
    "flag"
    "os"
)

func version() {
    fmt.Println("0.0.0")
    os.Exit(0)
}

func main() {

    versionPtr := flag.Bool("version", false, "print the version")
    flag.Parse()

    // only print the version if it's the only option
    if flag.NFlag() == 1 && *versionPtr == true {
        version()
    }

    dir, err := os.Getwd()

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Println(dir)
}
