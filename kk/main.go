package main

import (
	"fmt"
	"github.com/glyn/go-force/sobjects"
	"log"
	"os"

	"github.com/glyn/go-force/force"
)

type KarisAccount struct {
	sobjects.Account
	ParentId string `force:"ParentId"`
	LinkedFamilyId string `force:"kk_Linked_Family__c"` // reference to Account
	CcmId          string `force:"kk_Church_Community_Mobiliser_CCM__c"` // reference to Contact
}

type AccountQueryResponse struct {
	sobjects.BaseQuery
	Records []*sobjects.Account `force:"records"`
}


type Contact struct {
	sobjects.BaseSObject
}

func (c Contact) ApiName() string {
	return "Contact"
}

type ContactQueryResponse struct {
	sobjects.BaseQuery
	Records []*Contact `force:"records"`
}

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

	acQueryyResp := &AccountQueryResponse{}
	//err = forceApi.Query("SELECT Id FROM Account LIMIT 20", acQueryyResp)
	err = forceApi.Query("SELECT Id FROM Account", acQueryyResp)
	if err != nil {
		fmt.Println(err)
	}

	for _, c := range acQueryyResp.Records {
		var ac KarisAccount
		err := forceApi.GetSObject(c.Id, []string{"Name", "ParentId", "kk_Linked_Family__c"}, &ac)
		if err != nil {
			fmt.Println(err)
		}
		if ac.LinkedFamilyId != "" {
			var linkedFamily KarisAccount
			err := forceApi.GetSObject(ac.LinkedFamilyId, []string{"Name", "ParentId"}, &linkedFamily)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("%s (%s) linked to %s (%s)\n", ac.Name, parentId(forceApi, ac), linkedFamily.Name, parentId(forceApi, linkedFamily))
		}
	}

	contactQueryResp := &ContactQueryResponse{}
	//err = forceApi.Query("SELECT Id FROM Contact LIMIT 20", contactQueryResp)
	err = forceApi.Query("SELECT Id FROM Contact", contactQueryResp)
	if err != nil {
		fmt.Println(err)
	}

	for _, c := range contactQueryResp.Records {
		var contact Contact
		err := forceApi.GetSObject(c.Id, []string{}, &contact)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Contacct %s\n", contact.Name)
	}
}

func parentId(forceApi *force.ForceApi, ac KarisAccount) string {
	acChurch := ""
	if ac.ParentId != "" {
		var parent KarisAccount
		err := forceApi.GetSObject(ac.ParentId, []string{"Name"}, &parent)
		if err != nil {
			fmt.Println(err)
		}
		acChurch = parent.Name
	}
	return acChurch
}
