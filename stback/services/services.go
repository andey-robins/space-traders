package services

import "log"

func check(err error) {
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}
