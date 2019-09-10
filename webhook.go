package main

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"gopkg.in/src-d/go-git.v4"
)

type HookContext struct {
	Signature string
	Event     string
	Id        string
	Payload   []byte
}

func ParseHook(secret []byte, req *http.Request) (*HookContext, error) {
	hc := HookContext{}

	// if hc.Signature = req.Header.Get("x-gitea-signature"); len(hc.Signature) == 0 {
	// 	return nil, errors.New("No signature!")
	// }
	if hc.Event = req.Header.Get("x-github-event"); len(hc.Event) == 0 {
		return nil, errors.New("No event!")
	}
	if hc.Id = req.Header.Get("x-github-delivery"); len(hc.Id) == 0 {
		return nil, errors.New("No event Id!")
	}

	body, err := ioutil.ReadAll(req.Body)
	check(err)

	hc.Payload = body
	return &hc, nil
}

func GitWebHook(w http.ResponseWriter, r *http.Request) {
	// parse the hook, This code I got from git examples and have
	// adapted to support Gitea currently, will soon allow for more
	// providers to be supported.
	hc, err := ParseHook([]byte(cfg.Git.Webhook.Secret), r)

	// set our output header
	w.Header().Set("Content-type", "application/json")

	// gracefully handle any errors
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Failed processing hook! ('%s')", err)
		io.WriteString(w, "{}")
		return
	}

	// at the moment I only support gitea, this should change soon
	if strings.Contains(string(hc.Payload), "Gitea") {
		log.Println("Gitea it is!!!!")
	}
	if strings.Contains(string(hc.Payload), "Github") {
		log.Println("GITHUB!!!")
	}
	pload := GiteaPush{}
	json.Unmarshal(hc.Payload, &pload)

	// if our secret matches then we should go ahead and update our checkout
	if pload.Secret == cfg.Git.Webhook.Secret {
		log.Printf("%+v", pload)
		log.Println("---------------------------------")
		log.Printf("%+v", pload.Repository)
		log.Println("---------------------------------")
		log.Printf("Updating repo from: %s", pload.Repository.CloneURL)
		log.Printf("Updating repo branch: %s", pload.Repository.DefaultBranch)
		log.Printf("Updating repo to commit: %s", pload.After)

		// open the repository
		repo, err := git.PlainOpen(cfg.Dirs.Base)
		check(err)

		// get the working tree
		wdir, err := repo.Worktree()
		check(err)

		// and pull the updates
		err = wdir.Pull(&git.PullOptions{})
		lcheck(err)
		log.Println("Updating repo: SUCCESS")
	}

	// and return our successful status
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "{}")

	// and hand everything back
	return
}
