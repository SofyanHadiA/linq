package core

import (
	"io/ioutil"
	"os"
	"strings"

	"bitbucket.org/sofyan_a/linq.im/core/utils"

	"gopkg.in/yaml.v2"
)

const ENVAR_CONFIG_PREFIX = "LINQ_"

type Configs map[string]interface{}

var configs Configs

func init() {
	configs = loadConfig("conf/db.conf")
	appConfig := loadConfig("conf/app.conf")
	auth0Config := loadConfig("conf/auth0.conf")

	utils.MapCopy(configs, appConfig)
	utils.MapCopy(configs, auth0Config)

	utils.Log.Info("Application configs: ", configs)
}

func GetStrConfig(configKey string) string {
	return configs[configKey].(string)
}

func GetIntConfig(configKey string) int {
	return configs[configKey].(int)
}

func loadConfig(file string) Configs {

	var config Configs

	conf, errFile := ioutil.ReadFile(file)
	if errFile != nil {
		utils.Log.Fatal("error: %v", errFile)
	}

	errYaml := yaml.Unmarshal([]byte(conf), &config)
	if errYaml != nil {
		utils.Log.Fatal("error: %v", errYaml)
	}

	for k := range config {
		envarKey := ENVAR_CONFIG_PREFIX + strings.ToUpper(strings.Replace(k, ".", "_", -1))
		envarValue := os.Getenv(envarKey)
		if len(envarValue) > 0 {
			config[k] = envarValue
		}
	}

	return config
}
