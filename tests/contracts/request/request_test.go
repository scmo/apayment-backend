package request

import (
	"github.com/astaxie/beego"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/scmo/apayment-backend/smart-contracts/direct-payment-request"
	"github.com/scmo/apayment-backend/smart-contracts/rbac"
	. "github.com/smartystreets/goconvey/convey"
	"math/big"
	"testing"
)

var farmerAuth *bind.TransactOpts
var inspectorAuth *bind.TransactOpts
var fakeInspectorAuth *bind.TransactOpts
var adminAuth *bind.TransactOpts
var cantonAuth *bind.TransactOpts
var systemAuth *bind.TransactOpts
var sim *backends.SimulatedBackend

var rbacAddress common.Address
var requestContract *directpaymentrequest.RequestContract
var requestContract2 *directpaymentrequest.RequestContract

var amountPreviousYear *big.Int

/*
 Prepare accounts and simulated backend
 Deploy RBAC contract and add roles to address
*/
func init() {
	// Generate a new random account and a funded simulator
	key, _ := crypto.GenerateKey()
	farmerAuth = bind.NewKeyedTransactor(key)

	key, _ = crypto.GenerateKey()
	adminAuth = bind.NewKeyedTransactor(key)

	key, _ = crypto.GenerateKey()
	inspectorAuth = bind.NewKeyedTransactor(key)

	key, _ = crypto.GenerateKey()
	fakeInspectorAuth = bind.NewKeyedTransactor(key)

	key, _ = crypto.GenerateKey()
	cantonAuth = bind.NewKeyedTransactor(key)

	key, _ = crypto.GenerateKey()
	systemAuth = bind.NewKeyedTransactor(key)

	sim = backends.NewSimulatedBackend(core.GenesisAlloc{
		farmerAuth.From:        {Balance: big.NewInt(10000000000)},
		adminAuth.From:         {Balance: big.NewInt(10000000000)},
		systemAuth.From:        {Balance: big.NewInt(10000000000)},
		inspectorAuth.From:     {Balance: big.NewInt(10000000000)},
		cantonAuth.From:        {Balance: big.NewInt(10000000000)},
		fakeInspectorAuth.From: {Balance: big.NewInt(10000000000)},
	})
	ra, _, rbacContract, _ := rbac.DeployRBACContract(systemAuth, sim)
	sim.Commit()
	rbacAddress = ra
	rbacContract.AddFarmer(systemAuth, farmerAuth.From)
	rbacContract.AddCantonEmployee(systemAuth, cantonAuth.From)
	rbacContract.AddAdmin(systemAuth, adminAuth.From)
	rbacContract.AddInspector(systemAuth, inspectorAuth.From)
	rbacContract.AddInspector(systemAuth, fakeInspectorAuth.From)
	sim.Commit()

	amountPreviousYear = big.NewInt(500000)
}

/*
 Deploy DirectPaymentRequest
*/
func Test_DeployContract(t *testing.T) {
	contributionCodes := []uint16{5416}
	remark := "This is my remark"
	beego.Trace("Test: ", "DeployContractRequest", "ContributionCodes: ", contributionCodes)
	gvesList := []uint32{290000, 0, 180000, 210000, 0, 0, 50000, 90000, 10000}

	_, _, rc, err := directpaymentrequest.DeployRequestContract(farmerAuth, sim, contributionCodes, remark, rbacAddress, gvesList, amountPreviousYear)
	sim.Commit()
	requestContract = rc
	Convey("Subject: Deploy Request-Contract\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldBeNil)
		})
		Convey("Check remark", func() {
			contract_remark, _ := rc.Remark(nil)
			So(contract_remark, ShouldContainSubstring, remark)
		})
	})

}

/*
 Check if GVE is set correctly
*/
func Test_GVE(t *testing.T) {
	Convey("Subject: Check if GVE is set correctly\n", t, func() {
		Convey("Pointgroup 1110: GVE is 29", func() {
			cv, _ := requestContract.PointGroups(nil, 1110)
			So(cv.Gve, ShouldEqual, 290000)
		})
		Convey("Pointgroup 1150: GVE is 0", func() {
			cv, _ := requestContract.PointGroups(nil, 1150)
			So(cv.Gve, ShouldEqual, 0)
		})
		Convey("Pointgroup 1128: GVE is 18", func() {
			cv, _ := requestContract.PointGroups(nil, 1128)
			So(cv.Gve, ShouldEqual, 180000)
		})
		Convey("Pointgroup 1141: GVE is 21", func() {
			cv, _ := requestContract.PointGroups(nil, 1141)
			So(cv.Gve, ShouldEqual, 210000)
		})
		Convey("Pointgroup 1142: GVE is 0", func() {
			cv, _ := requestContract.PointGroups(nil, 1142)
			So(cv.Gve, ShouldEqual, 0)
		})
		Convey("Pointgroup 1124: GVE is 0", func() {
			cv, _ := requestContract.PointGroups(nil, 1124)
			So(cv.Gve, ShouldEqual, 0)
		})
		Convey("Pointgroup 1129: GVE is 5", func() {
			cv, _ := requestContract.PointGroups(nil, 1129)
			So(cv.Gve, ShouldEqual, 50000)
		})
		Convey("Pointgroup 1143: GVE is 9", func() {
			cv, _ := requestContract.PointGroups(nil, 1143)
			So(cv.Gve, ShouldEqual, 90000)
		})
		Convey("Pointgroup 1144: GVE is 1", func() {
			cv, _ := requestContract.PointGroups(nil, 1144)
			So(cv.Gve, ShouldEqual, 10000)
		})
	})
}

/*
 User from canton sets inspector
*/
func Test_SetInspector(t *testing.T) {
	beego.Trace("Test: ", "SetInspector ", "Inspector address: ", inspectorAuth.From.String())
	_, err := requestContract.SetInspectorId(cantonAuth, inspectorAuth.From)
	sim.Commit()
	Convey("Subject: Set Inspector Id\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldBeNil)
		})
		Convey("Check Inspector Address", func() {
			contract_inspectorId, _ := requestContract.InspectorAddress(nil)
			So(inspectorAuth.From, ShouldEqual, contract_inspectorId)
		})
	})
}

/*
 Inspect sets lacks
*/
func Test_AddLacks(t *testing.T) {
	cCodes := []uint16{5416, 5416, 5416, 5416, 5416, 5417, 5417, 5417, 5417, 5417}
	cCatIds := []int64{1, 1, 1, 1, 1, 2, 2, 2, 2, 2}
	pgCodes := []uint16{1110, 1110, 1128, 1129, 1144, 1110, 1110, 1128, 1129, 1144}
	cPoiIds := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	lackIds := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	points := []uint8{40, 20, 111, 0, 80, 40, 20, 111, 0, 80}

	_, err := requestContract.AddLacks(fakeInspectorAuth, cCodes, cCatIds, pgCodes, cPoiIds, lackIds, points)
	sim.Commit()
	Convey("Subject: Add Inspection Lack with fake Inspector\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldBeNil)
		})
		// Only the assigned Inspector is allow to add a Lack
		Convey("Number of Lacks should be null", func() {
			numLacks, _ := requestContract.NumLacks(nil)
			So(numLacks.Cmp(big.NewInt(0)), ShouldEqual, 0) //  Cmp: 0 if x == y
		})
	})

	_, err = requestContract.AddLacks(inspectorAuth, cCodes, cCatIds, pgCodes, cPoiIds, lackIds, points)

	sim.Commit()
	Convey("Subject: Add Inspection Lack as an Inspector\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldBeNil)
		})
		Convey("Number of Lacks should be null", func() {
			numLacks, _ := requestContract.NumLacks(nil)
			So(numLacks.Cmp(big.NewInt(0)), ShouldEqual, 1) // +1 if x >  y
		})
	})
}

/*
 Test if 'minus' points are stored
*/
func Test_PointGroupCalculation(t *testing.T) {
	Convey("Subject: Check if point are added to PointGroupCalculation struct\n", t, func() {
		Convey("Pointgroup 1110: BTS Point is 60", func() {
			pgc, _ := requestContract.PointGroups(nil, 1110)
			So(pgc.BtsPoints, ShouldEqual, 60)
		})
		Convey("Pointgroup 1128: BTS Point is 111", func() {
			pgc, _ := requestContract.PointGroups(nil, 1128)
			So(pgc.BtsPoints, ShouldEqual, 111)
		})
		Convey("Pointgroup 1144: BTS Point is 80", func() {
			pgc, _ := requestContract.PointGroups(nil, 1144)
			So(pgc.BtsPoints, ShouldEqual, 80)
		})
		Convey("Pointgroup 1110: RAUS Point is 60", func() {
			pgc, _ := requestContract.PointGroups(nil, 1110)
			So(pgc.RausPoints, ShouldEqual, 60)
		})
		Convey("Pointgroup 1128: RAUS Point is 111", func() {
			pgc, _ := requestContract.PointGroups(nil, 1128)
			So(pgc.RausPoints, ShouldEqual, 111)
		})
		Convey("Pointgroup 1144: RAUS Point is 80", func() {
			pgc, _ := requestContract.PointGroups(nil, 1144)
			So(pgc.RausPoints, ShouldEqual, 80)
		})
	})
}

/*
 Test if BTS Amount and deduction has been calculated correctly
*/
func Test_CalculateBTS(t *testing.T) {
	Convey("Subject: Check if total amount and deduction is calculated correctly for BTS\n", t, func() {
		Convey("Pointgroup 1110: Amount is 261000", func() {
			cv, _ := requestContract.PointGroups(nil, 1110)
			So(cv.BtsTotal.Cmp(big.NewInt(2610000000)), ShouldEqual, 0)
		})
		Convey("Pointgroup 1110: Deduction is 4500", func() {
			cv, _ := requestContract.PointGroups(nil, 1110)
			beego.Debug(cv)
			So(cv.BtsDeduction.Cmp(big.NewInt(45000000)), ShouldEqual, 0)
		})
		Convey("Pointgroup 1128: Amount is 162000", func() {
			cv, _ := requestContract.PointGroups(nil, 1128)
			So(cv.BtsTotal.Cmp(big.NewInt(1620000000)), ShouldEqual, 0)
		})
		Convey("Pointgroup 1128: Deduction is 162000", func() {
			cv, _ := requestContract.PointGroups(nil, 1128)
			So(cv.BtsDeduction.Cmp(big.NewInt(1620000000)), ShouldEqual, 0)
		})
		Convey("Pointgroup 1129: Amount is 45000", func() {
			cv, _ := requestContract.PointGroups(nil, 1129)
			So(cv.BtsTotal.Cmp(big.NewInt(450000000)), ShouldEqual, 0)
		})
		Convey("Pointgroup 1129: Deduction is 0", func() {
			cv, _ := requestContract.PointGroups(nil, 1129)
			So(cv.BtsDeduction.Cmp(big.NewInt(0)), ShouldEqual, 0)
		})
		Convey("Pointgroup 1144: Amount is 0, (not qualified)", func() {
			cv, _ := requestContract.PointGroups(nil, 1144)
			So(cv.BtsTotal.Cmp(big.NewInt(0)), ShouldEqual, 0)
		})
		Convey("Pointgroup 1144: Deduction is 0", func() {
			cv, _ := requestContract.PointGroups(nil, 1144)
			So(cv.BtsDeduction.Cmp(big.NewInt(0)), ShouldEqual, 0)
		})
	})
}

/*
 Test if RAUS Amount and deduction has been calculated correctly
*/
func Test_CalculateRAUS(t *testing.T) {
	Convey("Subject: Check if total amount and deduction is calculated correctly for RAUS\n", t, func() {
		Convey("Pointgroup 1110: Amount is 551000", func() {
			cv, _ := requestContract.PointGroups(nil, 1110)
			So(cv.RausTotal.Cmp(big.NewInt(5510000000)), ShouldEqual, 0)
		})
		Convey("Pointgroup 1110: Deduction is 9500", func() {
			cv, _ := requestContract.PointGroups(nil, 1110)
			So(cv.RausDeduction.Cmp(big.NewInt(95000000)), ShouldEqual, 0)
		})
		Convey("Pointgroup 1128: Amount is 342000", func() {
			cv, _ := requestContract.PointGroups(nil, 1128)
			So(cv.RausTotal.Cmp(big.NewInt(3420000000)), ShouldEqual, 0)
		})
		Convey("Pointgroup 1128: Deduction is 342000", func() {
			cv, _ := requestContract.PointGroups(nil, 1128)
			So(cv.RausDeduction.Cmp(big.NewInt(3420000000)), ShouldEqual, 0)
		})
		Convey("Pointgroup 1129: Amount is 95000", func() {
			cv, _ := requestContract.PointGroups(nil, 1129)
			So(cv.RausTotal.Cmp(big.NewInt(950000000)), ShouldEqual, 0)
		})
		Convey("Pointgroup 1129: Deduction is 0", func() {
			cv, _ := requestContract.PointGroups(nil, 1129)
			So(cv.RausDeduction.Cmp(big.NewInt(0)), ShouldEqual, 0)
		})
		Convey("Pointgroup 1144: Amount is 37000", func() {
			cv, _ := requestContract.PointGroups(nil, 1144)
			So(cv.RausTotal.Cmp(big.NewInt(370000000)), ShouldEqual, 0)
		})
		Convey("Pointgroup 1144: Deduction is 25900", func() {
			cv, _ := requestContract.PointGroups(nil, 1144)
			So(cv.RausDeduction.Cmp(big.NewInt(259000000)), ShouldEqual, 0)
		})
	})
}

/*
 Test if amout of first payout is correctly
*/
func Test_FirstPayoutAmount(t *testing.T) {
	payoutAmount := big.NewInt(0)
	payoutAmount.Div(amountPreviousYear, big.NewInt(2))
	Convey("Subject: Get amount for first payment\n", t, func() {
		amount, err := requestContract.GetFirstPaymentAmount(nil)
		Convey("No error", func() {
			So(err, ShouldBeNil)
		})
		Convey("First payment should be "+payoutAmount.String(), func() {
			So(amount.Cmp(payoutAmount), ShouldEqual, 0)
		})

	})
}

/*
 Test if amout of final payout is correctly
*/
func Test_FinalPayoutAmount(t *testing.T) {
	Convey("Subject: Get amount for final payment\n", t, func() {
		amount, err := requestContract.GetFinalPaymentAmount(nil)
		Convey("No error", func() {
			So(err, ShouldBeNil)
		})
		Convey("First payment should be ", func() {
			So(amount.Cmp(big.NewInt(1539100)), ShouldEqual, 0)
		})

	})
}


/*
 Deploy DirectPaymentRequest
*/
func Test_DeployContract2(t *testing.T) {
	contributionCodes := []uint16{5416, 5417}
	remark := "Real one"
	beego.Trace("Test: ", "DeployContractRequest", "ContributionCodes: ", contributionCodes)
	gvesList := []uint32{0, 0, 0, 0, 0, 0, 47748, 389934, 0}

	_, _, rc, err := directpaymentrequest.DeployRequestContract(farmerAuth, sim, contributionCodes, remark, rbacAddress, gvesList, big.NewInt(0))
	sim.Commit()
	requestContract2 = rc
	Convey("Subject: Deploy Request-Contract\n", t, func() {
		Convey("No error", func() {
			So(err, ShouldBeNil)
		})
		Convey("Check remark", func() {
			contract_remark, _ := rc.Remark(nil)
			So(contract_remark, ShouldContainSubstring, remark)
		})
	})
}


/*
 Test if BTS Amount and deduction has been calculated correctly
*/
func Test_CalculateBTS2(t *testing.T) {
	Convey("Subject: Check if total amount and deduction is calculated correctly for BTS\n", t, func() {
		Convey("Pointgroup 1129: Amount is 45000", func() {
			cv, _ := requestContract2.PointGroups(nil, 1129)
			So(cv.BtsTotal.Cmp(big.NewInt(429732000)), ShouldEqual, 0)
		})
		Convey("Pointgroup 1129: Deduction is 0", func() {
			cv, _ := requestContract2.PointGroups(nil, 1129)
			So(cv.BtsDeduction.Cmp(big.NewInt(0)), ShouldEqual, 0)
		})
		Convey("Pointgroup 1143: Amount is 0", func() {
			cv, _ := requestContract2.PointGroups(nil, 1143)
			So(cv.BtsTotal.Cmp(big.NewInt(3509406000)), ShouldEqual, 0)
		})
		Convey("Pointgroup 1143: Deduction is 0", func() {
			cv, _ := requestContract2.PointGroups(nil, 1143)
			So(cv.BtsDeduction.Cmp(big.NewInt(0)), ShouldEqual, 0)
		})
	})
}

/*
 Test if RAUS Amount and deduction has been calculated correctly
*/
func Test_CalculateRAUS2(t *testing.T) {
	Convey("Subject: Check if total amount and deduction is calculated correctly for RAUS\n", t, func() {
		Convey("Pointgroup 1129: Amount is 907212000", func() {
			cv, _ := requestContract2.PointGroups(nil, 1129)
			beego.Debug(cv)
			So(cv.RausTotal.Cmp(big.NewInt(907212000)), ShouldEqual, 0)
		})
		Convey("Pointgroup 1129: Deduction is 0", func() {
			cv, _ := requestContract2.PointGroups(nil, 1129)
			So(cv.RausDeduction.Cmp(big.NewInt(0)), ShouldEqual, 0)
		})
		Convey("Pointgroup 1143: Amount is 7408746000", func() {
			cv, _ := requestContract2.PointGroups(nil, 1143)
			So(cv.RausTotal.Cmp(big.NewInt(7408746000)), ShouldEqual, 0)
		})
		Convey("Pointgroup 1143: Deduction is 0", func() {
			cv, _ := requestContract2.PointGroups(nil, 1143)
			So(cv.RausDeduction.Cmp(big.NewInt(0)), ShouldEqual, 0)
		})
	})
}

/*
 Test if amout of first payout is correctly
*/
func Test_FirstPayoutAmount2(t *testing.T) {
	payoutAmount := big.NewInt(0)
	Convey("Subject: Get amount for first payment\n", t, func() {
		amount, err := requestContract2.GetFirstPaymentAmount(nil)
		Convey("No error", func() {
			So(err, ShouldBeNil)
		})
		Convey("First payment should be "+payoutAmount.String(), func() {
			So(amount.Cmp(payoutAmount), ShouldEqual, 0)
		})
	})
}

/*
 Test if amout of final payout is correctly
*/
func Test_FinalPayoutAmount2(t *testing.T) {
	Convey("Subject: Get amount for final payment\n", t, func() {
		amount, err := requestContract2.GetFinalPaymentAmount(nil)
		Convey("No error", func() {
			So(err, ShouldBeNil)
		})
		Convey("Second payment should be ", func() {
			beego.Debug(amount)
			So(amount.Cmp(big.NewInt(1225509	)), ShouldEqual, 0)
		})
	})
}