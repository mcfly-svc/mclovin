// ADD COMMANDS FOR USER

package main

import (
	"github.com/mikec/marsupi-api/client"
	"github.com/chrismrivera/cmd"
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
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

func outputResponse(cr *client.ClientResponse, res *http.Response) {

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println()
	fmt.Printf("StatusCode:          %d\n", cr.StatusCode)
	fmt.Printf("Data:                %s\n", fmt.Sprintf("%+v", cr.Data))
	fmt.Printf("Body:                %s\n", string(body))
	fmt.Println()
}