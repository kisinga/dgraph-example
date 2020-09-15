package config

import (
	"fmt"
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

type DatabaseConfig struct {
	URI string `yaml:"uri"`
}

//Config This maps the configuration in the yaml file into a struct
type Config struct {
	DBConfig DatabaseConfig `yaml:"database"`
}

//ReadYaml reads a yml file and returns the mapped config
func ReadYaml(path string) (*Config, error) {
	if path == "" {
		path = defaultYamlConfigPath()
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		fmt.Printf("error reading yaml file into config struct: %s\n", err)
		os.Exit(2)
	}
	return &cfg, nil
}

func defaultYamlConfigPath() string {
	// Reads the path of the current executable
	// goes up 2 directories and appends config.yaml
	// to the path.
	ex, err := os.Executable()
	if err != nil {
		log.Printf("error encountered reading path: %s\n", err)
		os.Exit(2)
	}

	filename := "config.yml"
	dir := path.Dir(path.Dir(ex))
	dir = path.Join(dir, filename)
	return dir
}
