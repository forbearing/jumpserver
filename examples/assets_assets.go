package main

import (
	"fmt"

	"github.com/forbearing/jumpserver"
)

func Asset_Asset(jms *jumpserver.Client) {
	var (
		asset    *jumpserver.Asset
		assets   []*jumpserver.Asset
		hostname = "myserver"
		ip       = "1.1.1.1"
		err      error
	)
	var _, _ = asset, assets
	jms.Asset().Asset().Delete(&jumpserver.DeleteAssetParam{Hostname: hostname})

	// 1. Create jumpserver asset.
	assets, err = jms.Asset().Asset().Create(&jumpserver.Asset{
		Hostname:     hostname,
		IP:           ip,
		Platform:     jumpserver.Linux,
		Protocols:    []string{"ssh/22"},
		Port:         22,
		IsActive:     true,
		NodesDisplay: []string{"/Default"},
	})
	handleObject(err, assets, "create asset")
	id := assets[0].ID
	for i := range assets {
		fmt.Println(assets[i].PrettyString())
	}

	// 2. Update jumpserver asset
	assets, err = jms.Asset().Asset().Update(&jumpserver.Asset{
		ID:       id,
		Hostname: hostname,
		Platform: jumpserver.Linux,
		IP:       "2.2.2.2",
	})
	handleObject(err, assets, "update assets")
	//time.Sleep(time.Second * 10)
	assets, err = jms.Asset().Asset().UpdatePartial(&jumpserver.Asset{
		ID:       id,
		Hostname: hostname,
		Platform: jumpserver.Linux,
		IP:       "3.3.3.3",
	})
	handleObject(err, assets, "partial update assets")
	//time.Sleep(time.Second * 10)

	// 3. List jumpserver assets.
	if assets, err = jms.Asset().Asset().List(nil); err != nil {
		panic(err)
	}
	handleObject(err, assets, "list all assets")
	assets, err = jms.Asset().Asset().List(&jumpserver.ListAssetParam{ID: id})
	handleObject(err, assets, "list assets by id")
	assets, err = jms.Asset().Asset().List(&jumpserver.ListAssetParam{Hostname: hostname})
	handleObject(err, assets, "list asset by hostname")
	assets, err = jms.Asset().Asset().List(&jumpserver.ListAssetParam{IP: ip})
	handleObject(err, assets, "list asset by ip")

	// 4. Get jumpserver assets
	asset, err = jms.Asset().Asset().Get(id)
	handleObject(err, asset, "get asset by id")

	// 5 Delete jumpserver asset by hostname.
	jms.Asset().Asset().Delete(&jumpserver.DeleteAssetParam{ID: id})
	handleObject(err, nil, "delete asset by id")
	jms.Asset().Asset().Create(&jumpserver.Asset{Hostname: hostname, IP: ip, Platform: jumpserver.Linux})
	jms.Asset().Asset().Delete(&jumpserver.DeleteAssetParam{Hostname: hostname})
	handleObject(err, nil, "delete asset by hostname")
	jms.Asset().Asset().Create(&jumpserver.Asset{Hostname: hostname, IP: ip, Platform: jumpserver.Linux})
	jms.Asset().Asset().Delete(&jumpserver.DeleteAssetParam{IP: ip})
	handleObject(err, nil, "delete asset by ip")

	// Output:
	//Successfully create asset (total 1)
	//Successfully update assets (total 2)
	//Successfully partial update assets (total 2)
	//Successfully list all assets (total 1)
	//Successfully list assets by id (total 1)
	//Successfully list asset by hostname (total 1)
	//Successfully list asset by ip (total 0)
	//Successfully get asset by id
	//Successfully delete asset by id
	//Successfully delete asset by hostname
	//Successfully delete asset by ip
}
