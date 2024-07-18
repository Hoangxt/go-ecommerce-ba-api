package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config/") // path to look for the config file in
	viper.SetConfigName("local")     // name of config file (without extension)
	viper.SetConfigType("yaml")      // REQUIRED if the config file does not have the extension in the name

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// read server configuration
	fmt.Println("Server Port:: ", viper.GetInt("server.port"))
	fmt.Println("Security Key:: ", viper.GetString("security.jwt.key"))

	// config struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}

	fmt.Println("Config Port:: ", config.Server.Port)

	for _, db := range config.Databases {
		fmt.Printf("Database User: %s, Password: %s, Host: %s \n", db.User, db.Password, db.Host)
	}
}
