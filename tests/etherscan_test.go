package test

import (
	"path/filepath"
	"github.com/astaxie/beego"
	"runtime"
	"testing"
	"github.com/scmo/apayment-backend/services"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func Test_GetTransactions(t *testing.T) {

	services.GetTransactions()
	//Convey("Subject: Test TVD Version\n", t, func() {
	//	Convey("Error should be nil", func() {
	//		So(err, ShouldEqual, nil)
	//	})
	//	Convey("Version should not be null (= '')", func() {
	//		So(len(versionResponse.VersionResult), ShouldBeGreaterThan, 0)
	//	})
	//})
}