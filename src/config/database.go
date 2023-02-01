package config

type Database struct {
	DB struct {
		ConnectionString string `yaml:"connection_string"`
		Name             string `yaml:"name"`
	}
}
