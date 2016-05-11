package main

import (
	"log"

	"github.com/chrismrivera/cmd"
	"github.com/mikec/msplapi/client"
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

	func(cmd *cmd.Command, clt *client.Client) error {

		cr, res, err := clt.AddProject(cmd.Arg("project-handle"), cmd.Arg("provider"))
		if err != nil {
			log.Fatal(err)
			return err
		}

		outputResponse(cr, res)

		return nil
	},
)

var getProviderProjects = NewAuthCommand(

	"get-provider-projects", "Projects", "Gets all projects that the authenticated user owns on a given provider",

	func(cmd *cmd.Command) {
		cmd.AppendArg("provider", "Project source provider (github, dropbox, ...)")
	},

	func(cmd *cmd.Command, clt *client.Client) error {

		cr, res, err := clt.GetProviderProjects(cmd.Arg("provider"))
		if err != nil {
			log.Fatal(err)
			return err
		}

		outputResponse(cr, res)

		return nil
	},
)

var getProjects = NewAuthCommand(

	"get-projects", "Projects", "Gets all projects that the authenticated user has added to marsupi",

	func(cmd *cmd.Command) {},

	func(cmd *cmd.Command, clt *client.Client) error {

		cr, res, err := clt.GetProjects()
		if err != nil {
			log.Fatal(err)
			return err
		}

		outputResponse(cr, res)

		return nil
	},
)

var deleteProject = NewAuthCommand(

	"delete-project", "Projects", "Deletes a project",

	func(cmd *cmd.Command) {
		cmd.AppendArg("project-handle", `A handle that uniquely identifies the project`)
		cmd.AppendArg("provider", "Provider (github, dropbox, ...)")
	},

	func(cmd *cmd.Command, clt *client.Client) error {

		cr, res, err := clt.DeleteProject(cmd.Arg("project-handle"), cmd.Arg("provider"))
		if err != nil {
			log.Fatal(err)
			return err
		}

		outputResponse(cr, res)

		return nil
	},
)
