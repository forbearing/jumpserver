package main

import (
	"fmt"

	"github.com/forbearing/jumpserver"
)

func Asset_Account(jms *jumpserver.Client) {
	//accounts, err := jms.Asset().Account().Create(
	//    &jumpserver.Account{Username: "hybfkuf"},
	//)
	//handleObject(err, accounts, "create asset accounts")

	accounts, err := jms.Asset().Account().List(nil)
	handleObject(err, accounts, "list all asset accounts")
	fmt.Println(accounts)
}
