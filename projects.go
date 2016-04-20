package main

import (
	"encoding/json"
	"strconv"

	"github.com/mikec/marsupi-api/models"
	"github.com/chrismrivera/cmd"
)

func init() {
	cmdr.AddCommand(getProjects)
	cmdr.AddCommand(getProject)
	cmdr.AddCommand(addProject)
	cmdr.AddCommand(deleteProject)
}

var getProjects = cmd.NewCommand(

	"get-projects", "Projects", "Gets all projects",

	func(cmd *cmd.Command) { }, 

	func(cmd *cmd.Command) error {
		res, err := apiClient.Projects.GetAll()
		if err != nil {
			return err
		}

		outputResponse(res)

		return nil
	},

)

var getProject = cmd.NewCommand(

	"get-project", "Projects", "Gets a project by ID",

	func(cmd *cmd.Command) {
		cmd.AppendArg("id", "Project ID")
	}, 

	func(cmd *cmd.Command) error {
		id, err := strconv.ParseInt(cmd.Arg("id"), 10, 64)
		if err != nil {
			return err
		}

		res, err := apiClient.Projects.Get(id)
		if err != nil {
			return err
		}

		outputResponse(res)

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

		res, err := apiClient.Projects.Create(string(pBytes))
		if err != nil {
			return err
		}

		outputResponse(res)

		return nil
	},

)

var deleteProject = cmd.NewCommand(
	"delete-project", "Projects", "Deletes a project",

	func(cmd *cmd.Command) {
		cmd.AppendArg("id", "Project ID")
	},

	func(cmd *cmd.Command) error {
		id, err := strconv.ParseInt(cmd.Arg("id"), 10, 64)
		if err != nil {
			return err
		}

		res, err := apiClient.Projects.Delete(id)
		if err != nil {
			return err
		}

		outputResponse(res)

		return nil
	},
)

