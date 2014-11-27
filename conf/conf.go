package conf

import "github.com/danryan/env"

// Conf ...
type Conf struct {
	DBHost string `env:"key=MONGO_ADDRESS"`
	DBName string `env:"key=MONGO_DATABASE"`
	Port   string `env:"key=HOST_PORT"`
}

// New returns a conf obj to be used by the app
func New() (*Conf, error) {
	conf := &Conf{}
	if err := env.Process(conf); err != nil {
		return conf, err
	}
	return conf, nil
}
