package main

import (
	"fmt"

	"github.com/forbearing/jumpserver"
)

func Asset_FavoriteAsset(jms *jumpserver.Client) {
	assets, err := jms.Asset().FavoriteAsset().List(nil)
	handleObject(err, assets, "list favorite assets")
	fmt.Println(assets)
}
