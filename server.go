package main

import (
	"os"
	"strconv"

	"github.com/valyala/fasthttp"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	hostbytes := []byte("hostname: " + hostname + "\ncounter: ")
	listenHost := os.Getenv("HOST")
	listenPort := os.Getenv("PORT")
	var lnHost string
	if listenHost == "" && listenPort == "" {
		lnHost = ":80"
	} else if listenHost == "" && listenPort != "" {
		lnHost = ":" + listenPort
	} else if listenHost != "" && listenPort == "" {
		lnHost = listenHost
	} else {
		lnHost = listenHost + ":" + listenPort
	}
	fasthttp.ListenAndServe(lnHost, func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Server", "php")
		ctx.Write(hostbytes)
		ctx.WriteString(strconv.FormatUint(GetCounter(), 10))
	})
}
