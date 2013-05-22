##ssws is a Simple Static Web Server written in [Go][1]##

It has 2 optional flags:

`-port` will set the TCP/IP port to listen on (defaults to port 80)

`-local` controls if the server is only listening on `localhost` (Set true or false, defaults to true)

It requires a directory names `www` to be in the same directory that the program is launched from.

###Example usage:###

`ssws -port=8080 -local=no` will run the server on TCP/IP port 8080 and will be accessible from all interfaces (depending on firewall and other network configuration).

`ssws` will run the server on TCP/IP port 80 and will only be accessible locally.

[1]: http://golang.org