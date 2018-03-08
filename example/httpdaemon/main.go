package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"

	"github.com/haxii/daemon"
)

var (
	port    = flag.Int("p", 8080, "server port")
	service = flag.String("s", daemon.UsageDefaultName, daemon.UsageMessage)
)

func main() {
	d := daemon.Make("-s", "httpdaemon", "simple http daemon service")

	d.Run(func() {
		flag.Parse()
		http.HandleFunc("/hello", helloServer)
		http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	})
}

func helloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}
