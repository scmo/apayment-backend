package test

import (
	"runtime"
	"path/filepath"
	_ "github.com/scmo/foodchain-backend/routers"

	"github.com/astaxie/beego"
	"testing"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"math/big"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/scmo/foodchain-backend/smart-contracts/request"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestEthereum (t * testing.T) {
	// Generate a new random account and a funded simulator
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)

	sim := backends.NewSimulatedBackend(core.GenesisAccount{ Balance: big.NewInt(10000000000)})

	// Deploy a token contract on the simulated blockchain

	smartcontracts.DeployRequestContract(auth, sim,  12, nil, "asdf")


}
//// TestGet is a sample to run an endpoint test
//func TestGet(t *testing.T) {
//	r, _ := http.NewRequest("GET", "/v1/object", nil)
//	w := httptest.NewRecorder()
//	beego.BeeApp.Handlers.ServeHTTP(w, r)
//
//	beego.Trace("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())
//
//	Convey("Subject: Test Station Endpoint\n", t, func() {
//	        Convey("Status Code Should Be 200", func() {
//	                So(w.Code, ShouldEqual, 200)
//	        })
//	        Convey("The Result Should Not Be Empty", func() {
//	                So(w.Body.Len(), ShouldBeGreaterThan, 0)
//	        })
//	})
//}

