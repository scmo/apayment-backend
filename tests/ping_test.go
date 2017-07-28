package test

import (
	"encoding/json"
	"github.com/astaxie/beego"
	_ "github.com/scmo/apayment-backend/routers"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// Test ping endpoint
func TestPing(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/ping", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestPing", "Code:", w.Code, "Response:", w.Body.String())

	Convey("Subject: Test Ping Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result be 'pong'", func() {
			pong, _ := json.Marshal("pong")
			So(w.Body.String(), ShouldEqual, string(pong))
		})
	})
}
