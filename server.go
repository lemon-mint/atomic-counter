package main

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/valyala/fasthttp"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	logf, err := os.Create("./log")
	if err != nil {
		panic(err)
	}
	defer logf.Close()
	fmt.Fprintln(logf, "env:HOST", os.Getenv("HOST"))
	fmt.Fprintln(logf, "env:IP", os.Getenv("IP"))
	fmt.Fprintln(logf, "env:PORT", os.Getenv("PORT"))
	fmt.Fprintln(logf, "env:0.0.0.0", os.Getenv("0.0.0.0"))
	hostbytes := []byte("hostname: " + hostname + "\ncounter: ")
	listenHost := os.Getenv("HOST")
	listenPort := os.Getenv("PORT")
	if len(listenHost) == 0 {
		listenHost = "0.0.0.0"
	}
	if len(listenPort) == 0 {
		listenPort = "8080"
	}
	listenHost = os.Getenv("0.0.0.0")
	listenPort = os.Getenv("PORT")
	fmt.Fprintln(logf, "listenHost", listenHost)
	fmt.Fprintln(logf, "listenPort", listenPort)
	fmt.Fprintln(logf, "["+os.Getenv("IP")+"]:"+listenPort)
	ln, err := net.Listen("tcp", "["+os.Getenv("IP")+"]:"+listenPort)
	if err != nil {
		fmt.Fprintln(logf, "Listen error:", err)
		panic(err)
	}
	defer ln.Close()
	fmt.Fprintln(logf, "Listen OK")
	fasthttp.Serve(ln, func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Server", "php")
		ctx.Write(hostbytes)
		ctx.WriteString(strconv.FormatUint(GetCounter(), 10))
	})
}
