package v0

type Auth struct {
	Token string `json:"token"`
}

type Config struct {
	Auth Auth `json:"auth"`
}
