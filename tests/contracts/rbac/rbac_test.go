package rbac

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"math/big"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/astaxie/beego"
	"testing"
	"github.com/scmo/apayment-backend/smart-contracts/rbac"
)

var farmerAuth *bind.TransactOpts
var inspectorAuth *bind.TransactOpts
var adminAuth *bind.TransactOpts
var systemAuth *bind.TransactOpts
var sim *backends.SimulatedBackend

var rbacContract *rbac.RBACContract

func init() {
	// Generate a new random account and a funded simulator
	key, _ := crypto.GenerateKey()
	farmerAuth = bind.NewKeyedTransactor(key)

	key, _ = crypto.GenerateKey()
	adminAuth = bind.NewKeyedTransactor(key)

	key, _ = crypto.GenerateKey()
	inspectorAuth = bind.NewKeyedTransactor(key)

	key, _ = crypto.GenerateKey()
	systemAuth = bind.NewKeyedTransactor(key)

	sim = backends.NewSimulatedBackend(core.GenesisAlloc{
		farmerAuth.From: {Balance: big.NewInt(10000000000)},
		adminAuth.From: {Balance: big.NewInt(10000000000)},
		systemAuth.From: {Balance: big.NewInt(10000000000)},
		inspectorAuth.From: {Balance: big.NewInt(10000000000)},
	})

}

func TestDeployRBACContract(t *testing.T) {
	beego.Trace("Test: ", "Deploy RBAC ContractRequest")
	_, _, rc, err := rbac.DeployRBACContract(systemAuth, sim)
	sim.Commit()

	Convey("Subject: Deploy Request-Contract\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldEqual, nil)
		})
	})
	rbacContract = rc
}

func TestAddFarmerRole(t *testing.T) {
	beego.Trace("Test: ", "Add farmer role with authorized user")
	_, err := rbacContract.AddFarmer(systemAuth, farmerAuth.From)
	sim.Commit()
	Convey("Subject: Add farmer (role) to Contract\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldEqual, nil)
		})
	})
}

func TestUserHasRole(t *testing.T) {
	isFarmer, err := rbacContract.IsFarmer(nil, farmerAuth.From)
	sim.Commit()
	Convey("Subject: Address has Farmer role?\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("Farmer has role 'Farmer'", func() {
			So(isFarmer, ShouldEqual, true)
		})
	})

	isNotFarmer, err := rbacContract.IsFarmer(nil, inspectorAuth.From)
	sim.Commit()
	Convey("Subject: Address has Farmer role?\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("Farmer is not anymore a farmer", func() {
			So(isNotFarmer, ShouldEqual, false)
		})
	})
}

func TestRemoveRole(t *testing.T) {
	_, err := rbacContract.RemoveFarmer(systemAuth, farmerAuth.From)
	sim.Commit()
	Convey("Subject: Remove farmer\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldEqual, nil)
		})
	})

	isNotFarmer, err := rbacContract.IsFarmer(nil, farmerAuth.From)
	sim.Commit()
	Convey("Subject: Address has Farmer role?\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("Farmer is not anymore a farmer", func() {
			So(isNotFarmer, ShouldEqual, false)
		})
	})
}