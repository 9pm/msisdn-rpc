package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

type Args struct {
	msisdn int
}

func main() {

	client, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	args := &Args{918369110173}
	var reply int
	c := jsonrpc.NewClient(client)

	err = c.Call("Parser.Extract", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}

	fmt.Printf("Req/Res: %d | %d\n", args.msisdn, reply)
}
