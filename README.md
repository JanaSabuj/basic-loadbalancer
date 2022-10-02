# basic-loadbalancer
- A basic load balancer written in Go 

- This LB load balances between 3 hardcoded servers
- Clone locally and run
``` go run main.go```
- Go to ```localhost:8000``` in incognito window and start hitting the ```/``` route
- You can see that the LB round robins between the 3 hardcoded servers
