package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
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

		project, _, err := api.GetProject(id)
		if err != nil {
			return err
		}

		b, err := json.Marshal(project)
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

		res, err := api.DeleteProject(id)
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

