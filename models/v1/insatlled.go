package v1

import (
	"gothub_release_dl/utils"
	"time"

	"github.com/pelletier/go-toml/v2"
)

type Installed struct {
	InstalledAt time.Time `toml:"installed_at"`
	Repository  string    `toml:"repository"`
	AssetName   string    `toml:"asset_name"`
}

func WriteToInstalledToml(installed map[string]Installed, filename string) error {
	root := map[string]any{
		"version": Version,
	}

	return utils.WriteToToml(root, installed, filename)
}

func parseInstalled(m map[string]any) (Installed, error) {
	b, err := toml.Marshal(m)
	if err != nil {
		return Installed{}, err
	}

	var r Installed
	if err := toml.Unmarshal(b, &r); err != nil {
		return Installed{}, err
	}

	return r, nil
}

func ReadFromInstalledToml(filename string) (map[string]Installed, error) {
	version, installed, err := utils.ReadFromToml(filename, parseInstalled)

	if version != Version {
		return nil, utils.ErrVersionMismatch
	}

	return installed, err
}
