package main

import (
	"fmt"

	"github.com/forbearing/jumpserver"
)

func Asset_CmdFilterRule(jms *jumpserver.Client) {
	rules, err := jms.Asset().CmdFilterRule().List(nil)
	handleObject(err, rules, "list all asset cmd filter rules")
	fmt.Println(rules)
}
