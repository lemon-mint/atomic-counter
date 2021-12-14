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
	hostbytes := []byte("hostname: " + hostname + "\ncounter: ")
	lnHost := ":8080"
	hostEnv := os.Getenv("HOST")
	if hostEnv != "" {
		lnHost = hostEnv
	}
	portEnv := os.Getenv("PORT")
	if portEnv != "" {
		lnHost = ":" + portEnv
	} else {
		portEnv = "8080"
	}
	ipEnv := os.Getenv("IP")
	if ipEnv != "" {
		ip := net.ParseIP(ipEnv)
		if ip != nil {
			lnHost = ipEnv + ":" + portEnv
		}
		if ip.To16() != nil {
			lnHost = "[" + ipEnv + "]:" + portEnv
		}
	}

	fmt.Fprintln(logf, lnHost)
	ln, err := net.Listen("tcp", lnHost)
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
