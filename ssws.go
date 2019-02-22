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
	"path"
	"strconv"
	"strings"
)

const (
	osps = os.PathSeparator
	ups  = "/"
)

var port int
var local bool
var filepath string
var urlpath string

func init() {
	flag.IntVar(&port, "port", 80, "TCP/IP Port to listen on")
	flag.BoolVar(&local, "local", true, "Listen on localhost only")
	flag.StringVar(&filepath, "path", "www", "Path to serve files from")
	flag.StringVar(&urlpath, "urlpath", "/", "URL Path to export")
}

func main() {
	flag.Parse()
	p := strconv.Itoa(port)
	addr := "localhost:" + p
	if local != true {
		addr = ":" + p
	}
	ps := string(osps)
	if !strings.HasSuffix(filepath, ps) {
		filepath = filepath + ps
	}
	if !strings.HasPrefix(urlpath, ups) {
		urlpath = ups + urlpath
	}
	if !strings.HasSuffix(urlpath, ups) {
		urlpath = urlpath + ups
	}

	http.Handle(urlpath, http.StripPrefix(urlpath, changeHeader(http.FileServer(http.Dir(filepath)))))
	log.Printf("\nServing files from %s on TCP/IP port %d\nlocalhost only=%t\nURL Path=%s\n", filepath, port, local, urlpath)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func changeHeader(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		switch path.Ext(r.URL.Path) {
		case ".js":
			w.Header().Add("Content-Type", "application/javascript")
		}
		// Set some header.
		//w.Header().Add("Keep-Alive", "300")
		// Serve with the actual handler.
		h.ServeHTTP(w, r)
	}
}
