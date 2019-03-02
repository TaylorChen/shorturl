package conf

import (
	"errors"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
)

var (
	Conf              config // holds the global app config.
	defaultConfigFile = "conf/dev.conf"
)

type config struct {
	// Redis
	Redis redis
}

type redis struct {
	Server string `toml:"server"`
	Pwd    string `toml:"pwd"`
}

func init() {
}

// initConfig initializes the app configuration by first setting defaults,
// then overriding settings from the app config file, then overriding
// It returns an error if any.
func InitConfig(configFile string) error {
	if configFile == "" {
		configFile = defaultConfigFile
	}

	// Set defaults.
	Conf = config{}

	if _, err := os.Stat(configFile); err != nil {
		return errors.New("config file err:" + err.Error())
	} else {
		configBytes, err := ioutil.ReadFile(configFile)
		if err != nil {
			return errors.New("config load err:" + err.Error())
		}
		_, err = toml.Decode(string(configBytes), &Conf)
		if err != nil {
			return errors.New("config decode err:" + err.Error())
		}
	}

	return nil
}
