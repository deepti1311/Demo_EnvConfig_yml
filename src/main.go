package main

import (
	c "Demo_EnvConfig_yml/src/config"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

func main() {

	f, err := os.Open("application.yml")
	if err != nil {
		fmt.Printf("Error of Configuration file, %s\n", err)
	}

	defer f.Close()

	var config c.Config

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}

	//Reading variable using the model
	fmt.Println("\n\nReading variables using the model...")
	fmt.Println("Database is\t", config.Database.DBName)
	fmt.Println("port is \t\t", config.Server.Port)
	fmt.Println("EXAMPLE_PATH is\t", config.EXAMPLE_PATH)
	fmt.Println("EXAMPLE_VAR is\t", config.EXAMPLE_var)

}
