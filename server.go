package main

import (
	"os"
	"strconv"

	"github.com/valyala/fasthttp"
)

func main() {
	host, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	hostbytes := []byte("hostname: " + host + "\ncounter: ")
	fasthttp.ListenAndServe(":80", func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Server", "php")
		ctx.Write(hostbytes)
		ctx.WriteString(strconv.FormatUint(GetCounter(), 10))
	})
}
