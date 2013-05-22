// +build windows linux darwin
// +build 386 amd64

package main

import "net/http"
import "log"
import "flag"

var port string
var local bool

func init() {
	flag.StringVar(&port, "port", "80", "TCP/IP Port to listen on")
	flag.BoolVar(&local, "local", true, "Listen on localhost only")
}

func main() {

	flag.Parse()

	adr := "localhost:" + port

	if local != true {
		adr = ":" + port
	}

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("www/"))))
	if err := http.ListenAndServe(adr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
