package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strconv"
)

// User : model to return
type User struct {
	CountryCode       int    `json:"countryCode"`
	Mno               int    `json:"mno"`
	CountryIdentifier string `json:"counrtyIdentifier"`
	Subscriber        int    `json:"subscriber"`
}

// Args : model for arguments
type Args struct {
	Msisdn int
}

// Parser : model with funcs
type Parser struct{}

func toStr(num int) string {
	s := strconv.Itoa(num)
	return s
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Print("")
	}
	return i
}

func getCC(num int) int {
	var cc int
	snum := toStr(num)
	switch len(snum) {
	case 13:
		cc = toInt(snum[:3])
	case 12:
		cc = toInt(snum[:2])
	case 11:
		cc = toInt(snum[:1])
	default:
	}
	return cc
}

// Extract : parse msisdn and return User
func (t *Parser) Extract(args *Args, reply *User) error {
	input := args.Msisdn

	chpok := User{
		CountryCode:       01,
		Mno:               15,
		CountryIdentifier: "test",
		Subscriber:        1509,
	}

	fmt.Printf("Req: %d\n", input)
	// fmt.Println("Res:", chpok)
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
