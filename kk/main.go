package main

import (
	"fmt"
	"github.com/glyn/go-force/sobjects"
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

	type AccountQueryResponse struct {
		sobjects.BaseQuery
		Records []*sobjects.Account `force:"records"`
	}

	someCustomSObjects := &AccountQueryResponse{}
	err = forceApi.Query("SELECT Id FROM Account LIMIT 20", someCustomSObjects)
	if err != nil {
		fmt.Println(err)
	}

	for _, c := range someCustomSObjects.Records {
		var out sobjects.Account
		err := forceApi.GetSObject(c.Id, []string{}, &out)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%#v\n", out)

	}

	fmt.Println("Hello Karis")
}
