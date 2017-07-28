package direct_payment_token

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/scmo/apayment-backend/smart-contracts/apayment-token"
	. "github.com/smartystreets/goconvey/convey"
	"math/big"
	"testing"
)

var farmerAuth *bind.TransactOpts
var cantonAuth *bind.TransactOpts
var systemAuth *bind.TransactOpts
var sim *backends.SimulatedBackend

var tokenContract *apaymenttoken.APaymentTokenContract

/*
 Prepare Simulated Blockchain inclusive accounts
*/
func init() {
	// Generate a new random account and a funded simulator
	key, _ := crypto.GenerateKey()
	farmerAuth = bind.NewKeyedTransactor(key)

	key, _ = crypto.GenerateKey()
	cantonAuth = bind.NewKeyedTransactor(key)

	key, _ = crypto.GenerateKey()
	systemAuth = bind.NewKeyedTransactor(key)

	sim = backends.NewSimulatedBackend(core.GenesisAlloc{
		farmerAuth.From: {Balance: big.NewInt(10000000000)},
		systemAuth.From: {Balance: big.NewInt(10000000000)},
		cantonAuth.From: {Balance: big.NewInt(10000000000)},
	})
	sim.Commit()
}

/*
 Deploy aPayment contract
*/
func Test_DeployContract(t *testing.T) {
	Convey("Subject: Deploy aPaymentToken-Contract\n", t, func() {
		supply := big.NewInt(1000000000000)
		_, _, tc, err := apaymenttoken.DeployAPaymentTokenContract(systemAuth, sim, supply)
		sim.Commit()
		tokenContract = tc
		Convey("No error", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("SystemAccount should have inital supply as balance", func() {
			balance, _ := tc.BalanceOf(nil, systemAuth.From)
			So(balance.Cmp(supply), ShouldEqual, 0) // -1 if x <  y, 0 if x == y, +1 if x >  y
		})
	})
}

/*
 Transfer token from system to canton
*/
func Test_TranferTokenToCanton(t *testing.T) {
	Convey("Subject: Transfer 100 Token from System to Canton\n", t, func() {
		amount := big.NewInt(100)
		tokenContract.Transfer(systemAuth, cantonAuth.From, amount)
		sim.Commit()
		Convey("Farmer should have 100 Token", func() {
			balance, _ := tokenContract.BalanceOf(nil, cantonAuth.From)
			So(balance.Cmp(amount), ShouldEqual, 0) // -1 if x <  y, 0 if x == y, +1 if x >  y
		})
	})
}

/*
 Transfer token from system to canton
*/
func Test_TranferTokenToSystem(t *testing.T) {
	Convey("Subject: Transfer 10 Token from Canton to Farmer\n", t, func() {
		amount := big.NewInt(10)
		tokenContract.Transfer(systemAuth, farmerAuth.From, amount)
		sim.Commit()
		Convey("Farmer should have 10 Token", func() {
			balance, _ := tokenContract.BalanceOf(nil, farmerAuth.From)
			So(balance.Cmp(amount), ShouldEqual, 0) // -1 if x <  y, 0 if x == y, +1 if x >  y
		})
	})
}
