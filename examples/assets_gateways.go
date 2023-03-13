package main

import (
	"fmt"

	"github.com/forbearing/jumpserver"
)

func Asset_Gateway(jms *jumpserver.Client) {
	var (
		name1 = "test1"
		name2 = "test2"
		name3 = "test3"

		ip1 = "192.168.1.1"
		ip2 = "192.168.1.2"
		ip3 = "192.168.1.3"

		domain1 = "test1"
		domain2 = "test2"
		domain3 = "test3"
	)
	jms.Asset().Domain().Delete(&jumpserver.DomainDeleteParam{Name: domain1})
	jms.Asset().Domain().Delete(&jumpserver.DomainDeleteParam{Name: domain2})
	jms.Asset().Domain().Delete(&jumpserver.DomainDeleteParam{Name: domain3})
	jms.Asset().Gateway().Delete(&jumpserver.GatewayDeleteParam{Name: name1})
	jms.Asset().Gateway().Delete(&jumpserver.GatewayDeleteParam{Name: name2})
	jms.Asset().Gateway().Delete(&jumpserver.GatewayDeleteParam{Name: name3})

	domains, err := jms.Asset().Domain().Create(
		&jumpserver.Domain{Name: domain1},
		&jumpserver.Domain{Name: domain2},
		&jumpserver.Domain{Name: domain3},
	)
	if err != nil {
		panic(err)
	}
	domainIds := make(map[int]string)
	for i := range domains {
		domainIds[i] = domains[i].ID
	}

	// 1.create
	gateways, err := jms.Asset().Gateway().Create(
		&jumpserver.Gateway{Name: name1, IP: ip1, Domain: domainIds[0]},
		&jumpserver.Gateway{Name: name2, IP: ip2, Domain: domainIds[1]},
		&jumpserver.Gateway{Name: name3, IP: ip3, Domain: domainIds[2]},
	)
	handleObject(err, gateways, "create asset gateways")
	ids := make(map[int]string)
	for i := range gateways {
		ids[i] = gateways[i].ID
		fmt.Println(gateways[i].PrettyString())
	}

	// 2.update
	gateways, err = jms.Asset().Gateway().Update(
		&jumpserver.Gateway{ID: ids[0], Name: name1, IP: ip2, Domain: domainIds[0]},
		&jumpserver.Gateway{ID: ids[1], Name: name2, IP: ip3, Domain: domainIds[1]},
		&jumpserver.Gateway{ID: ids[2], Name: name3, IP: ip1, Domain: domainIds[2]},
	)
	handleObject(err, gateways, "update asset gateways")
	gateways, err = jms.Asset().Gateway().UpdatePartial(
		&jumpserver.Gateway{ID: ids[0], Name: name1, IP: ip2},
		&jumpserver.Gateway{ID: ids[1], Name: name2, IP: ip3},
		&jumpserver.Gateway{ID: ids[2], Name: name3, IP: ip1},
	)
	handleObject(err, gateways, "partial update asset gateways")

	// 3.list
	gateways, err = jms.Asset().Gateway().List(nil)
	handleObject(err, gateways, "list all asset gateways")
	gateways, err = jms.Asset().Gateway().List(&jumpserver.GatewayListParam{Name: name1})
	handleObject(err, gateways, "list asset gateways by name")
	gateways, err = jms.Asset().Gateway().List(&jumpserver.GatewayListParam{Domain: domainIds[2]})
	handleObject(err, gateways, "list asset gateways by domain")
	gateways, err = jms.Asset().Gateway().List(&jumpserver.GatewayListParam{IP: ip3})
	handleObject(err, gateways, "list asset gateways by ip")

	// 4.get
	gateway, err := jms.Asset().Gateway().Get(ids[0])
	handleObject(err, gateway, "get asset gateway")

	// 5.delete
	err = jms.Asset().Gateway().Delete(&jumpserver.GatewayDeleteParam{ID: ids[0]})
	handleObject(err, gateway, "delete asset gateway by id")
	err = jms.Asset().Gateway().Delete(&jumpserver.GatewayDeleteParam{Name: name2})
	handleObject(err, gateway, "delete asset gateway by name")
	err = jms.Asset().Gateway().Delete(&jumpserver.GatewayDeleteParam{DomainName: domain3})
	handleObject(err, gateway, "delete asset gateway by domain name")

	/*
		Output:

		Successfully create asset gateways (total 3)
		{
		    "id": "a4a02930-602c-487d-8409-d47792f8b4ad",
		    "name": "test1",
		    "ip": "192.168.1.1",
		    "port": 22,
		    "protocol": "ssh",
		    "is_active": true,
		    "is_connective": true,
		    "date_created": "2023/03/13 22:48:24 +0800",
		    "date_updated": "2023/03/13 22:48:24 +0800",
		    "created_by": "Administrator",
		    "domain": "d476e51c-b3c9-404d-931f-f33effc73941",
		    "org_id": "00000000-0000-0000-0000-000000000002",
		    "org_name": "Default"
		}
		{
		    "id": "eb055711-020c-4752-8e9e-3838cf977f58",
		    "name": "test2",
		    "ip": "192.168.1.2",
		    "port": 22,
		    "protocol": "ssh",
		    "is_active": true,
		    "is_connective": true,
		    "date_created": "2023/03/13 22:48:24 +0800",
		    "date_updated": "2023/03/13 22:48:24 +0800",
		    "created_by": "Administrator",
		    "domain": "aeb73e2a-0fe5-46f1-8f02-9d3ddf02cd6f",
		    "org_id": "00000000-0000-0000-0000-000000000002",
		    "org_name": "Default"
		}
		{
		    "id": "85600b2a-52cf-4f0c-8838-41757be8f20c",
		    "name": "test3",
		    "ip": "192.168.1.3",
		    "port": 22,
		    "protocol": "ssh",
		    "is_active": true,
		    "is_connective": true,
		    "date_created": "2023/03/13 22:48:24 +0800",
		    "date_updated": "2023/03/13 22:48:24 +0800",
		    "created_by": "Administrator",
		    "domain": "18fa2f57-71c1-497c-8f06-2855a3a415b6",
		    "org_id": "00000000-0000-0000-0000-000000000002",
		    "org_name": "Default"
		}
		Successfully update asset gateways (total 3)
		Successfully partial update asset gateways (total 3)
		Successfully list all asset gateways (total 3)
		Successfully list asset gateways by name (total 1)
		Successfully list asset gateways by domain (total 1)
		Successfully list asset gateways by ip (total 1)
		Successfully get asset gateway
		Successfully delete asset gateway by id
		Successfully delete asset gateway by name
		Successfully delete asset gateway by domain name
	*/
}
