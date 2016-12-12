package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"runtime"
	"time"

)

var server = &http.Server{
	Addr:           ":80",
	Handler:        nil,
	ReadTimeout:    1000 * time.Second,
	WriteTimeout:   1000 * time.Second,
	MaxHeaderBytes: 1 << 20,
}

func main() {
	http.HandleFunc("/", requestHandler)
	fmt.Println("Server starts")
	fmt.Println()
	fmt.Println("Target URL's:")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Server wont started!!!")
		fmt.Println(err)
	}
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	switch r.URL.Path {
	case "/":
		http.ServeFile(w, r, "./content/html")
	case "/content/html/statistic.html":
		type statVars struct {
			Version string
			Os      string
			GoNum   int
		}
		p := statVars{Version: runtime.Version(), Os: runtime.GOOS, GoNum: runtime.NumGoroutine()}
		t, _ := template.ParseFiles("./content/templates/statistic.html")
		t.Execute(w, p)
	case "/stop server":
		os.Exit(0)
	default:
		http.ServeFile(w, r, r.URL.Path[1:])
	}
}
