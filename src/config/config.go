package config

//ServerConfiguration exported
type ServerConfigurations struct {
	Port int
}

//Configurations exported
type Configurations struct {
	Server       ServerConfigurations
	Database     DatabaseConfiguration
	EXAMPLE_PATH string
	EXAMPLE_var  string
}

// DatabaseConfiguration  exported
type DatabaseConfiguration struct {
	DBName     string
	DBUser     string
	DBPassword string
}
