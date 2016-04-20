package main

import (
	"github.com/mikec/marsupi-api/client"
	"github.com/chrismrivera/cmd"
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
)

var cmdr *cmd.App = cmd.NewApp()
var apiClient = client.NewClient("http://localhost:8080")

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

func outputResponse(res *http.Response) {
	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(contents))
}