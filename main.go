package main

import (
    "net/http"
    "fmt"
    "log"
    "strings"
    "flag"
)

var (
    url = flag.String("url", "", "URL address")
)

func main() {

    flag.Parse()
    if *url == "" {
        log.Fatal("not set url")
    }

    res, err := http.Get(*url)
    if err != nil {
        log.Fatal(err)
    }

    if err != nil {
        log.Fatal(err)
    }
    server, ok := res.Header["Server"]
    if ok {
        fmt.Printf("Server: %s\n", strings.Join(server, ", "))
    }

    power, ok := res.Header["X-Powered-By"]
    if ok {
        fmt.Printf("Powered: %s\n", strings.Join(power, ", "))
    }

    cookie, ok := res.Header["Set-Cookie"]
    if ok {
        fmt.Printf("Cookie: %s\n", strings.Join(cookie, ", "))
    }
}
