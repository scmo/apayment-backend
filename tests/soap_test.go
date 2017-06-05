package test

import (
	"testing"
	_ "github.com/scmo/apayment-backend/routers"
	"github.com/astaxie/beego"

	"path/filepath"
	"runtime"
	"github.com/scmo/apayment-backend/services/tvd"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// Test soap request
func Test_TvdVersion(t *testing.T) {
	auth := tvd.BasicAuth{
		Login: beego.AppConfig.String("tvd_username"),
		Password:beego.AppConfig.String("tvd_password"),
	}
	animalTracingPortType := tvd.NewAnimalTracingPortType("https://ws-in.wbf.admin.ch/Livestock/AnimalTracing/1", true, &auth)

	versionRequest := tvd.Version{
		PManufacturerKey:beego.AppConfig.String("tvd_manufacturerKey"),
		Action:"http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1/AnimalTracingPortType/Version",
	}
	versionResponse, err := animalTracingPortType.Version(&versionRequest)

	Convey("Subject: Test TVD Version\n", t, func() {
		Convey("Error should be nil", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("Version should not be null (= '')", func() {
			So(len(versionResponse.VersionResult), ShouldBeGreaterThan, 0)
		})

	})
}