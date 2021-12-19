package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
    "strings"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "<style>body { font-size: 3em; font-family: system-ui; display: flex; align-items: center; justify-content: center; height: 100 vh ; }</style><h1>Hello, %s</h1>", strings.TrimLeft(html.EscapeString(r.URL.Path), "/"))
    })

    fmt.Print("listening on 4000...")
    log.Fatal(http.ListenAndServe(":4000", nil))

}
