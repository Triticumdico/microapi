package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/Triticumdico/micro-api/args"
	"github.com/Triticumdico/micro-api/handler"
	"github.com/spf13/pflag"
)

var (
	argInsecurePort        = pflag.Int("insecure-port", 9090, "port to listen to for incoming HTTP requests")
	argInsecureBindAddress = pflag.IP("insecure-bind-address", net.IPv4(127, 0, 0, 1), "IP address on which to serve the --insecure-port, set to 127.0.0.1 for all interfaces")
)

func main() {

	// Set logging output to standard console out
	log.SetOutput(os.Stdout)

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	// Initializes dashboard arguments holder so we can read them in other packages
	initArgHolder()

	//TODO Client check connection
	log.Printf("No client connection tested")

	//TODO Create Http API Handler
	//apiHandler, err := handler.CreateHTTPAPIHandler()

	// Run a HTTP server that serves static public files from handles API calls.
	http.HandleFunc("/", handler.LocalPage.GetHomePage)

	// Listen for http
	log.Printf("Serving insecurely on HTTP port: %d", args.Holder.GetInsecurePort())
	addr := fmt.Sprintf("%s:%d", args.Holder.GetInsecureBindAddress(), args.Holder.GetInsecurePort())
	go func() { log.Fatal(http.ListenAndServe(addr, nil)) }()

	select {}

}

func initArgHolder() {
	builder := args.GetHolderBuilder()
	builder.SetInsecurePort(*argInsecurePort)
	builder.SetInsecureBindAddress(*argInsecureBindAddress)
}
