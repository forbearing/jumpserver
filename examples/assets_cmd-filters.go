package main

import (
	"fmt"

	"github.com/forbearing/jumpserver"
)

func Asset_CmdFilter(jms *jumpserver.Client) {
	var (
		name1    = "test1"
		name2    = "test2"
		name3    = "test3"
		comment1 = "just for test"
		comment2 = "update comment"
	)

	jms.Asset().CmdFilter().Delete(&jumpserver.CmdFilterDeleteParam{Name: name1})
	jms.Asset().CmdFilter().Delete(&jumpserver.CmdFilterDeleteParam{Name: name2})
	jms.Asset().CmdFilter().Delete(&jumpserver.CmdFilterDeleteParam{Name: name3})

	// 1.create
	filters, err := jms.Asset().CmdFilter().Create(
		&jumpserver.CmdFilter{Name: name1, Comment: comment1},
		&jumpserver.CmdFilter{Name: name2, Comment: comment1},
		&jumpserver.CmdFilter{Name: name3, Comment: comment1},
	)
	handleObject(err, filters, "create asset cmd filters")
	ids := make(map[int]string)
	for i := range filters {
		fmt.Println(filters[i].PrettyString())
		ids[i] = filters[i].ID
	}

	// 2.update
	// id is required field.
	filters, err = jms.Asset().CmdFilter().Update(
		&jumpserver.CmdFilter{ID: ids[0], Name: name1, Comment: comment2},
		&jumpserver.CmdFilter{ID: ids[1], Name: name2, Comment: comment2},
		&jumpserver.CmdFilter{ID: ids[2], Name: name3, Comment: comment2},
	)
	handleObject(err, filters, "update asset cmd filters")
	filters, err = jms.Asset().CmdFilter().UpdatePartial(
		&jumpserver.CmdFilter{ID: ids[0], Name: name1, Comment: comment2},
		&jumpserver.CmdFilter{ID: ids[1], Name: name2, Comment: comment2},
		&jumpserver.CmdFilter{ID: ids[2], Name: name3, Comment: comment2},
	)
	handleObject(err, filters, "partial update asset cmd filters")

	// 3.list
	filters, err = jms.Asset().CmdFilter().List(nil)
	handleObject(err, filters, "list all asset cmd filters")
	// not support list by id(default list all)
	filters, err = jms.Asset().CmdFilter().List(&jumpserver.CmdFilterListParam{ID: ids[0]})
	handleObject(err, filters, "list asset cmd filters by id")
	filters, err = jms.Asset().CmdFilter().List(&jumpserver.CmdFilterListParam{Name: name2})
	handleObject(err, filters, "list asset cmd filters by name")

	// 4.get
	filter, err := jms.Asset().CmdFilter().Get(ids[0])
	handleObject(err, filter, "get asset cmd filter")

	// 5.delete
	err = jms.Asset().CmdFilter().Delete(&jumpserver.CmdFilterDeleteParam{ID: ids[0]})
	handleObject(err, nil, "delete asset cmd filter by id")
	err = jms.Asset().CmdFilter().Delete(&jumpserver.CmdFilterDeleteParam{Name: name2})
	handleObject(err, nil, "delete asset cmd filter by name")

	// Output

	//Successfully create asset cmd filters (total 3)
	//{
	//    "id": "94e95556-9d2f-4b46-bfeb-57f1a73b11e6",
	//    "name": "test1",
	//    "org_id": "00000000-0000-0000-0000-000000000002",
	//    "org_name": "Default",
	//    "is_active": true,
	//    "date_created": "2023/03/13 18:40:30 +0800",
	//    "date_updated": "2023/03/13 18:40:30 +0800",
	//    "comment": "just for test",
	//    "created_by": "Administrator"
	//}
	//{
	//    "id": "847ea6c7-0611-480c-9aef-c19598142c0e",
	//    "name": "test2",
	//    "org_id": "00000000-0000-0000-0000-000000000002",
	//    "org_name": "Default",
	//    "is_active": true,
	//    "date_created": "2023/03/13 18:40:30 +0800",
	//    "date_updated": "2023/03/13 18:40:30 +0800",
	//    "comment": "just for test",
	//    "created_by": "Administrator"
	//}
	//{
	//    "id": "7fad8e82-12c0-4f8a-9e5e-26b0583a97dc",
	//    "name": "test3",
	//    "org_id": "00000000-0000-0000-0000-000000000002",
	//    "org_name": "Default",
	//    "is_active": true,
	//    "date_created": "2023/03/13 18:40:30 +0800",
	//    "date_updated": "2023/03/13 18:40:30 +0800",
	//    "comment": "just for test",
	//    "created_by": "Administrator"
	//}
	//Successfully update asset cmd filters (total 3)
	//Successfully partial update asset cmd filters (total 3)
	//Successfully list all asset cmd filters (total 3)
	//Successfully list asset cmd filters by id (total 3)
	//Successfully list asset cmd filters by name (total 1)
	//Successfully get asset cmd filter
	//Successfully delete asset cmd filter by id
	//Successfully delete asset cmd filter by name
}
