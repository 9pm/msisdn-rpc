package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// User : model to return
type User struct {
	countryCode       int
	mno               int
	countryIdentifier string
	subscriber        int
}

// Args : model for arguments
type Args struct {
	msisdn int
}

// Parser : model with funcs
type Parser struct{}

// Extract : parse msisdn and return User
func (t *Parser) Extract(args *Args, reply *User) error {
	chelebok := args.msisdn
	chpok := User{
		countryCode:       01,
		mno:               15,
		countryIdentifier: "test",
		subscriber:        1509,
	}
	fmt.Printf("Req: %d\n", chelebok)
	fmt.Println(chpok)
	*reply = chpok
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

	fmt.Println("RPC server starting on localhost:8080")
	for {
		if conn, err := listener.Accept(); err != nil {
			log.Fatal("accept error: " + err.Error())
		} else {
			log.Println("new connection established")
			go server.ServeCodec(jsonrpc.NewServerCodec(conn))
		}
	}

}
