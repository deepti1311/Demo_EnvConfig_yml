package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {

	//set the file name of the configurations file
	viper.SetConfigName("config")

	//set the path to look for the configurations file
	viper.AddConfigPath(".")

	//Enable VIPER to read Environment variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error of Configuration file, %s\n", err)
	}

	//set undefined variables
	viper.SetDefault("database.name", "test_db")

	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}

	//Reading variable without using the model
	fmt.Println("\n\nReading variables without using the model...")
	fmt.Println("Database is\t", viper.GetString("database.name"))
	fmt.Println("port is \t\t", viper.Get("Port"))
	fmt.Println("EXAMPLE_PATH is\t", viper.GetString("EXAMPLE_PATH"))
	fmt.Println("EXAMPLE_VAR is\t", viper.GetString("EXAMPLE_var"))

}
