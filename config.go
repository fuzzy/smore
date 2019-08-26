package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Interface string `yaml:"interface"`
	Port      int    `yaml:"port"`
	Template  string `yaml:"template"`
	Dirs      struct {
		Root      string
		Base      string `yaml:"base"`
		Static    string `yaml:"static"`
		Templates string `yaml:"templates"`
		Content   string `yaml:"content"`
	} `yaml:"dirs"`
	Git struct {
		Repo     string `yaml:"repo"`
		Interval int64    `yaml:"interval"`
		Webhook  struct {
			Enable bool   `yaml:"webhook"`
			Secret string `yaml:"secret"`
		} `yaml:"webhook"`
	} `yaml:"git"`
}

func ReadConfig(fn string) *Config {
	retv := &Config{}
	if _, err := os.Stat(fn); err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal([]byte(data), retv)
	if err != nil {
		log.Fatal(err)
	}

	if retv.Git.Repo != "" {
		data := strings.Split(retv.Git.Repo, "/")
		retv.Dirs.Root = retv.Dirs.Base
		tdir := fmt.Sprintf("%s/%s", retv.Dirs.Base, data[len(data)-1])
		retv.Dirs.Base = tdir
		tdir = fmt.Sprintf("%s/%s", retv.Dirs.Base, retv.Dirs.Content)
		retv.Dirs.Content = tdir
		tdir = fmt.Sprintf("%s/%s", retv.Dirs.Base, retv.Dirs.Static)
		retv.Dirs.Static = tdir
		tdir = fmt.Sprintf("%s/%s", retv.Dirs.Base, retv.Dirs.Templates)
		retv.Dirs.Templates = tdir
	} else {
		log.Fatal("No git repo configured")
	}

	log.Printf("Read config: %s", fn)
	return retv
}
