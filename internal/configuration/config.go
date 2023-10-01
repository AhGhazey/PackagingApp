package configuration

import (
	"path/filepath"
	"sync"
	"time"

	"github.com/spf13/viper"
)

var (
	configuration *AppConfiguration
	once          sync.Once
)

type AppConfiguration struct {
	// Add server configuration
	ServiceName    string        `mapstructure:"SERVICE_NAME"`
	ServerAddress  string        `mapstructure:"SERVER_ADDRESS"`
	ReadTimeout    time.Duration `mapstructure:"READ_TIMEOUT"`
	WriteTimeout   time.Duration `mapstructure:"WRITE_TIMEOUT"`
	IdleTimeout    time.Duration `mapstructure:"IDLE_TIMEOUT"`
	WaitingTimeout time.Duration `mapstructure:"WAITING_TIMEOUT"`
	LogLevel       string        `mapstructure:"LOG_LEVEL"`
	Environment    string        `mapstructure:"ENVIRONMENT"`
}

func loadConfig() (config AppConfiguration, err error) {
	configPath, _ := filepath.Abs(".")
	viper.AddConfigPath(configPath)
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")

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

func GetConfiguration() (*AppConfiguration, error) {
	var err error
	once.Do(func() {
		config, err := loadConfig()
		if err != nil {
			//log error
			return
		}

		configuration = &config
	})
	return configuration, err
}
