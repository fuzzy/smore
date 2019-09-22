package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/fsnotify/fsnotify"
	"github.com/urfave/cli"
)

var cfg *Config

func AppStart(c *cli.Context) error {
	// creates a new file watcher
	watcher, err := fsnotify.NewWatcher()
	check(err)
	defer watcher.Close()

	// read in our config initally
	cfg = ReadConfig(c.String("config"))

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op.String() == "WRITE" {
					cfg = ReadConfig(c.String("config"))
					log.Println("Rereading %s", c.String("config"))
				}
				// watch for errors
			case err := <-watcher.Errors:
				log.Println("ERROR", err)
			}
		}
	}()

	// out of the box fsnotify can watch a single file, or a single directory
	if err := watcher.Add(c.String("config")); err != nil {
		log.Fatal(err)
	}

	if !isDir(fmt.Sprintf("%s/%s", cfg.Dirs.Base, path.Base(cfg.Git.Repo))) {
		// clone and/or update the repo
		CloneRepo(cfg.Git.Repo, cfg.Dirs.Base)
	}

	// if we have a non 0 interval we should start the update thread
	if cfg.Git.Interval >= 1 {
		go UpdateRepo(cfg.Dirs.Base, cfg.Git.Interval)
	}

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
	app.Name = "SMORE"
	app.Usage = "Render Org/MarkDown files as HTML, on the fly."
	app.Version = "0.4.15"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "config", Value: "/app/smore.yml", Usage: "Specify the config file to use."},
	}
	app.Action = AppStart
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
