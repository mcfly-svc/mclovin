package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"

	"github.com/mikec/marsupi-api/models"
	"github.com/chrismrivera/cmd"
)

func init() {
	cmdr.AddCommand(getProjects)
	cmdr.AddCommand(addProject)
}

var getProjects = cmd.NewCommand(

	"get-projects", "Projects", "Gets all projects",

	func(cmd *cmd.Command) { }, 

	func(cmd *cmd.Command) error {
		projects, _, err := api.GetProjects()
		if err != nil {
			return err
		}

		b, err := json.Marshal(projects)
		if err != nil {
			return err
		}
		fmt.Println(string(b))

		return nil
	},

)

var addProject = cmd.NewCommand(

	"add-project", "Projects", "Adds a new project",

	func(cmd *cmd.Command) {
		cmd.AppendArg("name", "Project name")
		cmd.AppendArg("username", "Project owner user name")
		cmd.AppendArg("service", "Service where the project lives [ github | bitbucket(unsupported) ]")
	}, 

	func(cmd *cmd.Command) error {
		p := models.Project{
			Name: cmd.Arg("name"),
			Username: cmd.Arg("username"),
			Service: cmd.Arg("service"),
		}

		pBytes, err := json.Marshal(p)
		if err != nil {
			return err
		}

		res, err := api.CreateProject(string(pBytes))
		if err != nil {
			return err
		}

		bBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		fmt.Println(string(bBytes))

		return nil
	},

)
