package v1

import (
	models_v0 "gothub_release_dl/models/v0"
	"gothub_release_dl/utils"

	"github.com/pelletier/go-toml/v2"
)

const Version = 1

type Repo struct {
	Destination string `toml:"destination"`
	AssetRegex  string `toml:"asset_regex"`
	RenameTo    string `toml:"rename_to"`
	Installer   bool   `toml:"installer"`
	Repository  string `toml:"repository"`
}

func FromV0(packages map[string]models_v0.Package) (map[string]Repo, map[string]Installed, error) {
	var repos = make(map[string]Repo, len(packages))
	var installed = make(map[string]Installed, len(packages))

	for k, v := range packages {
		repos[k] = Repo{
			Destination: v.Destination,
			AssetRegex:  v.Filename,
			RenameTo:    v.Rename,
			Installer:   v.Installer,
			Repository:  v.Repository,
		}

		installed[k] = Installed{
			InstalledAt: v.Install.Local(),
			Repository:  v.Repository,
			AssetName:   "",
		}
	}

	return repos, installed, nil
}

func WriteToReposToml(repos map[string]Repo, filename string) error {
	root := map[string]any{
		"version": Version,
	}

	return utils.WriteToToml(root, repos, filename)
}

func parseRepo(m map[string]any) (Repo, error) {
	b, err := toml.Marshal(m)
	if err != nil {
		return Repo{}, err
	}

	var r Repo
	if err := toml.Unmarshal(b, &r); err != nil {
		return Repo{}, err
	}

	return r, nil
}

func ReadFromReposToml(filename string) (map[string]Repo, error) {
	version, repos, err := utils.ReadFromToml(filename, parseRepo)

	if version != Version {
		return nil, utils.ErrVersionMismatch
	}

	return repos, err
}
