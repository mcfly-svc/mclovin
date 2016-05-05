package main

import (
	"log"

	"github.com/chrismrivera/cmd"
)

func init() {
	cmdr.AddCommand(login)
}

var login = cmd.NewCommand(

	"login", "Auth", "Login with github token",

	func(cmd *cmd.Command) {
		cmd.AppendArg("token", "Provider auth token")
		cmd.AppendArg("token_type", "Provider (github, dropbox, ...)")
	},

	func(cmd *cmd.Command) error {
		cr, res, err := clt.Login(cmd.Arg("token"), cmd.Arg("token_type"))
		if err != nil {
			log.Fatal(err)
			return err
		}

		outputResponse(cr, res)

		return nil
	},
)
