package main

import (
    "fmt"
    "flag"
    "os"
    "io"
    "bufio"
    "github.com/ter0/goreutils/common"
    "crypto/md5"
)

func main() {
    versionPtr := flag.Bool("version", false, "print the version")
    flag.Parse()

    if *versionPtr == true {
        common.PrintVersion()
        os.Exit(0)
    }

    // a list of files to try and hash
    args := flag.Args()

    for _, fileName := range args {
        hash := md5.New()
        file, err := os.Open(fileName)

        if err != nil {
            fmt.Printf("md5sum: %s\n", err)
            continue
        }

        reader := bufio.NewReader(file)
        buffer := make([]byte, 1024)

        ok := true
        for {
            n, err := reader.Read(buffer)
            if err != nil && err != io.EOF {
                fmt.Printf("md5sum: %s\n", err)
                ok = false
                break
            }
            if n == 0 {
                break
            }
            io.WriteString(hash, string(buffer[:n]))
        }

        if ok == true {
            fmt.Printf("%x  %s\n", hash.Sum(nil), fileName)
        }

        ok = true
    }
}
