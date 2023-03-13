package main

import (
	"fmt"

	"github.com/forbearing/jumpserver"
)

func Asset_Domain(jms *jumpserver.Client) {
	var (
		name1    = "test1"
		name2    = "test2"
		name3    = "test3"
		comment1 = "comment1"
		comment2 = "comment2"
	)

	jms.Asset().Domain().Delete(&jumpserver.DomainDeleteParam{Name: name1})
	jms.Asset().Domain().Delete(&jumpserver.DomainDeleteParam{Name: name2})
	jms.Asset().Domain().Delete(&jumpserver.DomainDeleteParam{Name: name3})

	// 1.create
	domains, err := jms.Asset().Domain().Create(
		&jumpserver.Domain{Name: name1},
		&jumpserver.Domain{Name: name2},
		&jumpserver.Domain{Name: name3},
	)
	handleObject(err, domains, "create asset domains")
	ids := make(map[int]string)
	for i := range domains {
		ids[i] = domains[i].ID
		fmt.Println(domains[i].PrettyString())
	}

	// 2.update
	domains, err = jms.Asset().Domain().Update(
		&jumpserver.Domain{ID: ids[0], Name: name1, Comment: comment1},
		&jumpserver.Domain{ID: ids[1], Name: name2, Comment: comment1},
		&jumpserver.Domain{ID: ids[2], Name: name3, Comment: comment1},
	)
	handleObject(err, domains, "update asset domains")
	domains, err = jms.Asset().Domain().UpdatePartial(
		&jumpserver.Domain{ID: ids[0], Name: name1, Comment: comment2},
		&jumpserver.Domain{ID: ids[1], Name: name2, Comment: comment2},
		&jumpserver.Domain{ID: ids[2], Name: name3, Comment: comment2},
	)
	handleObject(err, domains, "partial update asset domains")

	// 3.list
	// list operation not support query by id
	domains, err = jms.Asset().Domain().List(nil)
	handleObject(err, domains, "list all asset domains")
	domains, err = jms.Asset().Domain().List(&jumpserver.DomainListParam{ID: ids[0]})
	handleObject(err, domains, "list asset domains by id")
	domains, err = jms.Asset().Domain().List(&jumpserver.DomainListParam{Name: name1})
	handleObject(err, domains, "list asset domains by name")

	// 4.get
	domain, err := jms.Asset().Domain().Get(ids[0])
	handleObject(err, domain, "get asset domain by id")

	// 5.delete
	err = jms.Asset().Domain().Delete(&jumpserver.DomainDeleteParam{ID: ids[0]})
	handleObject(err, nil, "delete asset domains by id")
	err = jms.Asset().Domain().Delete(&jumpserver.DomainDeleteParam{Name: name3})
	handleObject(err, nil, "delete asset domains by name")

	// Output

	// Successfully create asset domains (total 3)
	//
	//	{
	//	   "id": "36680af1-202b-46bc-917d-c5dc5b56bfaa",
	//	   "name": "test1",
	//	   "date_created": "2023/03/13 20:09:42 +0800",
	//	   "org_id": "00000000-0000-0000-0000-000000000002",
	//	   "org_name": "Default"
	//	}
	//
	//	{
	//	   "id": "810a83c6-6d81-48b8-a9cc-b7493bd53a2a",
	//	   "name": "test2",
	//	   "date_created": "2023/03/13 20:09:42 +0800",
	//	   "org_id": "00000000-0000-0000-0000-000000000002",
	//	   "org_name": "Default"
	//	}
	//
	//	{
	//	   "id": "8599f0a7-7467-4081-befa-3667d3e7440b",
	//	   "name": "test3",
	//	   "date_created": "2023/03/13 20:09:42 +0800",
	//	   "org_id": "00000000-0000-0000-0000-000000000002",
	//	   "org_name": "Default"
	//	}
	//
	// Successfully update asset domains (total 3)
	// Successfully partial update asset domains (total 3)
	// Successfully list all asset domains (total 3)
	// Successfully list asset domains by id (total 3)
	// Successfully list asset domains by name (total 1)
	// Successfully get asset domain by id
	// Successfully delete asset domains by id
	// Successfully delete asset domains by name
}
