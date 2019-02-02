package main

import (
	"fmt"
	"log"
	"os"

	"github.com/glyn/go-force/force"
	//"github.com/glyn/go-force/sobjects"
)

func main() {
	forceApi, err := force.Create(
		"v44.0",
		os.Getenv("KARIS_FORCE_CLIENT_ID"),
		os.Getenv("KARIS_FORCE_CLIENT_SECRET"),
		os.Getenv("KARIS_FORCE_USER_NAME"),
		os.Getenv("KARIS_FORCE_PASSWORD"),
		os.Getenv("KARIS_FORCE_SECURITY_TOKEN"),
		"sandbox",
		"https://eu12.salesforce.com/services/oauth2/token",
	)
	if err != nil {
		log.Fatal(err)
	}

	_ = forceApi

	fmt.Println("Hello Karis")
}
