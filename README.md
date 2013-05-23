###ssws is a Simple Static Web Server written in [Go][1]###

It has 3 optional flags:

`-port` will set the TCP/IP port to listen on (defaults to port 80)

`-path` will set the path to serve (defaults to "www" in the same directory, can be absolute or relative)

`-local` controls if the server is only listening on `localhost` (Set true or false, defaults to true)

It requires a directory named `www` to be in the same directory that the program is launched from.
Your site structure should go into the `www` folder.

####Example usage:####

`ssws -port=8080 -path=ht_docs -local=false` will run the server on TCP/IP port 8080, will serve the directory `ht_docs` in the current directory, and will be accessible from all interfaces (depending on firewall and other network configuration).

`ssws -port=8080 -local=true` will run the server on TCP/IP port 8080 and will only be accessible locally.

`ssws` will run the server on TCP/IP port 80 and will only be accessible locally.

[1]: http://golang.org