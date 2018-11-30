package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

// Args : model for arguments
type Args struct {
	msisdn int
}

// User : model to return
type User struct {
	countryCode       int
	mno               int
	countryIdentifier string
	subscriber        int
}

func main() {

	client, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	args := &Args{918369110173}
	var reply User
	c := jsonrpc.NewClient(client)
	err = c.Call("Parser.Extract", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}

	fmt.Printf("Req/Res: %d\n", args.msisdn)
	fmt.Println(reply)
}
