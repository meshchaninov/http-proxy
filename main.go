package main

import (
	"log"
	"net/http"
	"os"

	"github.com/elazarl/goproxy"
	"github.com/elazarl/goproxy/ext/auth"
)

const PORT = ":80"

func main() {
	username := os.Getenv("PROXY_USER")
	password := os.Getenv("PROXY_PASS")

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	if username + password != "" {
		auth.ProxyBasic(proxy, "Auth", func(reqUser,reqPasswd string) bool {
			return username == reqUser && password == reqPasswd
		})
	}
	log.Println("Listen and serve")
	log.Fatal(http.ListenAndServe(PORT, proxy))

}