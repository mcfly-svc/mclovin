package main

import (
	"github.com/chrismrivera/cmd"
	"github.com/mcfly-svc/mcfly/api/apidata"
	"github.com/mcfly-svc/mcfly/client"
)

func init() {
	cmdr.AddCommand(deploy)
}

var deploy = NewAuthCommand(
	"deploy", "Builds", "Starts a deploy for a build",
	func(cmd *cmd.Command) {
		cmd.AppendArg("build-handle", `A handle that uniquely identifies the build`)
		cmd.AppendArg("project-handle", `A handle that uniquely identifies the project`)
		cmd.AppendArg("provider", "Provider (github, dropbox, ...)")
	},
	func(cmd *cmd.Command, clt client.Client) error {
		return handleClientResponse(clt.Deploy(&apidata.DeployReq{
			BuildHandle:         cmd.Arg("build-handle"),
			SourceProjectHandle: cmd.Arg("project-handle"),
			Provider:            cmd.Arg("provider"),
		}))
	},
)
