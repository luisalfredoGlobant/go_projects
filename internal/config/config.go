package config

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/qiangxue/go-env"
	"github.com/globant/crud_project/pkg/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const (
	defaultServerPort         = 8080
)

// Config represents an application configuration.
type Config struct {
	// the server port. Defaults to 8080
	ServerPort int `yaml:"server_port" env:"SERVER_PORT"`
	// the data source name (DSN) for connecting to the database. required.
	DSN string `yaml:"dsn" env:"DSN,secret"`
}

func (c Config) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.DSN, validation.Required),
	)
}

func Load(file string, logger log.Logger) (*Config, error) {
	// default config
	c := Config {
		ServerPort: defaultServerPort,
	}

	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(bytes, &c); err != nil {
		return nil, err
	}

	// load from environment variables prefixed with "APP_"
	if err = env.New("APP_", logger.Infof).Load(&c); err != nil {
		return nil, err
	}

	// validation
	if err = c.Validate(); err != nil {
		return nil, err
	}

	return &c, err
}
