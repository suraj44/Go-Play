# TCP Server-Client

There are probably 100s of these out there but I think it was a good place to start.

For some extra fun, ask for a joke from the server as a client!

Commands recognized by the server:
1. `/quit`
2. `/time`
3. `/joke`

Instructions to run:
In two separate terminal tabs/windows run:
For the server:
```go run server.go -port 8000```

For the client:
```go run client.go -host localhost -port 8000 ```