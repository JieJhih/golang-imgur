package config

import "io/ioutil"
import "gopkg.in/yaml.v2"

type ConfYaml struct {
	Auth   SectionAuth   `yaml:"auth"`
	Server SectionServer `yaml:"server"`
	API    SectionAPI    `yaml:"api"`
}

type SectionAuth struct {
	ClientID string `yaml:"client_id"`
}

type SectionServer struct {
	Port         string `yaml:"port"`
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
}

type SectionAPI struct {
	UploadImage string `yaml:"upload_image"`
}

// BuildDefaultPushConf is default config setting.
func BuildDefaultPushConf() ConfYaml {
	var conf ConfYaml

	conf.Auth.ClientID = ""
	conf.Server.Port = "8081"
	conf.Server.ReadTimeout = 10
	conf.Server.WriteTimeout = 10

	return conf
}

func LoadConfYaml(confPath string) (ConfYaml, error) {
	var config ConfYaml

	configFile, err := ioutil.ReadFile(confPath)

	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(configFile, &config)

	if err != nil {
		return config, err
	}

	return config, nil
}
