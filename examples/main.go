package main

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/forbearing/jumpserver"
)

var (
	jmsServerURL = "http://10.250.100.40"
	accessKey    = "cfa5d3d1-8fe1-496b-92c0-ff8e5024cff5"
	secretKey    = "374c71eb-9a2c-4ec1-a35c-ec78a9c4a300"
)

func main() {
	jms, err := jumpserver.New(jmsServerURL,
		jumpserver.AuthWithAccessKey(accessKey, secretKey),
		jumpserver.WithHTTPClient(&http.Client{Timeout: time.Second * 3}),
	)
	if err != nil {
		panic(err)
	}

	Asset_Asset(jms)
	//Asset_AccountBackupPlanExecution(jms)
	//Asset_AccountBackupPlan(jms)
	//Asset_AccountHistory(jms)
	//Asset_Account(jms)
	//Asset_CmdFilterRule(jms)
	//Asset_CmdFilter(jms)
	//Asset_Domain(jms)
}

func handleObject(err error, obj any, action string) {
	if err != nil {
		panic(err)
	}
	switch obj.(type) {
	case nil:
		fmt.Printf("Successfully %s \n", action)
	case jumpserver.Object:
		fmt.Printf("Successfully %s \n", action)
	default:
		fmt.Printf("Successfully %s (total %d)\n", action, reflect.ValueOf(obj).Len())
	}
}
