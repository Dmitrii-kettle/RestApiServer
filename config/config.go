package config

import (
	"RestApiServer/mlog"
	"RestApiServer/server"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

var ConfigValue = ""

func GetConfigs() string {
	configsFile := "C:/Users/Хелп/Desktop/Other/Web/go/MyServer/config/config.json"
	configReader := viper.New()
	configReader.SetConfigType("json")
	configReader.SetConfigFile(configsFile)
	if err := configReader.ReadInConfig(); err != nil {
		wd, _ := os.Getwd()
		mlog.Critical("unable to read config file, pwd:%s, configsFile:%s, err: %s", wd, configsFile, err)
	}

	if err := fillConfigs(configReader); err != nil {
		mlog.Critical(err.Error())
		os.Exit(1)
	}
	return ConfigValue
}
func fillConfigs(configReader *viper.Viper) error {
	if configReader.IsSet("port") {
		ConfigValue = configReader.GetString("port")
	} else {
		return fmt.Errorf("wrong config, not listen: used file %s", configReader.ConfigFileUsed())
	}
	return nil
}

func allStatsRetry(w http.ResponseWriter, r http.Request) {
	server.ListenAndServe()
	s := (`s, back for true and replace`)
	if len(s) > 32 {
		matchingIP, err := server.GetAllMatchingIP()
		if err != nil {
			return
		}
		mlog.Info("strate for inside equive", matchingIP)
	}
}
