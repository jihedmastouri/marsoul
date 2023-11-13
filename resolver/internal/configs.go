package internal

import (
	"github.com/spf13/viper"
)

type Configs struct {
	Addr      string
	Port      int
	CertLoc   string
	PrvKeyLoc string
}

func initConfigs() {
	viper.SetDefault("addr", "locahost")
	viper.SetDefault("port", "4220")
	viper.SetDefault("certLoc", "server-cert.pem")
	viper.SetDefault("prvKeyLoc", "server-key.pem")
}

func NewConfigs() (*Configs, error) {
	viper.SetConfigName("resolver.config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.marsoul")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			initConfigs()
			viper.WriteConfig()
		} else {
			return nil, err
		}
	}

	return &Configs{
		Addr:      viper.GetString("addr"),
		Port:      viper.GetInt("port"),
		CertLoc:   viper.GetString("certLoc"),
		PrvKeyLoc: viper.GetString("prvKeyLoc"),
	}, nil
}
