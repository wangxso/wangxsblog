package utils

import (
	"io"
	"os"
	"runtime"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`

	Db struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"db"`
}

func GetConfig() (*Config, error) {
	config := Config{}
	// 访问配置文件
	var path string
	if runtime.GOOS == "windows" {
		path = "C:\\Users\\admin\\wangxsblog\\config\\config.yaml"
	} else if runtime.GOOS == "darwin" {
		path = "/Users/wangxs/wangxsoblog/config/config.yaml"
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
