package request

import (
	"github.com/astaxie/beego"
	"testing"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"math/big"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/scmo/apayment-backend/smart-contracts/rbac"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/scmo/apayment-backend/smart-contracts/direct-payment-request"
)

var farmerAuth *bind.TransactOpts
var inspectorAuth *bind.TransactOpts
var adminAuth *bind.TransactOpts
var systemAuth *bind.TransactOpts
var sim *backends.SimulatedBackend

var rbacAddress common.Address
var requestContract *directpaymentrequest.RequestContract

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
	ra, _, rbacContract, _ := rbac.DeployRBACContract(systemAuth, sim)
	sim.Commit()
	rbacAddress = ra
	rbacContract.AddFarmer(systemAuth, farmerAuth.From)
	rbacContract.AddAdmin(systemAuth, adminAuth.From)
	rbacContract.AddInspector(systemAuth, inspectorAuth.From)
	sim.Commit()
}

func TestDeployContract(t *testing.T) {
	contributionCodes := []uint16{5416}
	remark := "This is my remark"
	beego.Trace("Test: ", "DeployContractRequest", "ContributionCodes: ", contributionCodes)
	_, _, rc, err := directpaymentrequest.DeployRequestContract(farmerAuth, sim, 3, contributionCodes, remark, rbacAddress)
	sim.Commit()
	requestContract = rc
	Convey("Subject: Deploy Request-Contract\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("Check remark", func() {
			contract_remark, _ := rc.Remark(nil)
			So(remark, ShouldEqual, contract_remark)
		})
	})

}

func TestSetInspector(t *testing.T) {
	beego.Trace("Test: ", "SetInspectorId", "InspectorId: ", inspectorAuth.From.String())
	_, err := requestContract.SetInspectorId(adminAuth, inspectorAuth.From)
	sim.Commit()
	Convey("Subject: Set Inspector Id\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("Check Inspector Address", func() {
			contract_inspectorId, _ := requestContract.InspectorAddress(nil)
			So(inspectorAuth.From, ShouldEqual, contract_inspectorId)
		})
	})
}

func TestAddLacks(t *testing.T) {
	beego.Trace("Test: ", "AddLacks", "ContributionCode: ", 5416)
	_, err := requestContract.AddLack(farmerAuth, 5416, "cat", 1110, "controlPoint", 1, 10)
	sim.Commit()
	Convey("Subject: Add Inspection Lack with wrong User\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldEqual, nil)
		})
		// Only the Inspector is allow to add a Lack
		Convey("Number of Lacks should be null", func() {
			numLacks, _ := requestContract.NumLacks(nil)
			So(numLacks.Cmp(big.NewInt(0)), ShouldEqual, 0) //  Cmp: 0 if x == y
		})
	})

	_, err = requestContract.AddLack(inspectorAuth, 5416, "cat", 1110, "controlPoint", 1, 10)
	sim.Commit()
	Convey("Subject: Add Inspection Lack as an Inspector\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldEqual, nil)
		})
		// Only the Inspector is allow to add a Lack
		Convey("Number of Lacks should be null", func() {
			numLacks, _ := requestContract.NumLacks(nil)
			So(numLacks.Cmp(big.NewInt(0)), ShouldEqual, 1) // +1 if x >  y
		})
	})
}

func Test_UpdateBTSGVE(t *testing.T) {
	_, err := requestContract.SetGVE(adminAuth, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	sim.Commit()
	Convey("Subject: Add Inspection Lack with wrong User\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("GVE should be 1", func() {
			gve, _ := requestContract.GetBTSGVE(nil, 1110)
			So(gve, ShouldEqual, 1)
		})
		Convey("GVE should be 2", func() {
			gve, _ := requestContract.GetBTSGVE(nil, 1150)
			So(gve, ShouldEqual, 2)
		})
	})
}

func TestCalculateBTS(t *testing.T) {
	beego.Trace("Test: ", "Calculate BTS", 5416)

	// Add some more lacks
	requestContract.AddLack(inspectorAuth, 5416, "cat", 1128, "controlPoint", 2, 110)
	requestContract.AddLack(inspectorAuth, 5416, "dog", 1141, "controlPoint", 1, 20)
	requestContract.AddLack(inspectorAuth, 5416, "dog", 1141, "controlPoint", 2, 60)
	sim.Commit()

	_, err := requestContract.CalculateBTS(adminAuth)
	if (err != nil) {
		beego.Error("Error while calculating BTS")
		t.Failed()
	}
	sim.Commit()

	Convey("Subject: Calculate BTS Direct Payment Amount\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("1128 gets no directpayment", func() {
			amount1128, _ := requestContract.PointGroups(nil, 1128)
			So(amount1128.BtsTotal, ShouldEqual, 0)
		})
		Convey("1141 should get 297", func() {
			amount1141, _ := requestContract.PointGroups(nil, 1141)
			So((amount1141.BtsTotal - amount1141.BtsDeduction) / 100, ShouldEqual, 297)
		})
	})
}


