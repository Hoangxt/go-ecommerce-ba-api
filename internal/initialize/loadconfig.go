package initialize

import (
	"fmt"

	"github.com/hoangxt/go-ecommerce-ba-api/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	// Load config from file
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
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}
}
