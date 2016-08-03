package main

import (
    "fmt"
    "net/http"
    "log"
)

var Alice = `
[{"url": "cat1.jpg", "description": "burrito cat"},
{"url": "cat2.jpg", "description": "orange grass cat"},
{"url": "cat3.jpg", "description": "spotted lynx cat"}]   
`
var Bob = `
[{"url": "dog1.jpg", "description": "tilted head dog"},
{"url": "dog3.jpg", "description": "puppy red ball dog"},
{"url": "dog4.jpg", "description": "puppy stairs colloar dog"}]   
`

var Carol = `
[{"url": "moose1.jpg", "description": "moose"},
{"url": "moose2.jpg", "description": "braying moose molting"},
{"url": "moose4.jpg", "description": "dark moose"}]   
`

var David = `
[{"url": "pokemon1.jpg", "description": "pikachu sky"},
{"url": "pokemon2.jpg", "description": "ash pikachu angry"},
{"url": "pokemon3.jpg", "description": "bulbasaur charizard squirtle"}]   
`

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("path:", r.URL.Path)
    w.Header().Set("Access-Control-Allow-Origin", "*")
    switch r.URL.Path {
    case "/Alice":
        fmt.Fprintf(w, Alice)
    case "/Bob":
        fmt.Fprintf(w, Bob)
    case "/Carol":
        fmt.Fprintf(w, Carol)
    case "/David":
        fmt.Fprintf(w, David)
    default:
        w.WriteHeader(http.StatusNotFound)    
    }
}

func main() {
    http.HandleFunc("/", handler)
    fs := http.FileServer(http.Dir(""))
    http.Handle("/images/", fs)
    fmt.Println("listening...")
    err := http.ListenAndServe(":8080", nil) 
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}