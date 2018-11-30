package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type User struct {
	countryCode       int
	mno               int
	countryIdentifier string
	subscriber        int
}

type Args struct {
	msisdn int
}

type Parser struct{}

func (t *Parser) Extract(args *Args, reply *int) error {
	*reply = args.msisdn
	return nil
}

func main() {
	cal := new(Parser)
	server := rpc.NewServer()

	server.Register(cal)
	server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)

	listener, e := net.Listen("tcp", ":8080")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	fmt.Println("RPC server sstarting on localhost:8080")
	for {
		if conn, err := listener.Accept(); err != nil {
			log.Fatal("accept error: " + err.Error())
		} else {
			log.Printf("new connection established\n")
			go server.ServeCodec(jsonrpc.NewServerCodec(conn))
		}
	}

}
