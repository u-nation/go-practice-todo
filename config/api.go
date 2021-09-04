package config

type APIConfig struct {
	AppName  string    `required:"true" envconfig:"APP_NAME"`
	Profile  string    `required:"true" envconfig:"PROFILE"`
	MaxConns int       `required:"true" split_words:"true"`
	DB       *DBConfig `ignored:"true"`
	// Redis     *RedisConfig `ignored:"true"`
	// Storage   *AWSS3Config `ignored:"true"`
	// XRay      *XRayConfig  `ignored:"true"`
}

func (c *APIConfig) IsLocal() bool {
	return c.Profile == "local"
}

func (c *APIConfig) IsDev() bool {
	return c.Profile == "dev"
}

func (c *APIConfig) IsStg() bool {
	return c.Profile == "stg"
}

func (c *APIConfig) IsPrd() bool {
	return c.Profile == "prd"
}