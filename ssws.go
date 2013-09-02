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
	"strconv"
	"strings"
)

const (
	ps  = os.PathSeparator
	ups = "/"
)

var port int
var local bool
var path string
var urlpath string

func init() {
	flag.IntVar(&port, "port", 80, "TCP/IP Port to listen on")
	flag.BoolVar(&local, "local", true, "Listen on localhost only")
	flag.StringVar(&path, "path", "www", "Path to serve files from")
	flag.StringVar(&urlpath, "urlpath", "/", "URL Path to export")
}

func main() {
	flag.Parse()
	p := strconv.Itoa(port)
	addr := "localhost:" + p
	if local != true {
		addr = ":" + p
	}
	ps := string(ps)
	if !strings.HasSuffix(path, ps) {
		path = path + ps
	}
	if !strings.HasPrefix(urlpath, ups) {
		urlpath = ups + urlpath
	}
	if !strings.HasSuffix(urlpath, ups) {
		urlpath = urlpath + ups
	}
	http.Handle(urlpath, http.StripPrefix(urlpath, http.FileServer(http.Dir(path))))
	log.Printf("\nServing files from %s on TCP/IP port %d\nlocalhost only=%t\nURL Path=%s\n", path, port, local, urlpath)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
