package otp

type Config struct {
	baseUrl     string
	application string
	secret      string
	devMode     bool
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) WithBaseUrl(baseUrl string) *Config {
	c.baseUrl = baseUrl
	return c
}

func (c *Config) WithApplication(application string) *Config {
	c.application = application
	return c
}

func (c *Config) WithSecret(secret string) *Config {
	c.secret = secret
	return c
}

func (c *Config) WithDevMode(devMode bool) *Config {
	c.devMode = devMode
	return c
}
