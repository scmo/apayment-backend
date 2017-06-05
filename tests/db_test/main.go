// +build !test
package db_test

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"github.com/scmo/apayment-backend/db"
)

func Setup() {
	beego.Info("Initialize Database")
	// Register Driver
	orm.RegisterDriver("postgres", orm.DRPostgres)

	dataSource := "port=9032 user=postgres password=test123456 dbname=db_foodchain_test sslmode=disable"
	//dataSource := "user=postgres password=test123456 dbname=db_foodchain_test sslmode=disable"

	beego.Info(dataSource)

	// set default database
	err := orm.RegisterDataBase("default", "postgres", dataSource, 30, 30)
	// Error.
	err = orm.RunSyncdb("default", true, false)
	if err != nil {
		beego.Error("RunSyncdb Error")
	}
	// Log every SQL Query
	orm.Debug = false

	// Populate DB
	db.Seed_LegalForm()
	db.Seed_PlantType()

	db.Seed_Contributions()
	db.Seed_ControlPoints()
}