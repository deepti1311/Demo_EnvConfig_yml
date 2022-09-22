# GOLANG Environment Configuration file



How to read variables from the list of Environment Variable and Configuration Files in GO using VIPER

This branch contains the implementation of reading from Configuration Files and Environment Variables using the VIPER package from https://github.com/spf13/viper



### TO read the my application.yml file of Environment variable configuration 

* To use data from a .yml file in Go you need to unmarshal it into the structure like you do for JSON.

* I use gopkg.in/yaml.v2 library by Cannonical to parse YAML files.