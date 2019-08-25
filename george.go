package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

var cfg *Config

func AppStart(c *cli.Context) error {
	cfg = ReadConfig(c.String("config"))

	// clone and/or update the repo
	CloneRepo(cfg.Git.Repo, cfg.Dirs.Base)
	UpdateRepo(cfg.Dirs.Base)

	// setup our http handlers
	if cfg.Git.Webhook.Enable {
		http.HandleFunc("/webhook", GitWebHook)
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(cfg.Dirs.Static))))
	http.HandleFunc("/", Router)

	// and finally
	log.Printf("Serving at: http://%s:%d/", cfg.Interface, cfg.Port)
	return http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Interface, cfg.Port), nil)
}

func main() {
	app := cli.NewApp()
	app.Name = "GeORGe"
	app.Usage = "Render Org/MarkDown files as HTML, on the fly."
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "config", Value: "/app/george.yml", Usage: "Specify the config file to use."},
	}
	app.Action = AppStart
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
