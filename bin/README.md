###ssws is a Simple Static Web Server written in [Go][1]###

It has 4 optional flags:

`-port` will set the TCP/IP port to listen on (defaults to port 80)

`-local` controls if the server is only listening on `localhost` (Set true or false, defaults to true)

`-path` will set the path to serve (defaults to "www" in the same directory, can be absolute or relative)

`-urlpath` will set the url path to export (defaults to "/")

####Example usage:####

`ssws` will run the server on TCP/IP port 80 and will only be accessible locally.

`ssws -port=8080 -local=true` will run the server on TCP/IP port 8080 and will only be accessible at `http://localhost/` .

`ssws -port=8080 -local=false -path=ht_docs` will run the server on TCP/IP port 8080, will serve the directory `ht_docs` in the current directory, and will be accessible from all interfaces (depending on firewall and other network configuration).

`ssws -port=8080 -path=ht_docs -urlpath=files` will run the server on TCP/IP port 8080, will serve the directory `ht_docs` in the current directory, and will only be accessible at `http://localhost/files/` .

[1]: http://golang.org