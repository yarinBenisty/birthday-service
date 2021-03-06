package util

import (
	"github.com/spf13/viper"
)

const (
	// ConfigMongoConnectionString IS ..
	ConfigMongoConnectionString = "mongo_host"
)

// Config is used to declare every variable in
// app.env file in order to use it in go files
type Config struct {
	MongoURL string `mapstructure:"MONGO_URL"`
}

// LoadConfig loads all variables from app.env file
func LoadConfig() (err error) {
	viper.SetDefault(ConfigMongoConnectionString, "mongodb://root:example@0.0.0.0:27017")
	viper.AutomaticEnv()

	return nil
}
