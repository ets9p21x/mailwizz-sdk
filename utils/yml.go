package utils

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path/filepath"
)

// Config 配置结构
type Config struct {
	PublicKey  string `yaml:"publicKey"`
	PrivateKey string `yaml:"privateKey"`
	Host       string `yaml:"host"`
}

// GobalConfig 配置全局变量
var GobalConfig = Config{
	PublicKey: "d253918744bfxxx6c87f9460ce22d6522",
	PrivateKey: "a41eb8891bbxxxe6394e055d175b00021f",
	Host: "https://xxx.com/latest/api/index.php",
}

// ReadYML 读取配置文件
func ReadYML(pathname string) {

	configFileAbsPath, err := filepath.Abs(pathname)
	if err != nil {
		panic("File Path error: " + err.Error())
	}

	fileBytes, err := ioutil.ReadFile(configFileAbsPath)
	if err != nil {
		panic("File to bytes error: " + err.Error())
	}
	err = yaml.Unmarshal(fileBytes, &GobalConfig)
	if err != nil {
		panic("Yaml to struct error: " + err.Error())
	}

}
