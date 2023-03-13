package main

import (
	"fmt"

	"github.com/forbearing/jumpserver"
)

func Asset_AccountBackupPlan(jms *jumpserver.Client) {
	var (
		name1    = "test-1"
		name2    = "test-2"
		name3    = "test-3"
		crontab  = "*/1 * * * *"
		crontab2 = "*/2 * * * *"
		comment  = "test comment"
	)

	jms.Asset().AccountBackupPlan().Delete(&jumpserver.AccountBackupPlanDeleteParam{Name: name1})
	jms.Asset().AccountBackupPlan().Delete(&jumpserver.AccountBackupPlanDeleteParam{Name: name2})
	jms.Asset().AccountBackupPlan().Delete(&jumpserver.AccountBackupPlanDeleteParam{Name: name3})

	// 1.create
	plans, err := jms.Asset().AccountBackupPlan().Create(
		&jumpserver.AccountBackupPlan{Name: name1, Crontab: crontab},
		&jumpserver.AccountBackupPlan{Name: name2, Crontab: crontab},
		&jumpserver.AccountBackupPlan{Name: name3, Crontab: crontab},
	)
	handleObject(err, plans, "create account backup plans")
	ids := make(map[int]string)
	for i := range plans {
		fmt.Println(plans[i].PrettyString())
		ids[i] = plans[i].ID
	}

	// 2.update
	// plan's "Crontab" field is must set.
	plans, err = jms.Asset().AccountBackupPlan().Update(
		&jumpserver.AccountBackupPlan{ID: ids[0], Name: name1, Crontab: crontab2, Comment: comment},
		&jumpserver.AccountBackupPlan{ID: ids[1], Name: name2, Crontab: crontab2, Comment: comment},
		&jumpserver.AccountBackupPlan{ID: ids[2], Name: name3, Crontab: crontab2, Comment: comment},
	)
	// partial update
	// plan's "Crontab" field is optional.
	handleObject(err, plans, "update account backup plans")
	plans, err = jms.Asset().AccountBackupPlan().UpdatePartial(
		&jumpserver.AccountBackupPlan{ID: ids[0], Name: name1, Comment: comment},
		&jumpserver.AccountBackupPlan{ID: ids[1], Name: name2, Comment: comment},
		&jumpserver.AccountBackupPlan{ID: ids[2], Name: name3, Comment: comment},
	)
	handleObject(err, plans, "partial update account backup plans")

	// 3.list
	plans, err = jms.Asset().AccountBackupPlan().List(nil)
	handleObject(err, plans, "list all account backup plans")
	plans, err = jms.Asset().AccountBackupPlan().List(&jumpserver.AccountBackupPlanListParam{Name: name1})
	handleObject(err, plans, "list account backup plans by name")
	// 4.get
	plan, err := jms.Asset().AccountBackupPlan().Get(ids[0])
	handleObject(err, plan, "get account backup plans")

	// 5.delete
	err = jms.Asset().AccountBackupPlan().Delete(&jumpserver.AccountBackupPlanDeleteParam{Name: name1})
	handleObject(err, nil, "delete account backup plan by name")
	jms.Asset().AccountBackupPlan().Delete(&jumpserver.AccountBackupPlanDeleteParam{ID: ids[3]})
	handleObject(err, nil, "delete account backup plan by id")

	// Output
	//Successfully create account backup plans (total 3)
	//{
	//    "id": "3e475956-3277-4d17-becf-6d4e269140e8",
	//    "name": "test-1",
	//    "is_periodic": true,
	//    "interval": 24,
	//    "crontab": "*/1 * * * *",
	//    "date_created": "2023/03/13 11:27:43 +0800",
	//    "date_updated": "2023/03/13 11:27:43 +0800",
	//    "created_by": "Administrator",
	//    "periodic_display": "定期执行 ( */1 * * * * )",
	//    "types": [
	//        "all",
	//        "asset",
	//        "application"
	//    ],
	//    "org_id": "00000000-0000-0000-0000-000000000002",
	//    "org_name": "Default"
	//}
	//{
	//    "id": "7d271006-419c-4fba-97f2-0708a1a202ef",
	//    "name": "test-2",
	//    "is_periodic": true,
	//    "interval": 24,
	//    "crontab": "*/1 * * * *",
	//    "date_created": "2023/03/13 11:27:43 +0800",
	//    "date_updated": "2023/03/13 11:27:43 +0800",
	//    "created_by": "Administrator",
	//    "periodic_display": "定期执行 ( */1 * * * * )",
	//    "types": [
	//        "all",
	//        "asset",
	//        "application"
	//    ],
	//    "org_id": "00000000-0000-0000-0000-000000000002",
	//    "org_name": "Default"
	//}
	//{
	//    "id": "7dad30a1-6330-4093-b393-a192f55d6139",
	//    "name": "test-3",
	//    "is_periodic": true,
	//    "interval": 24,
	//    "crontab": "*/1 * * * *",
	//    "date_created": "2023/03/13 11:27:43 +0800",
	//    "date_updated": "2023/03/13 11:27:43 +0800",
	//    "created_by": "Administrator",
	//    "periodic_display": "定期执行 ( */1 * * * * )",
	//    "types": [
	//        "all",
	//        "asset",
	//        "application"
	//    ],
	//    "org_id": "00000000-0000-0000-0000-000000000002",
	//    "org_name": "Default"
	//}
	//Successfully update account backup plans (total 3)
	//Successfully partial update account backup plans (total 3)
	//Successfully list all account backup plans (total 4)
	//Successfully list account backup plans by name (total 1)
	//Successfully get account backup plans
	//Successfully delete account backup plan by name
	//Successfully delete account backup plan by id
}
