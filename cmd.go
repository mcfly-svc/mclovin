package main

import (
	"github.com/mikec/marsupi-api/client"
	"github.com/chrismrivera/cmd"
	"fmt"
	"encoding/json"
	"os"
)

var cmdr *cmd.App = cmd.NewApp()
var clt = client.NewClient("http://localhost:8080")

func main() {
	cmdr.Description = "A command line client for the marsupi API"
	if err := cmdr.Run(os.Args); err != nil {
		if ue, ok := err.(*cmd.UsageErr); ok {
			ue.ShowUsage()
		} else {
			fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
		}
		os.Exit(1)
	}
}

func outputResponse(res *client.ClientResponse) {
	b, err := json.Marshal(res.Data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println()
	fmt.Printf("StatusCode: %d\n", res.StatusCode)
	fmt.Printf("Data:       %s\n", string(b))
	fmt.Println()
}