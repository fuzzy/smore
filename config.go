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
	Port      int64  `yaml:"port"`
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
		Interval int64  `yaml:"interval"`
		Webhook  struct {
			Enable bool   `yaml:"webhook"`
			Secret string `yaml:"secret"`
		} `yaml:"webhook"`
	} `yaml:"git"`
}

func envConfigUpdate(c *Config) {
	vars := map[string]string{
		"SMORE_INTERFACE":      c.Interface,
		"SMORE_PORT":           string(c.Port),
		"SMORE_TEMPLATE":       c.Template,
		"SMORE_DIRS_ROOT":      c.Dirs.Root,
		"SMORE_DIRS_BASE":      c.Dirs.Base,
		"SMORE_DIRS_STATIC":    c.Dirs.Static,
		"SMORE_DIRS_TEMPLATES": c.Dirs.Templates,
		"SMORE_DIRS_CONTENT":   c.Dirs.Content,
		"SMORE_GIT_REPO":       c.Git.Repo,
		"SMORE_GIT_INTERVAL":   "", // emp disable the update interval, set to a number otherwise
		"SMORE_GIT_WEBHOOK":    "", // empty disables the webhook, contains secret otherwise
	}
	for k, v := range vars {
		ev, ef := os.LookupEnv(k)
		if ef {
			switch k {
			case "SMORE_INTERFACE":
				c.Interface = ev
			case "SMORE_PORT":
				c.Port = int(ev)
			case "SMORE_TEMPLATE":
				c.Template = ev
			case "SMORE_DIRS_ROOT":
				c.Dirs.Root = ev
			case "SMORE_DIRS_BASE":
				c.Dirs.Base = ev
			case "SMORE_DIRS_STATIC":
				c.Dirs.Static = ev
			case "SMORE_DIRS_TEMPLATES":
				c.Dirs.Templates = ev
			case "SMORE_DIRS_CONTENT":
				c.Dirs.Content = ev
			case "SMORE_GIT_REPO":
				c.Git.Repo = ev
			case "SMORE_GIT_INTERVAL":
				c.Git.Interval = int64(ev)
			case "SMORE_GIT_WEBHOOK":
				c.Git.Webhook.Enable = true
				c.Git.Webhook.Secret = ev
			}
		} else {
			switch k {
			case "SMORE_GIT_INTERVAL":
				c.Git.Interval = 0
			case "SMORE_GIT_WEBHOOK":
				c.Git.Webhook.Enable = false
				c.Git.Webhook.Secret = ""
			}
		}
	}
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

	// ok, now we're done with all that
	log.Printf("Read config: %s", fn)

	// let's see if the environment has anything to say about it
	envConfigUpdate(&retv)

	return retv
}
