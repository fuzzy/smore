package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"gopkg.in/src-d/go-git.v4"
)

func signBody(secret, body []byte) []byte {
	computed := hmac.New(sha1.New, secret)
	computed.Write(body)
	return []byte(computed.Sum(nil))
}

func verifySignature(secret []byte, signature string, body []byte) bool {

	const signaturePrefix = "sha1="
	const signatureLength = 45 // len(SignaturePrefix) + len(hex(sha1))

	if len(signature) != signatureLength || !strings.HasPrefix(signature, signaturePrefix) {
		return false
	}

	actual := make([]byte, 20)
	hex.Decode(actual, []byte(signature[5:]))

	return hmac.Equal(signBody([]byte(cfg.Git.Webhook.Secret), body), actual)
}

type HookContext struct {
	Signature string
	Event     string
	Id        string
	Payload   []byte
}

func ParseHook(secret []byte, req *http.Request) (*HookContext, error) {
	hc := HookContext{}

	if hc.Signature = req.Header.Get("x-gitea-signature"); len(hc.Signature) == 0 {
		return nil, errors.New("No signature!")
	}

	if hc.Event = req.Header.Get("x-github-event"); len(hc.Event) == 0 {
		return nil, errors.New("No event!")
	}

	if hc.Id = req.Header.Get("x-github-delivery"); len(hc.Id) == 0 {
		return nil, errors.New("No event Id!")
	}

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		return nil, err
	}

	// if !verifySignature(secret, hc.Signature, body) {
	// 	return nil, errors.New("Invalid signature")
	// }

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
	pload := GiteaPush{}
	json.Unmarshal(hc.Payload, &pload)

	// if our secret matches then we should go ahead and update our checkout
	if pload.Secret == cfg.Git.Webhook.Secret {
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
		check(err)
		log.Println("Updating repo: SUCCESS")
	}

	// and return our successful status
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "{}")

	// and hand everything back
	return
}
