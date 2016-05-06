package main

import (
	"fmt"

	"github.com/chrismrivera/cmd"
	"github.com/mikec/msplapi/client"
)

func init() {
	cmdr.AddCommand(addProject)
}

var addProject = NewAuthCommand(

	"add-project", "Projects", "Add a new project",

	func(cmd *cmd.Command) {
		cmd.AppendArg("project-name", `Name of the project to add. 
																		Must match the name of the project on the 
																		given provider`)
		cmd.AppendArg("provider", "Provider (github, dropbox, ...)")
	},

	func(cmd *cmd.Command, clt *client.Client) error {

		/*cr, res, err := clt.Login(cmd.Arg("project-name"), cmd.Arg("provider"))
		if err != nil {
			log.Fatal(err)
			return err
		}

		outputResponse(cr, res)*/

		fmt.Println("TOKEN:", clt.Context.AccessToken)

		return nil
	},
)
