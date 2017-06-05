package integrationtests
//
//import (
//	"path/filepath"
//	"github.com/astaxie/beego"
//	"net/http/httptest"
//
//	"testing"
//	"runtime"
//	"net/http"
//	_ "github.com/scmo/apayment-backend/routers"
//	. "github.com/smartystreets/goconvey/convey"
//	"strings"
//	"github.com/scmo/apayment-backend/tests/db_test"
//	"encoding/json"
//	"github.com/ethereum/go-ethereum/accounts/abi/bind"
//	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
//	"github.com/ethereum/go-ethereum/core"
//	"math/big"
//	"github.com/ethereum/go-ethereum/crypto"
//	"github.com/scmo/apayment-backend/services"
//	"github.com/scmo/apayment-backend/models"
//)
//
//
//var farmerAuth *bind.TransactOpts
//var inspectorAuth *bind.TransactOpts
//var adminAuth *bind.TransactOpts
//var systemAuth *bind.TransactOpts
//
//var sim *backends.SimulatedBackend
//
//func init() {
//	_, file, _, _ := runtime.Caller(1)
//	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, "../.." + string(filepath.Separator))))
//	beego.TestBeegoInit(apppath)
//	db_test.Setup()
//	initSimulatedBackend()
//	seed_Users()
//}
//
//func initSimulatedBackend() {
//	// Generate a new random account and a funded simulator
//	key, _ := crypto.GenerateKey()
//	farmerAuth = bind.NewKeyedTransactor(key)
//
//	key, _ = crypto.GenerateKey()
//	adminAuth = bind.NewKeyedTransactor(key)
//
//	key, _ = crypto.GenerateKey()
//	inspectorAuth = bind.NewKeyedTransactor(key)
//
//	key, _ = crypto.GenerateKey()
//	systemAuth = bind.NewKeyedTransactor(key)
//
//	sim = backends.NewSimulatedBackend(core.GenesisAlloc{
//		farmerAuth.From: {Balance: big.NewInt(10000000000)},
//		adminAuth.From: {Balance: big.NewInt(10000000000)},
//		systemAuth.From: {Balance: big.NewInt(10000000000)},
//		inspectorAuth.From: {Balance: big.NewInt(10000000000)},
//	})
//}
//
//func seed_Users() {
//	if cnt, _ := services.CountRoles(); cnt > 0 {
//		return
//	}
//	admin := models.Role{Name:"Admin"}
//	farmer := models.Role{Name:"Farmer"}
//	canton := models.Role{Name:"Canton"}
//	inspector := models.Role{Name:"Inspector"}
//	bund := models.Role{Name:"bund"}
//
//	services.CreateRole(&admin);
//	services.CreateRole(&farmer);
//	services.CreateRole(&canton);
//	services.CreateRole(&inspector);
//	services.CreateRole(&bund);
//
//	if cnt, _ := services.CountUsers(); cnt > 0 {
//		return
//	}
//
//	roles := make([]*models.Role, 1)
//	roles[0] = &farmer
//	services.CreateUser(&models.User{Username:"farmer1", Password:"farmer1", Firstname: "Florian", Lastname:"Meisterhans", Address: farmerAuth.From.String(), Roles: roles})
//
//	roles = make([]*models.Role, 1)
//	roles[0] = &inspector
//	services.CreateUser(&models.User{Username:"inspect", Password:"inspect", Firstname: "Inspector", Lastname: "Gadget", Address: inspectorAuth.From.String(), Roles: roles})
//
//	roles = make([]*models.Role, 1)
//	roles[0] = &admin
//	services.CreateUser(&models.User{Username:"admin", Password:"admin", Firstname: "Admin", Lastname: "Admin", Address: adminAuth.From.String(), Roles: roles})
//
//}
//
//
//
//
//type LoginResponse struct {
//	Token string `json:"token"`
//}
//var lr LoginResponse
//func ssTestLogin(t *testing.T){
//	r, _ := http.NewRequest("POST", "/v1/user/login", strings.NewReader("{\"username\": \"farmer1\", \"password\":\"farmer1\"}" ))
//	w := httptest.NewRecorder()
//	beego.BeeApp.Handlers.ServeHTTP(w, r)
//	json.Unmarshal(w.Body.Bytes(), &lr)
//	beego.Trace("Test:", "Login", "Code:", w.Code, "Response:", w.Body.String())
//	Convey("Subject: Test Login Endpoint\n", t, func() {
//		Convey("Status Code Should Be 200", func() {
//			So(w.Code, ShouldEqual, 200)
//		})
//		Convey("Response contains token", func() {
//			So(lr.Token, ShouldNotEqual, nil)
//		})
//	})
//	services.CreateRequest(models.Request{}, nil)
//}
