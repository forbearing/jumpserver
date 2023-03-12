package main

import (
	"github.com/forbearing/jumpserver"
)

func Example_Asset(jms *jumpserver.Client) {
	var (
		asset    *jumpserver.Asset
		assets   []*jumpserver.Asset
		hostname = "myserver"
		ip       = "1.1.1.1"
		err      error
	)
	var _, _ = asset, assets
	jms.Asset().Delete(&jumpserver.DeleteAssetParam{Hostname: hostname})

	// 1. Create jumpserver asset.
	assets, err = jms.Asset().Create(&jumpserver.Asset{
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

	// 2. Update jumpserver asset
	assets, err = jms.Asset().Update(&jumpserver.Asset{
		ID:       id,
		Hostname: hostname,
		Platform: jumpserver.Linux,
		IP:       "2.2.2.2",
	})
	handleObject(err, assets, "update assets")
	//time.Sleep(time.Second * 10)
	assets, err = jms.Asset().UpdatePartial(&jumpserver.Asset{
		ID:       id,
		Hostname: hostname,
		Platform: jumpserver.Linux,
		IP:       "3.3.3.3",
	})
	handleObject(err, assets, "partial update assets")
	//time.Sleep(time.Second * 10)

	// 3. List jumpserver assets.
	if assets, err = jms.Asset().List(); err != nil {
		panic(err)
	}
	handleObject(err, assets, "list asset")

	// 4. Get jumpserver assets
	assets, err = jms.Asset().Get(&jumpserver.GetAssetParam{ID: id})
	handleObject(err, assets, "get asset by id")
	assets, err = jms.Asset().Get(&jumpserver.GetAssetParam{Hostname: hostname})
	handleObject(err, assets, "get asset by hostname")
	assets, err = jms.Asset().Get(&jumpserver.GetAssetParam{IP: ip})
	handleObject(err, assets, "get asset by ip")

	// 5 Delete jumpserver asset by hostname.
	jms.Asset().Delete(&jumpserver.DeleteAssetParam{ID: id})
	handleObject(err, nil, "delete asset by id")
	jms.Asset().Create(&jumpserver.Asset{Hostname: hostname, IP: ip, Platform: jumpserver.Linux})
	jms.Asset().Delete(&jumpserver.DeleteAssetParam{Hostname: hostname})
	handleObject(err, nil, "delete asset by hostname")
	jms.Asset().Create(&jumpserver.Asset{Hostname: hostname, IP: ip, Platform: jumpserver.Linux})
	jms.Asset().Delete(&jumpserver.DeleteAssetParam{IP: ip})
	handleObject(err, nil, "delete asset by ip")

	// Output:
	//Successfully create asset (total 1)
	//Successfully update assets (total 1)
	//Successfully partial update assets (total 1)
	//Successfully list asset (total 1)
	//Successfully get asset by id (total 1)
	//Successfully get asset by hostname (total 1)
	//Successfully get asset by ip (total 0)
	//Successfully delete asset by id
	//Successfully delete asset by hostname
	//Successfully delete asset by ip
}
