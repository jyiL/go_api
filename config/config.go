package config

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
)

type RpcConf struct {
    Url string   `yaml:"url"`
}
var RpcYamlConfig RpcConf

func init() {
    RpcYamlConfig.getConf()
}

func (c *RpcConf) getConf() *RpcConf {

    yamlFile, err := ioutil.ReadFile("config/rpc.yaml")
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }

    return c
}