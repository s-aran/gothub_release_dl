package v1

type Auth struct {
	GithubApi string `toml:"github_api"`
}

type Config struct {
	auth Auth
}
