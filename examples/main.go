package main

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/forbearing/jumpserver"
)

var (
	////jmsServerURL = "http://oajumper.cowarobot.cn"
	////accessKey    = "9ae9c9d7-4baf-41c0-b215-3ace8d91d783"
	////secretKey    = "e97e87b2-e418-4018-8397-6c86eab1392c"

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

	Example_Asset(jms)
	Example_User(jms)
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
