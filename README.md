calc-svc
-------------

[![calc-svc](https://circleci.com/gh/hazbo/calc-svc.svg?style=svg)](https://circleci.com/gh/hazbo/calc-svc)

A simple calculator as a Go microservice.

### Building the client

```bash
$ go build -o calc-client cmd/calc-client/main.go
```

You'll now see the client binary at `./calc-client`. Don't run it just yet though.

### Building the server

Building the server can be useful for local development or for building a new
container image for this service. To quickly test it out with the client, you
may build like so:

```bash
$ go build -o calc-server cmd/calc-server/main.go
```

Open a new terminal window and start the server there:

```bash
$ ./calc-server
```

In your original terminal window, start the client:

```bash
$ ./calc-client sum 10 10
$ ./calc-client mul 30 40
$ ./calc-client div 100 5
```

You may also define the port and host in each of these programs. By default the
server listens on port 2020.

This can be changed like so:

```bash
$ ./calc-server -p 3000
```

And for the client:

```bash
$ ./calc-client -r 127.0.0.1:3000 mul 50 50
```

### Running on Minikube

Start Minikube if you haven't already

```
$ minikube start
```

Open up a new terminal window and start the Minikube tunnel so that services
can be assigned an external IP.

```bash
$ minikube tunnel
```

Keep that running and switch back to your original terminal window. Next, we'll
create the service and deployment.

```bash
$ kubectl apply -f deployments/minikube
```

You may inspect the newly created deployment and service.

```bash
# For example...

$ kubectl get pods
$ kubectl get deployments
$ kubectl get services
```

Find the external IP for the newly created service.

```bash
$ kubectl get services | grep calc\-service
```

Test against the client with the newly created external IP.

```bash
$ ./calc-client -remote 10.102.113.144:2020 add 20 40
```

given that in the above example `10.102.113.144` is the IP address generated
for us. Yours may be different, though `:2020` is the default port unless you
specify otherwise.
