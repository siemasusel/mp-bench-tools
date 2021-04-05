package main

import (
	"fmt"
	"net"
	"time"

	"github.com/alecthomas/kong"
)

var Config struct {
	Concurrency  int           `short:"c" default:"1" help:"Number of multiple requests to perform at a time."`
	Protocol     string        `short:"P" default:"TCP" enum:"TCP" help:"Protocol to use during benchmark. Currently only TCP is supported."`
	Destinations []net.IP      `short:"d" help:"Connect to destination IP addresses. If not specified program will only run only as a server."`
	Sources      []net.IP      `short:"s" help:"Bind to host IP addresses. If not specified program will run only as a client."`
	Port         int           `short:"p" default:"3030" help:"The server port for the server to listen on and the client to connect to (this should be the same in both client and server)."`
	TimeLimit    time.Duration `short:"t" help:"Maximum number of time to spend for benchmarking. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'. Per default there is no timelimit."`
	WindowSize   int           `short:"b" default:"1024" help:"Size of TCP send/receive buffer, in bytes."`
}

func main() {
	_ = kong.Parse(&Config)
	fmt.Println(Config.Sources)
	fmt.Println(Config.WindowSize)
}
