package config

import (
	"log"

	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

// Configuration - configuration model
type Configuration struct {
	Port   int    `json:"port" mapstructure:"PORT"`
	DBURL  string `json:"dbURL" mapstructure:"DB_URL"`
	DBName string `json:"dbName" mapstructure:"DB_NAME"`
	DBUser string `json:"dbUser" mapstructure:"DB_USER"`
	DBPass string `json:"dbPass" mapstructure:"DB_PASSWORD"`
}

type configuration struct {
	config Configuration
	logger *logrus.Logger
}

// LoadConfig - load project configuration
func LoadConfig(path string, logger *logrus.Logger) (*configuration, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	viper.SetDefault("PORT", 8080)

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Unable to read the config file: %s", err)
		return nil, err
	}
	config := Configuration{}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	logger.Info("[config] LoadConfig - readConfigFile")
	logger.Debugf("... context : %v", config)

	return &configuration{
		config: config,
		logger: logger,
	}, nil
}

func (c *configuration) GetConfig() *Configuration {
	return &c.config
}

func (c *configuration) ValidateConfig() {
	if len(c.config.DBName) == 0 {
		log.Fatal("Missing required 'DBName' parameter")
	}

	if len(c.config.DBURL) == 0 {
		log.Fatal("Missing required 'DBURL' parameter")
	}

	if len(c.config.DBName) == 0 {
		log.Fatal("Missing required 'DBName' parameter")
	}
}
