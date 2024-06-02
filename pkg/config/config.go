package config

type UrlCheckerConfig struct {
	Limit  int    `json:"rate_limit"`
	Format string `json:"format"`
}
