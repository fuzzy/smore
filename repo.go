package main

import (
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

func CloneRepo(r, b string) error {
	_, err := git.PlainClone(b, false, &git.CloneOptions{URL: r})
	if lcheck(err) != nil && !safeError(err) {
		return err
	}
	return nil
}

func UpdateRepo(r string) error {
	repo, err := git.PlainOpen(r)
	if lcheck(err) != nil && !safeError(err) {
		return err
	}
	wdir, err := repo.Worktree()
	if lcheck(err) != nil && !safeError(err) {
		return err
	}
	err = wdir.Pull(&git.PullOptions{})
	if lcheck(err) != nil && !safeError(err) {
		return err
	}
	return nil
}
