package util

import "github.com/spf13/viper"

// Configuration structure holding all configuration variables
// The values are read by the viper plugin from a config file
type Config struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	LargestPrime  int    `mapstructure:"LARGEST_PRIME"`
}

// LoadConfig reads configuration from a file
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
