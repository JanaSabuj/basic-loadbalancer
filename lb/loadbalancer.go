package lb

import (
	"fmt"
	"net/http"

	"github.com/JanaSabuj/server"
)

type LoadBalancer struct {
	Port            string
	Ip              string
	roundRobinCount int
	serverList      []server.Server
}

func NewLoadBalancer(serverList []server.Server) *LoadBalancer {
	return &LoadBalancer{
		Port:            "8000",
		Ip:              "http://www.localhost",
		roundRobinCount: 0,
		serverList:      serverList,
	}
}

func (lb *LoadBalancer) getRoundRobinServer() server.Server {
	server := lb.serverList[lb.roundRobinCount%len(lb.serverList)] // get the current available server from count
	if !server.IsAlive() {
		// if not alive, round robin to next
		lb.roundRobinCount++
		server = lb.serverList[lb.roundRobinCount%len(lb.serverList)]
	}
	lb.roundRobinCount++ // update the count
	return server
}

// main LB logic
func (lb *LoadBalancer) Serve(rw http.ResponseWriter, req *http.Request) {
	// 1. find the available server
	targetServer := lb.getRoundRobinServer()

	// 1a. log it
	fmt.Printf("Forwarding this request to backend server: %v\n", targetServer.Address())

	// 2. call the reverse proxy layer of the server
	targetServer.Serve(rw, req)
}
