/*
 * @author Daniel Popov <lalabuy9948@gmail.com>
 * @copyright <Do whatever you want>
 */

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
	Mno               []string `json:"mno"`
	DialingCode       string   `json:"dialingCode"`
	Subscriber        string   `json:"subscriber"`
	CountryIdentifier string   `json:"counrtyIdentifier"`
	CountryName       string   `json:"countryName"`
}

// Args : model for arguments
type Args struct {
	Msisdn int
}

// Parser : model for funcs
type Parser struct{}

// ToStr : convert int to string
func ToStr(num int) string {
	s := strconv.Itoa(num)
	return s
}

// ToInt : convert string to int
func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	return i
}

func getCC(num int) string {
	var cc string
	snum := ToStr(num)
	switch len(snum) {
	case 14:
		cc = snum[:4]
	case 13:
		cc = snum[:3]
	case 12:
		cc = snum[:2]
	case 11:
		cc = snum[:1]
	default:
		cc = "000"
	}
	return cc
}

func getSubscriber(num int) string {
	var cc string
	snum := ToStr(num)
	switch len(snum) {
	case 14:
		cc = snum[8:]
	case 13:
		cc = snum[7:]
	case 12:
		cc = snum[6:]
	case 11:
		cc = snum[5:]
	default:
		cc = "000000"
	}
	return cc
}

func getContryName() {
	FindCountry("004")
}

// Extract : parse msisdn and return User
func (t *Parser) Extract(args *Args, reply *User) error {
	input := args.Msisdn

	cc := getCC(input)
	alpha := GetAlpha(cc)
	country := FindCountry(alpha)
	countryName := country.Name
	mnos := FindMNO(countryName)
	countryIdentifier := country.CC1
	dialingCode := cc

	chpok := User{
		Mno:               mnos,
		DialingCode:       dialingCode,
		CountryIdentifier: countryIdentifier,
		CountryName:       countryName,
		Subscriber:        getSubscriber(input),
	}

	fmt.Printf("Req: %d\n", input)
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
