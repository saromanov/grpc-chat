package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"

	"github.com/saromanov/grpc-chat/app/service"
	"github.com/saromanov/grpc-chat/app/service/config"
	"gopkg.in/yaml.v2"
)

const defaultPort = 14200

var configFile = flag.String("config", "", "config file")

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

// makePort provides making of the port into string representation
func makePort(port int) string {
	return fmt.Sprintf(":%d", port)
}

func main() {
	flag.Parse()
	if *configFile == "" {
		panic("config file is not defined")
	}

	conf, err := parseConfig(*configFile)
	if err != nil {
		panic(err)
	}

	port := makePort(defaultPort)
	if conf.Port != 0 {
		port = makePort(conf.Port)
	}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic("unable to listen: %v")
	}
	defer lis.Close()

	_, err = service.New(nil, lis)
	if err != nil {
		panic(err)
	}
}
