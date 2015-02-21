package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/glycerine/goproxy"
)

func main() {

	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	verbose := flag.Bool("v", false, "should every proxy request be logged to stdout")
	addr := flag.String("addr", ":8080", "proxy listen address")
	flag.Parse()
	proxy := goproxy.NewProxyHttpServer()
	if verbose != nil && *verbose {
		fmt.Printf("\n being verbose!\n")
		proxy.Verbose = *verbose
	}
	log.Fatal(http.ListenAndServe(*addr, proxy))
	fmt.Printf("\n basic example done.\n")
}
