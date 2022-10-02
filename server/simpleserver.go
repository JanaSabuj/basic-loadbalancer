package server

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/JanaSabuj/customerror"
)

type SimpleServer struct {
	address string
	proxy   *httputil.ReverseProxy
}

// Factory
func NewSimpleServer(address string) *SimpleServer {
	url, err := url.Parse(address)
	if err != nil {
		customerror.HandleError(err)
	}
	return &SimpleServer{
		address: address,
		proxy:   httputil.NewSingleHostReverseProxy(url), // proxy that points to the actual address
	}
}

func (ss *SimpleServer) Address() string {
	return ss.address
}

func (ss *SimpleServer) IsAlive() bool {
	return true
}

func (ss *SimpleServer) Serve(rw http.ResponseWriter, req *http.Request) {
	ss.proxy.ServeHTTP(rw, req)
}
