# MSISDN RPC Server

1. Create an application with following requirements:

- latest PHP or Golang
- takes MSISDN as an input
- returns MNO identifier, country dialling code, subscriber number and country identifier as defined with ISO 3166-1-alpha-2
- do not care about number portability

2. Write all needed tests.

3. Expose the package through an RPC API, select one and explain why have you chosen it.

4. Use git, vagrant and/or docker, and a configuration management tool (puppet, chef, ansible ...).

5. Other:

- a git repository with full commit history is expected to be part of the delivered solution
- if needed, provide additional installation instructions, but there shouldn't be much more than running a simple command to set everything up
- use best practices all around. For PHP, good source of that would be http://www.phptherightway.com

```go
MSISDN = CC + NPA + SN

CC = Country Code

NPA = Number Planning Area

SN = Subscriber Number
```

## How to

```sh
git clone git@github.com:lalabuy948/msisdn-rpc.git --depth=0
cd msisdn-rpc

# server
go run *.go

# client
go run client/client.go
```

## RPC

I chose [JSON-RPC 2.0](https://golang.org/pkg/net/rpc/jsonrpc/) because of JSON simplicity and it's standard built in library.

`method: parse`

`parameters: msisdn`