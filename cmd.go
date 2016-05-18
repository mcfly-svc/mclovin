package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/chrismrivera/cmd"
	"github.com/mikec/msplapi/client"
	"github.com/mikec/msplapi/config"
)

var cmdr *cmd.App = cmd.NewApp()
var cfg *config.Config

type AuthCommandRunFunc func(cmd *cmd.Command, ac client.Client) error

var ErrNoCredentials = errors.New("Please login first.")

func main() {
	_cfg, err := config.NewConfigFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}
	cfg = _cfg
	cmdr.Description = "A command line client for the marsupi API"
	if err = cmdr.Run(os.Args); err != nil {
		if ue, ok := err.(*cmd.UsageErr); ok {
			ue.ShowUsage()
		} else {
			fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
		}
		os.Exit(1)
	}
}

func NewAuthCommand(name, group, desc string, setup cmd.SetupFunc, run AuthCommandRunFunc) *cmd.Command {
	wrappingSetup := func(cmd *cmd.Command) {
		setup(cmd)
		cmd.AddFlag("token", "", "Override the authentication token")
	}

	wrappingRun := func(cmd *cmd.Command) error {
		s, err := NewSimpleCredentialStore()
		if err != nil {
			return err
		}

		u, err := s.GetUserCreds()
		if err != nil {
			return err
		}

		overrideToken := cmd.Flag("token")

		if u == nil {
			if overrideToken != "" {
				u = &UserCreds{}
			} else {
				return ErrNoCredentials
			}
		}

		if overrideToken != "" {
			u.Token = overrideToken
		}

		clt := client.NewMsplClient(cfg.ApiUrl, u.Token)

		return run(cmd, clt)
	}

	return cmd.NewCommand(name, group, desc, wrappingSetup, wrappingRun)
}

func handleClientResponse(cr *client.ClientResponse, res *http.Response, err error) error {
	if err != nil {
		log.Fatal(err)
	}
	return outputResponse(cr, res)
}

func outputResponse(cr *client.ClientResponse, res *http.Response) error {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
		return err
	}

	fmt.Println()
	if cr != nil {
		fmt.Printf("StatusCode:          %d\n", cr.StatusCode)
		//fmt.Printf("Data:                %s\n", fmt.Sprintf("%+v", cr.Data))
	}
	fmt.Printf("Body:                %s\n", string(body))
	fmt.Println()
	return nil
}
