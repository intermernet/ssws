###ssws is a Simple Static Web Server written in [Go][1]###

It's suitable for simple testing of html, css and javascript on your local machine. You can also use it for demonstrating ideas or mockups without a net connection, as a simple, single binary replacement to things like [Python's][2] `SimpleHTTPServer`, or just to play around with web development.

For more serious use, because it's so simple (59 lines of code), and written in [Go's][1] standard libraries, it should be quite reliable and performant, and may be a good choice as a static file server behind a proxy.

It has 4 optional flags:

`-port` will set the TCP/IP port to listen on (defaults to port 80)

`-local` controls if the server is only listening on `localhost` (Set true or false, defaults to true)

`-path` will set the directory to serve files from (defaults to "www" in the same directory, can be absolute or relative)

`-urlpath` will set the url path to export (defaults to "/")

####Example usage:####

`ssws` will run the server on TCP/IP port 80, will serve files from the directory `www` in the current directory and will only be accessible locally.

`ssws -port=8080` will run the server on TCP/IP port 8080, will serve files from the directory `www` in the current directory and will only be accessible at `http://localhost:8080/` .

`ssws -port=8080 -local=false -path=ht_docs` will run the server on TCP/IP port 8080, will serve files from the directory `ht_docs` in the current directory, and will be accessible from all interfaces (depending on firewall and other network configuration).

`ssws -port=8080 -path=/home/bob/ht_docs -urlpath=files` will run the server on TCP/IP port 8080, will serve files from the directory `/home/bob/ht_docs`, and will only be accessible at `http://localhost:8080/files/` .

####Downloads####
*~ 1.3MB*
- Linux ([64 bit tar.gz][3] | [32 bit tar.gz][4])
- Windows ([64 bit zip][5] | [32 bit zip][6])
- OSX ([64 bit zip][7] | [32 bit zip][8])

NOTE: This program has been tested on Windows 8 64bit, and Ubuntu 13.04

[1]: http://golang.org
[2]: http://python.org
[3]: http://downloads.intermer.net/ssws/bin/linux_amd64/ssws_linux_amd64.tar.gz
[4]: http://downloads.intermer.net/ssws/bin/linux_386/ssws_linux_386.tar.gz
[5]: http://downloads.intermer.net/ssws/bin/windows_amd64/ssws_windows_amd64.zip
[6]: http://downloads.intermer.net/ssws/bin/windows_386/ssws_windows_386.zip
[7]: http://downloads.intermer.net/ssws/bin/darwin_amd64/ssws_darwin_amd64.zip
[8]: http://downloads.intermer.net/ssws/bin/darwin_386/ssws_darwin_386.zip
[8]: http://downloads.intermer.net/ssws/bin/darwin_386/ssws_darwin_386.zip