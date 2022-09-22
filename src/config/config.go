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

//var configuration c.Configurations

////Reading variable using the model
//fmt.Println("\n\nReading variables using the model...")
//fmt.Println("Database is\t", configuration.Database.DBName)
//fmt.Println("port is \t\t", configuration.Server.Port)
//fmt.Println("EXAMPLE_PATH is\t", configuration.EXAMPLE_PATH)
//fmt.Println("EXAMPLE_VAR is\t", configuration.EXAMPLE_var)
