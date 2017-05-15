package test

import (
	"runtime"
	"path/filepath"
	_ "github.com/scmo/foodchain-backend/routers"
	"github.com/astaxie/beego"
)


func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)

	//ethereum_test.Init()
	//beego.Debug(TestBackend.FarmerAuth.From.String())
	//db_test.Init()
}



