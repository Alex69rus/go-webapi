package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
)

func readConfiguration(key string, cfg interface{}) {
	file, err := os.Open("config.json")
	processError(err)
	defer file.Close()

	fullCfg := map[string]interface{}{}
	err = json.NewDecoder(file).Decode(&fullCfg)
	processError(err)

	cfgByKey, ok := fullCfg[key]
	if !ok {
		processError(fmt.Errorf("Coud not found key: %v", key))
	}

	res, err := json.Marshal(cfgByKey)
	processError(err)

	json.Unmarshal(res, cfg)

	readEnv(cfg)
}

func readEnv(cfg interface{}) {
	err := envconfig.Process("", cfg)
	processError(err)
}

func processError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
