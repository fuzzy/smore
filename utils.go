package main

import (
	"log"
	"os"
)

func lcheck(e error) error {
	if e != nil {
		log.Println(e)
		return e
	}
	return nil
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func isFile(t string) bool {
	info, err := os.Stat(t)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
