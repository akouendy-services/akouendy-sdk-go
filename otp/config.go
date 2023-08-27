package otp

type Config struct {
	BaseUrl     string
	Application string
	Secret      string
	DevMode     bool
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) WithBaseUrl(baseUrl string) *Config {
	c.BaseUrl = baseUrl
	return c
}

func (c *Config) WithApplication(application string) *Config {
	c.Application = application
	return c
}

func (c *Config) WithSecret(secret string) *Config {
	c.Secret = secret
	return c
}

func (c *Config) WithDevMode(devMode bool) *Config {
	c.DevMode = devMode
	return c
}
