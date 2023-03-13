package main

import (
	"fmt"

	"github.com/forbearing/jumpserver"
)

func Asset_AccountBackupPlanExecution(jms *jumpserver.Client) {
	val, err := jms.Asset().AccountBackupPlanExecution().List(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}
