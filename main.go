package main

import (
	"fmt"
	"net/http"

	"github.com/JanaSabuj/lb"
	"github.com/JanaSabuj/server"
)

func main() {
	fmt.Println("Hello to basic-elb")

	// 1. configure the list of backend servers
	servers := []server.Server{
		server.NewSimpleServer("http://www.janasabuj.github.io"),
		server.NewSimpleServer("http://www.duckduckgo.com"),
		server.NewSimpleServer("http://www.bing.com"),
	}

	// 2. create a new LB with serverList
	lb := lb.NewLoadBalancer(servers)

	// forwardToLB
	forwardToLB := func(rw http.ResponseWriter, req *http.Request) {
		lb.Serve(rw, req)
	}

	// 3. wait for requests at the main route
	http.HandleFunc("/", forwardToLB)

	// 4. configure the logs
	fmt.Printf("Serving requests at %s:%s \n", lb.Ip, lb.Port)

	// 5. start the server
	http.ListenAndServe(":"+lb.Port, nil)
}
