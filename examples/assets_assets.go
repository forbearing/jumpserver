package main

import (
	"fmt"

	"github.com/forbearing/jumpserver"
)

func Asset_Asset(jms *jumpserver.Client) {
	var (
		hostname1 = "server1"
		hostname2 = "server2"
		hostname3 = "server3"

		ip1 = "1.1.1.1"
		ip2 = "2.2.2.2"
		ip3 = "3.3.3.3"

		protocols    = []string{"ssh/22"}
		nodesDisplay = []string{"/Default"}
	)

	jms.Asset().Asset().Delete(&jumpserver.DeleteAssetParam{Hostname: hostname1})
	jms.Asset().Asset().Delete(&jumpserver.DeleteAssetParam{Hostname: hostname2})
	jms.Asset().Asset().Delete(&jumpserver.DeleteAssetParam{Hostname: hostname3})

	// 1. Create jumpserver asset.
	assets, err := jms.Asset().Asset().Create(
		&jumpserver.Asset{Hostname: hostname1, IP: ip1, Platform: jumpserver.Linux, Protocols: protocols, Port: 22, IsActive: true, NodesDisplay: nodesDisplay},
		&jumpserver.Asset{Hostname: hostname2, IP: ip2, Platform: jumpserver.Linux, Protocols: protocols, Port: 22, IsActive: true, NodesDisplay: nodesDisplay},
		&jumpserver.Asset{Hostname: hostname3, IP: ip3, Platform: jumpserver.Linux, Protocols: protocols, Port: 22, IsActive: true, NodesDisplay: nodesDisplay},
	)
	handleObject(err, assets, "create assets")
	ids := make(map[int]string)
	for i := range assets {
		ids[i] = assets[i].ID
		fmt.Println(assets[i].PrettyString())
	}

	// 2. Update jumpserver asset
	// asset id is required field
	assets, err = jms.Asset().Asset().Update( // update to windows os.
		&jumpserver.Asset{ID: ids[0], Hostname: hostname1, IP: ip1, Platform: jumpserver.Windows, Protocols: protocols, Port: 22, IsActive: true, NodesDisplay: nodesDisplay},
		&jumpserver.Asset{ID: ids[1], Hostname: hostname2, IP: ip2, Platform: jumpserver.Windows, Protocols: protocols, Port: 22, IsActive: true, NodesDisplay: nodesDisplay},
		&jumpserver.Asset{ID: ids[2], Hostname: hostname3, IP: ip3, Platform: jumpserver.Windows, Protocols: protocols, Port: 22, IsActive: true, NodesDisplay: nodesDisplay},
	)
	// asset id is required field
	handleObject(err, assets, "update assets")
	assets, err = jms.Asset().Asset().UpdatePartial( // update to BSD os
		&jumpserver.Asset{ID: ids[0], Hostname: hostname1, IP: ip1, Platform: jumpserver.BSD},
		&jumpserver.Asset{ID: ids[1], Hostname: hostname2, IP: ip2, Platform: jumpserver.BSD},
		&jumpserver.Asset{ID: ids[2], Hostname: hostname3, IP: ip3, Platform: jumpserver.BSD},
	)
	handleObject(err, assets, "partial update assets")

	// 3. List jumpserver assets.
	assets, err = jms.Asset().Asset().List(nil)
	handleObject(err, assets, "list all assets")

	assets, err = jms.Asset().Asset().List(&jumpserver.ListAssetParam{ID: ids[0]})
	handleObject(err, assets, "list assets by id")
	assets, err = jms.Asset().Asset().List(&jumpserver.ListAssetParam{Hostname: hostname1})
	handleObject(err, assets, "list asset by hostname")
	assets, err = jms.Asset().Asset().List(&jumpserver.ListAssetParam{IP: ip1})
	handleObject(err, assets, "list asset by ip")
	assets, err = jms.Asset().Asset().List(&jumpserver.ListAssetParam{IsActive: true})
	handleObject(err, assets, "list asset by IsActive")

	// 4. Get jumpserver assets
	asset, err := jms.Asset().Asset().Get(ids[0])
	handleObject(err, asset, "get asset by id")

	// 5 Delete jumpserver asset by hostname.
	jms.Asset().Asset().Delete(&jumpserver.DeleteAssetParam{ID: ids[0]})
	handleObject(err, nil, "delete asset by id")
	jms.Asset().Asset().Delete(&jumpserver.DeleteAssetParam{Hostname: hostname2})
	handleObject(err, nil, "delete asset by hostname")
	jms.Asset().Asset().Delete(&jumpserver.DeleteAssetParam{IP: ip3})
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

	/*
		Output

		Successfully create assets (total 3)
		{
		    "id": "d460227e-8fd9-431b-a3dd-7e61f16c6f5f",
		    "hostname": "server1",
		    "ip": "1.1.1.1",
		    "platform": "Linux",
		    "protocols": [
		        "ssh/22"
		    ],
		    "is_active": true,
		    "nodes": [
		        "27e08786-5281-45f3-93de-08a0a5cac1d3"
		    ],
		    "nodes_display": [
		        "/Default"
		    ],
		    "connectivity": "unknown",
		    "created_by": "Administrator",
		    "date_created": "2023/03/13 21:12:50 +0800",
		    "org_id": "00000000-0000-0000-0000-000000000002",
		    "org_name": "Default"
		}
		{
		    "id": "c8d52f66-79fc-45e0-bd48-afc08c016676",
		    "hostname": "server2",
		    "ip": "2.2.2.2",
		    "platform": "Linux",
		    "protocols": [
		        "ssh/22"
		    ],
		    "is_active": true,
		    "nodes": [
		        "27e08786-5281-45f3-93de-08a0a5cac1d3"
		    ],
		    "nodes_display": [
		        "/Default"
		    ],
		    "connectivity": "unknown",
		    "created_by": "Administrator",
		    "date_created": "2023/03/13 21:12:50 +0800",
		    "org_id": "00000000-0000-0000-0000-000000000002",
		    "org_name": "Default"
		}
		{
		    "id": "89077ffe-73eb-4c04-a6c7-26435f9279ed",
		    "hostname": "server3",
		    "ip": "3.3.3.3",
		    "platform": "Linux",
		    "protocols": [
		        "ssh/22"
		    ],
		    "is_active": true,
		    "nodes": [
		        "27e08786-5281-45f3-93de-08a0a5cac1d3"
		    ],
		    "nodes_display": [
		        "/Default"
		    ],
		    "connectivity": "unknown",
		    "created_by": "Administrator",
		    "date_created": "2023/03/13 21:12:50 +0800",
		    "org_id": "00000000-0000-0000-0000-000000000002",
		    "org_name": "Default"
		}
		Successfully update assets (total 3)
		Successfully partial update assets (total 3)
		Successfully list all assets (total 4)
		Successfully list assets by id (total 4)
		Successfully list asset by hostname (total 1)
		Successfully list asset by ip (total 1)
		Successfully list asset by IsActive (total 4)
		Successfully get asset by id
		Successfully delete asset by id
		Successfully delete asset by hostname
		Successfully delete asset by ip
	*/
}
