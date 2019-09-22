package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"gopkg.in/src-d/go-git.v4"
)

func safeError(e error) bool {
	safe_messages := []string{
		"repository already exists",
		"already up-to-date",
	}
	for _, v := range safe_messages {
		if e.Error() == v {
			return true
		}
	}
	return false
}

func cleanupAttribution(s string) string {
	start := s[1:][:len(s)-2]
	re := regexp.MustCompile("([a-zA-Z\\.\\-\\_0-9]+)@([a-zA-Z0-9\\.\\-\\_]+)")
	next := re.ReplaceAllString(start, "<${1} _at_ ${2}>>")
	bits := strings.Split(next, "> ")
	date := strings.Split(bits[1], " ")
	retv := fmt.Sprintf("%s at %s %s", bits[0], date[0], date[1])
	return retv
}

func getFileAuthor(f string) string {
	retv := ""
	nf := strings.Split(f, cfg.Dirs.Base)[1][1:]
	repo, err := git.PlainOpen(cfg.Dirs.Base)
	if lcheck(err) != nil && !safeError(err) {
		log.Fatal(err)
	}
	cIter, err := repo.Log(&git.LogOptions{FileName: &nf})
	c, err := cIter.Next()
	for err == nil {
		retv = fmt.Sprint(c.Author)
		c, err = cIter.Next()
	}
	return cleanupAttribution(retv)
}

func getFileLastChanged(f string) string {
	retv := ""
	nf := strings.Split(f, cfg.Dirs.Base)[1][1:]
	repo, err := git.PlainOpen(cfg.Dirs.Base)
	if lcheck(err) != nil && !safeError(err) {
		log.Fatal(err)
	}
	cIter, err := repo.Log(&git.LogOptions{FileName: &nf})
	c, err := cIter.Next()
	if err != nil {
		log.Println(err)
	}
	retv = fmt.Sprint(c.Committer)
	return cleanupAttribution(retv)
}

func CloneRepo(r, b string) error {
	_, err := git.PlainClone(b, false, &git.CloneOptions{URL: r})
	if lcheck(err) != nil && !safeError(err) {
		return err
	}
	return nil
}

func UpdateRepo(r string, i int64) {
	for {
		repo, err := git.PlainOpen(r)
		if lcheck(err) != nil && !safeError(err) {
			log.Fatal(err)
		}
		wdir, err := repo.Worktree()
		if lcheck(err) != nil && !safeError(err) {
			log.Fatal(err)
		}
		err = wdir.Pull(&git.PullOptions{})
		if lcheck(err) != nil && !safeError(err) {
			log.Fatal(err)
		}
		time.Sleep(time.Duration(i) * time.Second)
	}
}
