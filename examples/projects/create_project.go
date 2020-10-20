package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

func CreateProjectExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		name           string = "project-name"
		lifecycleID    string = "lifecycle-id"
		projectGroupID string = "project-group-id"
	)

	apiURL, err := url.Parse(octopusURL)
	if err != nil {
		_ = fmt.Errorf("error parsing URL for Octopus API: %v", err)
		return
	}

	client, err := octopusdeploy.NewClient(nil, apiURL, apiKey, spaceID)
	if err != nil {
		_ = fmt.Errorf("error creating API client: %v", err)
		return
	}

	// NOTE: the lifecycle and project group (below) can be obtained through the
	// respective services in the client

	// create project
	project := octopusdeploy.NewProject(name, lifecycleID, projectGroupID)

	// update any additional project fields here...

	// create project through Add(); returns error if fails
	createdProject, err := client.Projects.Add(project)
	if err != nil {
		_ = fmt.Errorf("error creating project: %v", err)
		return
	}

	fmt.Printf("project created: (%s)\n", createdProject.GetID())
}
