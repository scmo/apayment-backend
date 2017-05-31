package test

import (
	"testing"
	_ "github.com/scmo/foodchain-backend/routers"
	"github.com/astaxie/beego"

	"path/filepath"
	"runtime"
	"github.com/scmo/foodchain-backend/services/tvd"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// Test soap request
func TestSoap(t *testing.T) {
	auth := tvd.BasicAuth{
		Login: "2000319",
		Password:"qikCloud2017",
	}
	animalTracingPortType := tvd.NewAnimalTracingPortType("https://ws-in.wbf.admin.ch/Livestock/AnimalTracing/1", true, &auth)

	versionRequest := tvd.Version{
		P_ManufacturerKey:"bebc4e6a-2477-4ec6-8837-d503a87e85f2",
	}
	versionResponse, err := animalTracingPortType.Version(&versionRequest)

	beego.Debug(versionResponse, err)

}