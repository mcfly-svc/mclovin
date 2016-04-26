package main

import (
	"github.com/chrismrivera/cmd"
)

func init() {
	cmdr.AddCommand(login)
}

var login = cmd.NewCommand(

	"login", "Auth", "Login with github token",

	func(cmd *cmd.Command) {
		cmd.AppendArg("github_token", "GitHub auth token")
	},

	func(cmd *cmd.Command) error {

		cr, res, err := clt.Login(cmd.Arg("github_token"))
		if err != nil {
			return err
		}

		outputResponse(cr, res)

		return nil
	},
)
