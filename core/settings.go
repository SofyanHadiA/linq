package core

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"linq/core/utils"

	"gopkg.in/yaml.v2"
)

const ENVAR_CONFIG_PREFIX = "LINQ_"

type Configs map[string]interface{}

var configs Configs

func init() {
	configs = loadConfig("db.conf")
	appConfig := loadConfig("app.conf")

	utils.MapCopy(configs, appConfig)
}

func GetStrConfig(configKey string) string {
	return configs[configKey].(string)
}

func GetIntConfig(configKey string) string {
	return strconv.Itoa(configs[configKey].(int))
}

func loadConfig(file string) Configs {

	var config Configs

	conf, errFile := ioutil.ReadFile(file)
	if errFile != nil {
		log.Fatalf("error: %v", errFile)
	}

	errYaml := yaml.Unmarshal([]byte(conf), &config)
	if errYaml != nil {
		log.Fatalf("error: %v", errYaml)
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
