package config

type Config struct {
	HTTP HTTP
}

func Load() *Config {
	return &Config{HTTP: loadHTTP()}
}
