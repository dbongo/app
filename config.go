package main

import "github.com/danryan/env"

// config references the env vars used to configure the app
type config struct {
	DBHost     string `env:"key=DB_HOST"`
	DBName     string `env:"key=DB_NAME"`
	DBUser     string `env:"key=DB_USER"`
	DBPassword string `env:"key=DB_PASS"`
	ServerPort string `env:"key=PORT"`
}

// createConfig returns a config obj to be used by the app
func createConfig() (*config, error) {
	config := &config{}
	if err := env.Process(config); err != nil {
		return config, err
	}
	return config, nil
}
