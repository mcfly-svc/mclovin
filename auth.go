package main

import (
	"github.com/chrismrivera/cmd"
	"github.com/mcfly-svc/mcfly/api/apidata"
	"github.com/mcfly-svc/mcfly/client"
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
		clt := client.NewMcflyClient(cfg.ApiUrl, "")

		cr, res, err := clt.Login(&apidata.LoginReq{
			Token:    cmd.Arg("token"),
			Provider: cmd.Arg("provider"),
		})
		if err != nil {
			panic(err)
			return err
		}

		s, err := NewSimpleCredentialStore()
		if err != nil {
			panic(err)
			return err
		}

		u := cr.Data.(*apidata.LoginResp)

		err = s.SaveUserCreds(&UserCreds{
			Token: u.AccessToken,
		})
		if err != nil {
			panic(err)
			return err
		}

		return outputResponse(cr, res)
	},
)

var logout = cmd.NewCommand(
	"logout", "Auth", "Logout from mcflyapi",
	func(cmd *cmd.Command) {},
	func(cmd *cmd.Command) error {
		s, err := NewSimpleCredentialStore()
		if err != nil {
			return err
		}
		return s.ClearUserCreds()
	},
)
