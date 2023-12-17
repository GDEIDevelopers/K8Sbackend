package config

import (
	"encoding/json"
	"log"
	"os"
	"reflect"
	"strings"
)

type Config struct {
	HTTPServerListen string

	DatabaseAddr string
	DatabasePort string
	DatabaseUser string
	DatabasePass string
	DatabaseDB   string

	RedisTokenAddr string
	RedisTokenDB   int `skip:"true"`
}

func Read(filename string) *Config {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var c Config

	json.Unmarshal(b, &c)

	c.assert()

	return &c
}

// assert函数要求配置必须填写完整
func (c *Config) assert(skip ...string) {
	elem := reflect.Indirect(reflect.ValueOf(c))
	for i := 0; i < elem.NumField(); i++ {
		current := elem.Field(i)
		if current.IsZero() && elem.Type().Field(i).Tag.Get("skip") == "" &&
			strings.ToLower(elem.Type().Field(i).Name[0:1]) != elem.Type().Field(i).Name[0:1] {
			log.Fatal("配置缺失: ", elem.Type().Field(i).Name)
		}
	}
}
