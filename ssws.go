package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strings"
)

var port string
var path string
var local bool

func init() {
	flag.StringVar(&port, "port", "80", "TCP/IP Port to listen on")
	flag.StringVar(&path, "path", "www", "Path to export")
	flag.BoolVar(&local, "local", true, "Listen on localhost only")
}

func main() {
	flag.Parse()
	ps := string(os.PathSeparator)
	adr := "localhost:" + port
	if !strings.HasSuffix(path, ps) {
		path = path + ps
	}
	if local != true {
		adr = ":" + port
	}
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(path))))
	if err := http.ListenAndServe(adr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
