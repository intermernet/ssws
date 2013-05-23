/*
Copyright Mike Hughes 2013 (intermernet AT gmail DOT com)

ssws is a Simple Static Web Server

LICENSE: BSD 3-Clause License (see http://opensource.org/licenses/BSD-3-Clause)
*/

package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strings"
)

var port string
var local bool
var path string
var urlpath string

func init() {
	flag.StringVar(&port, "port", "80", "TCP/IP Port to listen on")
	flag.BoolVar(&local, "local", true, "Listen on localhost only")
	flag.StringVar(&path, "path", "www", "Path to serve files from")
	flag.StringVar(&urlpath, "urlpath", "/", "URL Path to export")
}

func main() {
	flag.Parse()
	addr := "localhost:" + port
	if local != true {
		addr = ":" + port
	}
	ps := string(os.PathSeparator)
	if !strings.HasSuffix(path, ps) {
		path = path + ps
	}
	if !strings.HasPrefix(urlpath, "/") {
		urlpath = "/" + urlpath
	}
	http.Handle(urlpath, http.StripPrefix(urlpath, http.FileServer(http.Dir(path))))
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
