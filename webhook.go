package main

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/go-playground/webhooks.v5/github"
)

func GitWebHook(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%+v\n", r)
	switch r.Method {
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func StartWebHook(secret string) {
	log.Println("Enabling webhook endpoint.")
	hook, err := github.New(github.Options.Secret(secret))
	check(err)

	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		log.Println("webhook payload triggered")
		payload, err := hook.Parse(r, github.ReleaseEvent, github.PullRequestEvent)
		if err != nil {
			if err != github.ErrEventNotFound {
				log.Fatal(err)
			}
		}

		switch payload.(type) {

		case github.ReleasePayload:
			release := payload.(github.ReleasePayload)
			// Do whatever you want from here...
			fmt.Printf("%+v\n", release)
		default:
			release := payload.(github.ReleasePayload)
			fmt.Printf("%+v\n", release)
		}

	})
	http.ListenAndServe(":3000", nil)
}
