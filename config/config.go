package config

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
)

type Cfg struct {
	Env string
	Db  struct {
		Url string
	}
	Jwt struct {
		SignKey string
	}
}

var Conf Cfg

func Parse(cfg string) (*Cfg, error) {
	conf := Cfg{}

	if data, err := ioutil.ReadFile(cfg); err != nil {
		log.Fatalf("Could not read config file=%s because of %s.", cfg, err)
	} else {
		if _, err := toml.Decode(string(data), &conf); err != nil {
			return nil, errors.New("Could not read TOML config")
		}
	}
	Conf = conf
	log.Println("Read config=", fmt.Sprintf("%+v", conf))
	return &conf, nil
}
