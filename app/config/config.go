package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	// CF config
	CF = &Configs{}
)

// Configs config
type Configs struct {
	App struct {
		Port int `mapstructure:"PORT"`
	} `mapstructure:"APP"`
	Swagger struct {
		Title       string `mapstructure:"TITLE"`
		Host        string `mapstructure:"HOST"`
		Description string `mapstructure:"DESCRIPTION"`
	} `mapstructure:"SWAGGER"`
}

// InitConfig init config
func InitConfig() error {
	v := viper.New()
	v.AddConfigPath("configs")
	v.SetConfigName("config")
	v.AutomaticEnv()
	v.SetConfigType("yml")

	if err := v.ReadInConfig(); err != nil {
		logrus.Error("read config file error:", err)
		return err
	}

	if err := v.Unmarshal(&CF); err != nil {
		logrus.Error("unmarshal config error:", err)
		return err
	}

	return nil
}
