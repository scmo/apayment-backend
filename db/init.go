package db

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
)

func Init() {
	beego.Info("Initialize Database")
	// Register Driver
	orm.RegisterDriver("postgres", orm.DRPostgres)

	dataSource := "user=" + beego.AppConfig.String("db_user") + " password=" + beego.AppConfig.String("db_password") + " dbname=" + beego.AppConfig.String("db_name") + " sslmode=" + beego.AppConfig.String("db_sslmode")

	beego.Info(dataSource)

	max_idle, _ := beego.AppConfig.Int("max_idle")
	max_conn, _ := beego.AppConfig.Int("max_conn")

	force_table_recreate, _ := beego.AppConfig.Bool("force_table_recreate")
	db_log_verbose, _ := beego.AppConfig.Bool("db_log_verbose")

	// set default database
	orm.RegisterDataBase("default", "postgres", dataSource, max_idle, max_conn)

	// Error.
	err := orm.RunSyncdb("default", force_table_recreate, db_log_verbose)
	if err != nil {
		beego.Error("RunSyncdb Error")
	}


	// Populate DB
	SeedContributions()
	SeedInspectionCriterion()
}

