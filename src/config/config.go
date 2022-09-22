package config

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		DBName     string `yaml:"DBName"`
		DBUser     string `yaml:"DBUser"`
		DBPassword string `yaml:"DBPassword"`
	} `yaml:"database"`

	EXAMPLE_PATH string `yaml:"EXAMPLE_PATH"`
	EXAMPLE_var  string `yaml:"EXAMPLE_VAR"`
}
