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
	Mno               []string `json:"mno"`
	DialingCode       string   `json:"dialingCode"`
	Subscriber        string   `json:"subscriber"`
	CountryIdentifier string   `json:"counrtyIdentifier"`
	CountryName       string   `json:"countryName"`
}

func main() {

	client, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	// India 918369110173
	// Russia 78369111222
	// American Samoa 16848369555888
	args := &Args{16848369555888}
	var reply User
	c := jsonrpc.NewClient(client)
	err = c.Call("Parser.Extract", args, &reply)
	if err != nil {
		log.Fatal("parse error:", err)
	}

	fmt.Printf("MSISDN: %d\n\n", args.Msisdn)

	fmt.Println("MNO: ", reply.Mno)
	fmt.Printf("Dialing code: +%s\n", reply.DialingCode)
	fmt.Printf("County identifier: %s\n", reply.CountryIdentifier)
	fmt.Printf("County name: %s\n", reply.CountryName)
	fmt.Printf("Subscriber: %s\n", reply.Subscriber)
}
