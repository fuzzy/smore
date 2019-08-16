package main

import "log"

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
