package main

import (
	"github.com/chrismrivera/cmd"
	"github.com/mikec/msplapi/api"
	"github.com/mikec/msplapi/client"
)

func init() {
	cmdr.AddCommand(login)
	cmdr.AddCommand(logout)
}

var login = cmd.NewCommand(

	"login", "Auth", "Login with a provider token",

	func(cmd *cmd.Command) {
		cmd.AppendArg("token", "Provider auth token")
		cmd.AppendArg("provider", "Provider (github, dropbox, ...)")
	},

	func(cmd *cmd.Command) error {
		clt := client.NewClient(cfg.ApiUrl, "")

		cr, res, err := clt.Login(cmd.Arg("token"), cmd.Arg("provider"))

		s, err := NewSimpleCredentialStore()
		if err != nil {
			return err
		}

		u := cr.Data.(*api.LoginResp)

		err = s.SaveUserCreds(&UserCreds{
			Token: u.AccessToken,
		})
		if err != nil {
			return err
		}

		return handleClientResponse(cr, res, err)
	},
)

var logout = cmd.NewCommand("logout", "Auth", "Logout from msplapi",
	func(cmd *cmd.Command) {},

	func(cmd *cmd.Command) error {
		s, err := NewSimpleCredentialStore()
		if err != nil {
			return err
		}

		err = s.ClearUserCreds()
		if err != nil {
			return err
		}

		return nil
	},
)
