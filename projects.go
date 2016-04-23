package main

import (
	"github.com/mikec/marsupi-api/models"
)

func init() {
	cmdr.AddCommand(getProjects)
	cmdr.AddCommand(getProject)
	cmdr.AddCommand(addProject)
	cmdr.AddCommand(deleteProject)
}

var getProjects = NewGetAllCommand(
	CommandProperties{"get-projects", "Projects", "Gets all projects"},
	&clt.Projects,
)

var getProject = NewGetCommand(
	CommandProperties{"get-project", "Projects", "Gets a project by ID"},
	&clt.Projects,
)

var deleteProject = NewDeleteCommand(
	CommandProperties{"delete-project", "Projects", "Deletes a project"},
	&clt.Projects,
)

var addProject = NewAddCommand(
	CommandProperties{"add-project","Projects","Adds a new project"},
	&clt.Projects,
	models.Project{},
	AddCommandArg{"name", "Name", "Project name"},
	AddCommandArg{"username", "Username", "Project owner user name"},
	AddCommandArg{"service", "Service", "Service where the project lives [ github | bitbucket(unsupported) ]"},
)
