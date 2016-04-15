package main

import (
	"github.com/mikec/marsupi-api/apiutil"
	"github.com/chrismrivera/cmd"
	"fmt"
	"os"
)

var cmdr *cmd.App = cmd.NewApp()
var api = apiutil.ApiUtil{"http://localhost:8080"}

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
