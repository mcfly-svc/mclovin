package main

import (
	"github.com/chrismrivera/cmd"
	"github.com/mcfly-svc/mcfly/api/apidata"
	"github.com/mcfly-svc/mcfly/client"
)

func init() {
	cmdr.AddCommand(addProject)
	cmdr.AddCommand(getProviderProjects)
	cmdr.AddCommand(getProjects)
	cmdr.AddCommand(deleteProject)
}

var addProject = NewAuthCommand(
	"add-project", "Projects", "Add a new project",
	func(cmd *cmd.Command) {
		cmd.AppendArg("project-handle", `A handle that uniquely identifies the project`)
		cmd.AppendArg("provider", "Provider (github, dropbox, ...)")
	},
	func(cmd *cmd.Command, clt client.Client) error {
		cr, res, err := clt.AddProject(&apidata.ProjectReq{
			Handle:   cmd.Arg("project-handle"),
			Provider: cmd.Arg("provider"),
		})
		return handleClientResponse(cr, res, err)
	},
)

var getProviderProjects = NewAuthCommand(
	"get-provider-projects", "Projects", "Gets all projects that the authenticated user owns on a given provider",
	func(cmd *cmd.Command) {
		cmd.AppendArg("provider", "Project source provider (github, dropbox, ...)")
	},
	func(cmd *cmd.Command, clt client.Client) error {
		return handleClientResponse(clt.GetProviderProjects(cmd.Arg("provider")))
	},
)

var getProjects = NewAuthCommand(
	"get-projects", "Projects", "Gets all projects that the authenticated user has added to mcfly",
	func(cmd *cmd.Command) {},
	func(cmd *cmd.Command, clt client.Client) error {
		return handleClientResponse(clt.GetProjects())
	},
)

var deleteProject = NewAuthCommand(
	"delete-project", "Projects", "Deletes a project",
	func(cmd *cmd.Command) {
		cmd.AppendArg("project-handle", `A handle that uniquely identifies the project`)
		cmd.AppendArg("provider", "Provider (github, dropbox, ...)")
	},
	func(cmd *cmd.Command, clt client.Client) error {
		return handleClientResponse(clt.DeleteProject(&apidata.ProjectReq{
			Handle:   cmd.Arg("project-handle"),
			Provider: cmd.Arg("provider"),
		}))
	},
)
