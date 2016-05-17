package main

import (
	"github.com/chrismrivera/cmd"
	"github.com/mikec/msplapi/client"
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
	func(cmd *cmd.Command, clt *client.Client) error {
		return handleClientResponse(clt.Deploy(
			cmd.Arg("build-handle"),
			cmd.Arg("project-handle"),
			cmd.Arg("provider"),
		))
	},
)
