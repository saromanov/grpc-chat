package main

import (
	"fmt"
	"io/ioutil"
	"net"

	"github.com/saromanov/grpc-chat/app/service"
	"github.com/saromanov/grpc-chat/app/service/config"
	"gopkg.in/yaml.v2"
)

// parseConfig provides parsing of the config .yml file
func parseConfig(path string) (*config.Config, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open config file: %v", err)
	}
	var c *config.Config
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return nil, fmt.Errorf("unable to parse .config.yml: %v", err)
	}

	return c, nil
}

func main() {
	lis, err := net.Listen("tcp", ":14200")
	if err != nil {
		panic("unable to listen: %v")
	}
	defer lis.Close()

	_, err = service.New(nil, lis)
	if err != nil {
		panic(err)
	}
}
