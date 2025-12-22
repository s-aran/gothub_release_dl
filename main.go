package main

import (
	"fmt"
	"os"

	models_v0 "gothub_release_dl/models/v0"
	models_v1 "gothub_release_dl/models/v1"
)

func main() {
	package_v0, err := models_v0.LoadJson("package.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", fmt.Sprintf("%s", err))
	}

	// fmt.Println(package_v0)

	repos, installed, err := models_v1.FromV0(package_v0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", fmt.Sprintf("%s", err))
	}

	// fmt.Println(repos)
	// fmt.Println(installed)

	models_v1.WriteToReposToml(repos, "repos.toml")
	models_v1.WriteToInstalledToml(installed, "installed.toml")

	repos, err = models_v1.ReadFromReposToml(("repos.toml"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", fmt.Sprintf("%s", err))
	}

	fmt.Println(repos)

	installed, err = models_v1.ReadFromInstalledToml(("installed.toml"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", fmt.Sprintf("%s", err))
	}

	fmt.Println(installed)

	config, err := models_v1.ReadFromConfigToml("config.toml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", fmt.Sprintf("%s", err))
	}
	config = models_v1.LoadEnvironmentVariable(config)

	fmt.Println(config.Auth.GithubApi)

}
