package main

import (
	"fmt"
	"jumpserver"
)

var (
	jmsServerURL = "http://oajumper.cowarobot.cn"
	accessKey    = "9ae9c9d7-4baf-41c0-b215-3ace8d91d783"
	secretKey    = "e97e87b2-e418-4018-8397-6c86eab1392c"
)

func main() {
	jms, err := jumpserver.New(jmsServerURL, accessKey, secretKey)
	handleErr(err)

	//fmt.Println(string(jms.MustGetUsers()))
	//fmt.Println(string(jms.MustGetAssets()))
	fmt.Println(string(jms.MustGetNodes()))

}

// handleErr
func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
