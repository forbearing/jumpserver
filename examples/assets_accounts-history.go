package main

import "github.com/forbearing/jumpserver"

func Asset_AccountHistory(jms *jumpserver.Client) {
	historys, err := jms.Asset().AccountHistory().List(nil)
	handleObject(err, historys, "list all account historys")
}
