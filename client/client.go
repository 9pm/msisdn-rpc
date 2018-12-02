/*
 * @author Daniel Popov <lalabuy9948@gmail.com>
 * @copyright <Do whatever you want>
 */

package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

// Args : model for arguments
type Args struct {
	Msisdn int
}

// User : model to return
type User struct {
	CountryCode       int    `json:"countryCode"`
	Mno               int    `json:"mno"`
	CountryIdentifier string `json:"counrtyIdentifier"`
	Subscriber        string `json:"subscriber"`
}

func main() {

	client, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	args := &Args{6438369110173}
	var reply User
	c := jsonrpc.NewClient(client)
	err = c.Call("Parser.Extract", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}

	fmt.Printf("Req/Res: %d\n", args.Msisdn)

	fmt.Printf("County code: %d\n", reply.CountryCode)
	fmt.Printf("MNO: %d\n", 0)
	fmt.Printf("County identifier: %s\n", reply.CountryIdentifier)
	fmt.Printf("Subscriber: %s\n", reply.Subscriber)
}
