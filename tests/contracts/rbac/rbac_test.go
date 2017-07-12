package rbac

import (
	"github.com/astaxie/beego"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/scmo/apayment-backend/smart-contracts/rbac"
	. "github.com/smartystreets/goconvey/convey"
	"math/big"
	"testing"
)

var farmerAuth *bind.TransactOpts
var inspectorAuth *bind.TransactOpts
var adminAuth *bind.TransactOpts
var cantonAuth *bind.TransactOpts
var systemAuth *bind.TransactOpts
var sim *backends.SimulatedBackend

var rbacContract *rbac.RBACContract

/*
 Prepare Simulated Blockchain inclusive accounts
*/
func init() {
	// Generate a new random account and a funded simulator
	key, _ := crypto.GenerateKey()
	farmerAuth = bind.NewKeyedTransactor(key)

	key, _ = crypto.GenerateKey()
	adminAuth = bind.NewKeyedTransactor(key)

	key, _ = crypto.GenerateKey()
	cantonAuth = bind.NewKeyedTransactor(key)

	key, _ = crypto.GenerateKey()
	inspectorAuth = bind.NewKeyedTransactor(key)

	key, _ = crypto.GenerateKey()
	systemAuth = bind.NewKeyedTransactor(key)

	sim = backends.NewSimulatedBackend(core.GenesisAlloc{
		farmerAuth.From:    {Balance: big.NewInt(10000000000)},
		adminAuth.From:     {Balance: big.NewInt(10000000000)},
		cantonAuth.From:    {Balance: big.NewInt(10000000000)},
		systemAuth.From:    {Balance: big.NewInt(10000000000)},
		inspectorAuth.From: {Balance: big.NewInt(10000000000)},
	})

}

/*
 Deploy RBAC contract
*/
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

/*
 Add Role 'Farmer' to farmer
*/
func Test_AddFarmerRole(t *testing.T) {
	beego.Trace("Test: ", "Add farmer role with authorized user")
	_, err := rbacContract.AddFarmer(systemAuth, farmerAuth.From)
	sim.Commit()
	Convey("Subject: Add farmer (role) to Contract\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldEqual, nil)
		})
	})
}

/*
 Check if farmer has role 'Farmer'
*/
func Test_UserHasFarmerRole(t *testing.T) {
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
	Convey("Subject: Inspector Address has Farmer role?\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("Inspector does not have 'Farmer' role", func() {
			So(isNotFarmer, ShouldEqual, false)
		})
	})
}

/*
 Remove role 'Farmer' from farmer
*/
func Test_RemoveFarmerRole(t *testing.T) {
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
		Convey("Farmer does not have role 'Farmer'", func() {
			So(isNotFarmer, ShouldEqual, false)
		})
	})
}

/*
 Add Role 'Admin' to admin
*/
func Test_AddAdminRole(t *testing.T) {
	beego.Trace("Test: ", "Add admin role with authorized user")
	_, err := rbacContract.AddAdmin(systemAuth, adminAuth.From)
	sim.Commit()
	Convey("Subject: Add user to admin roles\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldEqual, nil)
		})
	})
}

/*
 Check if admin has role 'Admin'
*/
func Test_UserHasAdminRole(t *testing.T) {
	isAdmin, err := rbacContract.IsAdmin(nil, adminAuth.From)
	sim.Commit()
	Convey("Subject: Address has Admin role?\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("Admin has role 'Admin'", func() {
			So(isAdmin, ShouldEqual, true)
		})
	})

	isNotAdmin, err := rbacContract.IsFarmer(nil, inspectorAuth.From)
	sim.Commit()
	Convey("Subject: Address has Admin role?\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldEqual, nil)
		})

		Convey("Inspector does not have 'Admin' role", func() {
			So(isNotAdmin, ShouldEqual, false)
		})
	})
}

/*
 Remove role 'Admin' from admin
*/
func Test_RemoveAdminRole(t *testing.T) {
	_, err := rbacContract.RemoveAdmin(systemAuth, adminAuth.From)
	sim.Commit()
	Convey("Subject: Remove farmer\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldEqual, nil)
		})
	})

	isNotAdmin, err := rbacContract.IsAdmin(nil, adminAuth.From)
	sim.Commit()
	Convey("Subject: Address has Farmer role?\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("AdminUser does not have role 'Admin'", func() {
			So(isNotAdmin, ShouldEqual, false)
		})
	})
}

/*
 Add Role 'Canton' to canton
*/
func Test_AddCantonRole(t *testing.T) {
	beego.Trace("Test: ", "Add canton role with authorized user")
	_, err := rbacContract.AddCantonEmployee(systemAuth, cantonAuth.From)
	sim.Commit()
	Convey("Subject: Add user to canton roles\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldEqual, nil)
		})
	})
}

/*
 Check if canton has role 'Canton'
*/
func Test_UserHasCantonRole(t *testing.T) {
	isCanton, err := rbacContract.IsCantonEmployee(nil, cantonAuth.From)
	sim.Commit()
	Convey("Subject: Address has Canton role?\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("Canton has role 'Canton'", func() {
			So(isCanton, ShouldEqual, true)
		})
	})

	isNotCanton, err := rbacContract.IsCantonEmployee(nil, farmerAuth.From)
	sim.Commit()
	Convey("Subject: Address has Canton role?\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("Farmer does not have role 'Canton'", func() {
			So(isNotCanton, ShouldEqual, false)
		})
	})
}

/*
 Remove role 'Canton' from canton
*/
func Test_RemoveCantonRole(t *testing.T) {
	_, err := rbacContract.RemovecantonEmployee(systemAuth, cantonAuth.From)
	sim.Commit()
	Convey("Subject: Remove farmer\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldEqual, nil)
		})
	})

	isNotCanton, err := rbacContract.IsCantonEmployee(nil, cantonAuth.From)
	sim.Commit()
	Convey("Subject: Address has canton role?\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("CantonUser does not have role 'Canton'", func() {
			So(isNotCanton, ShouldEqual, false)
		})
	})
}

/*
 Add Role 'Inspector' to inspector
*/
func Test_AddInspectorRole(t *testing.T) {
	beego.Trace("Test: ", "Add inspector role with authorized user")
	_, err := rbacContract.AddInspector(systemAuth, inspectorAuth.From)
	sim.Commit()
	Convey("Subject: Add user to inspector roles\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldEqual, nil)
		})
	})
}

/*
 Check if inspector has role 'Inspector'
*/
func Test_UserHasInspectorRole(t *testing.T) {
	isInspector, err := rbacContract.IsInspector(nil, inspectorAuth.From)
	sim.Commit()
	Convey("Subject: Address has Inspector role?\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("Inspector has role 'Inspector'", func() {
			So(isInspector, ShouldEqual, true)
		})
	})

	isNotInspector, err := rbacContract.IsInspector(nil, farmerAuth.From)
	sim.Commit()
	Convey("Subject: Address has Inspector role?\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("Farmer does not have role 'Inspector'", func() {
			So(isNotInspector, ShouldEqual, false)
		})
	})
}

/*
 Remove role 'Inspector' from inspector
*/
func Test_RemoveInspectorRole(t *testing.T) {
	_, err := rbacContract.RemoveInspector(systemAuth, inspectorAuth.From)
	sim.Commit()
	Convey("Subject: Remove farmer\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldEqual, nil)
		})
	})

	isNotInspector, err := rbacContract.IsInspector(nil, inspectorAuth.From)
	sim.Commit()
	Convey("Subject: Address has inspector role?\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("InspectorUser does not have role 'Inspector'", func() {
			So(isNotInspector, ShouldEqual, false)
		})
	})
}
