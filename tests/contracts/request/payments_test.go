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
	. "github.com/smartystreets/goconvey/convey"
	"github.com/scmo/apayment-backend/smart-contracts/direct-payment-request"
)

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

func Test_DeployContract(t *testing.T) {
	contributionCodes := []uint16{5416}
	remark := "This is my remark"
	beego.Trace("Test: ", "DeployContractRequest", "ContributionCodes: ", contributionCodes)
	//gvesList := [9]uint16{29,0,18,21,0,0,5,9,1}
	_, _, rc, err := directpaymentrequest.DeployRequestContract(farmerAuth, sim, contributionCodes, remark, rbacAddress)
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

	// SET GVE
	requestContract.SetGVE(farmerAuth, 0, 0, 0, 0, 0, 0, 0, 0, 0)
	sim.Commit();
}

func Test_CalcTotal(t *testing.T) {
	requestContract.SetGVE(farmerAuth, 29, 0, 18, 21, 0, 0, 5, 9, 1)
	sim.Commit();

	Convey("Subject: CalcVariable 1110\n", t, func() {
		cv, _ := requestContract.PointGroups(nil, 1110)
		Convey("GVE ist 29", func() {
			So(cv.Gve, ShouldEqual, 29)
		})
		Convey("BTS Total is 29 * 9000", func() {
			So(cv.BtsTotal.Cmp(big.NewInt(261000)), ShouldEqual, 0)
		})
	})
	Convey("Subject: CalcVariable 1128\n", t, func() {
		cv, _ := requestContract.PointGroups(nil, 1128)
		Convey("GVE ist 18", func() {
			So(cv.Gve, ShouldEqual, 18)
		})
		//Convey("BTS Total is 9 * 9000", func() {
		//	So(cv.BtsTotal, ShouldEqual, 261000)
		//})
	})
	Convey("Subject: CalcVariable 1141\n", t, func() {
		cv, _ := requestContract.PointGroups(nil, 1141)
		Convey("GVE ist 21", func() {
			So(cv.Gve, ShouldEqual, 21)
		})
		//Convey("BTS Total is 9 * 9000", func() {
		//	So(cv.BtsTotal, ShouldEqual, 261000)
		//})
	})
	Convey("Subject: CalcVariable 1143\n", t, func() {
		cv, _ := requestContract.PointGroups(nil, 1143)
		Convey("GVE ist 9", func() {
			So(cv.Gve, ShouldEqual, 9)
		})
		Convey("BTS Total is 9 * 9000", func() {
			So(cv.BtsTotal.Cmp(big.NewInt(81000)), ShouldEqual, 0)
		})
	})
}

func T_est_FirstPaymentAmount(t *testing.T) {
	Convey("Subject: Calculation of first payment amount with no GVE\n", t, func() {
		amount, _ := requestContract.GetFirstPaymentAmount(nil)
		Convey("Amount should be 0", func() {
			So(amount.Cmp(big.NewInt(0)), ShouldEqual, 0)
		})
	})
	// SET GVE
	requestContract.SetGVE(farmerAuth, 1, 0, 0, 0, 0, 0, 0, 0, 0)
	sim.Commit();
	Convey("Subject: Calculation of first payment amount with 1110 = 1\n", t, func() {
		amount, _ := requestContract.GetFirstPaymentAmount(nil)
		Convey("Amount should be 9000", func() {
			beego.Debug(amount)
			So(amount.Cmp(big.NewInt(9000)), ShouldEqual, 0)
		})
	})
	// SET GVE
	requestContract.SetGVE(farmerAuth, 1, 1, 1, 1, 1, 1, 1, 1, 1)
	sim.Commit();
	Convey("Subject: Calculation of first payment amount with every GVE cat = 1\n", t, func() {
		amount, _ := requestContract.GetFirstPaymentAmount(nil)
		Convey("Amount should be 9000", func() {
			So(amount.Cmp(big.NewInt(63000)), ShouldEqual, 0)
		})
	})



	//Convey("Subject: Calculation of first payment amount with every GVE cat = 1\n", t, func() {
	//	amount, _ := requestContract.GetFirstPaymentAmount(nil)
	//	beego.Debug(amount)
	//	//Convey("Amount should be 9000", func() {
	//	//	So(amount.Cmp(big.NewInt(63000)), ShouldEqual, 0)
	//	//})
	//})


}

