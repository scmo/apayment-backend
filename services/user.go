package services

import (
	"context"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/scmo/apayment-backend/ethereum"
	"github.com/scmo/apayment-backend/models"
	"github.com/scmo/apayment-backend/services/tvd"
	"github.com/scmo/apayment-backend/smart-contracts/rbac"
	"golang.org/x/crypto/bcrypt"
	"math/big"
)

func CreateUser(u *models.User) error {
	o := orm.NewOrm()
	hash, err := hashPassword(u.Password)
	if err != nil {
		beego.Error("HashPassword ", err.Error())
		return err
	}
	u.Password = hash

	if beego.BConfig.RunMode != "test" {
		accountAddress, err := createNewEthereumAccount()
		if err != nil {
			beego.Error("Creating Ethereum Account ", err.Error())
			return err
		}
		u.EtherumAddress = accountAddress
	}

	_, err = o.Insert(u)
	if err != nil {
		beego.Error("Inser User ", err.Error())
		return err
	}
	m2m := o.QueryM2M(u, "Roles")
	for _, rolePtr := range u.Roles {
		if _, id, err := o.ReadOrCreate(rolePtr, "Name"); err == nil {
			rolePtr.Id = id
		} else {
			beego.Error("ReadOrCreate ", err.Error())
			return err
		}
		_, err := m2m.Add(rolePtr)
		addUserToRBAC(u.EtherumAddress, rolePtr.Name)
		if err != nil {
			beego.Error("Many2Many Add ", err.Error())
			return err
		}
	}
	return err
}
func addUserToRBAC(address string, role string) {
	ethereumController := ethereum.GetEthereumController()
	rbacContract, err := rbac.NewRBACContract(common.HexToAddress(beego.AppConfig.String("accessControlContract")), ethereumController.Client)
	if err != nil {
		beego.Error("Error while creating a new instace of RBAC")
	}
	tx := new(types.Transaction)
	switch role {
	case "Farmer":
		tx, err = rbacContract.AddFarmer(ethereumController.Auth, common.HexToAddress(address))
	case "Inspector":
		tx, err = rbacContract.AddInspector(ethereumController.Auth, common.HexToAddress(address))
	case "Admin":
		tx, err = rbacContract.AddAdmin(ethereumController.Auth, common.HexToAddress(address))
	case "Canton":
		tx, err = rbacContract.AddCantonEmployee(ethereumController.Auth, common.HexToAddress(address))
	default:
		beego.Error("Unknown role - address was not added to the RoleBasedAccessControl")
	}
	if err != nil {
		beego.Critical("Error while adding User to RBAC.", err)
	}
	beego.Info("Transaction waiting to be mined: ", tx.Hash().String())

}

func createNewEthereumAccount() (string, error) {
	ethereumController := ethereum.GetEthereumController()
	account, err := ethereumController.Keystore.NewAccount(beego.AppConfig.String("userAccountPassword"))
	// TODO: Uncomment for production
	//if beego.BConfig.RunMode == "dev" {
	amountEther := 0.5
	amount := new(big.Float).Mul(big.NewFloat(amountEther), big.NewFloat(params.Ether))
	amountWei := new(big.Int)
	amount.Int(amountWei)
	ethereum.SendWei(beego.AppConfig.String("systemAccountAddress"), account.Address.String(), amountWei)
	//}
	return account.Address.String(), err
}

func CheckLoginWithUsername(_username string, _password string) (models.User, error) {
	o := orm.NewOrm()

	user := models.User{Username: _username}
	err := o.Read(&user, "Username")
	if checkPasswordHash(_password, user.Password) == false {
		return user, errors.New("Wrong password")
	}
	o.LoadRelated(&user, "Roles")
	return user, err
}

func CheckLoginWithEmail(_email string, _password string) (models.User, error) {
	o := orm.NewOrm()

	user := models.User{Email: _email}
	err := o.Read(&user, "Email")
	if checkPasswordHash(_password, user.Password) == false {
		return user, errors.New("Wrong password")
	}
	o.LoadRelated(&user, "Roles")
	return user, err
}

func GetAllUsers() []*models.User {
	o := orm.NewOrm()
	var users []*models.User
	o.QueryTable(new(models.User)).All(&users)
	for _, user := range users {
		o.LoadRelated(user, "Roles")
		setTVD(user)
		setEtherBalance(user)
		setAPaymentTokenBalance(user)
	}
	return users
}

func GetAllUsersByRole(_role string) ([]*models.User, error) {
	o := orm.NewOrm()
	role := models.Role{Name: _role}
	err := o.Read(&role, "Name")

	if err == orm.ErrNoRows {
		beego.Error("No result found.")
		return nil, err
	} else if err == orm.ErrMissPK {
		beego.Error("No primary key found.")
		return nil, err
	}
	o.LoadRelated(&role, "Users")
	return role.Users, nil
}

func GetUserById(_id int64) (*models.User, error) {
	o := orm.NewOrm()
	user := models.User{Id: _id}
	err := o.Read(&user)
	if err == orm.ErrNoRows {
		beego.Error("No result found.")
		return nil, err
	} else if err == orm.ErrMissPK {
		beego.Error("No primary key found.")
		return nil, err
	}
	o.LoadRelated(&user, "Roles")
	setTVD(&user)
	setEtherBalance(&user)
	setAPaymentTokenBalance(&user)
	return &user, nil
}

func GetUserByUsername(_username string) (*models.User, error) {
	o := orm.NewOrm()
	user := models.User{Username: _username}
	err := o.Read(&user, "Username")
	if err == orm.ErrNoRows {
		beego.Error("No result found.")
		return nil, err
	} else if err == orm.ErrMissPK {
		beego.Error("No primary key found.")
		return nil, err
	}
	o.LoadRelated(&user, "Roles")
	setTVD(&user)
	setEtherBalance(&user)
	setAPaymentTokenBalance(&user)
	return &user, nil
}

func GetUserByAddress(etherumAddress string) (*models.User, error) {
	o := orm.NewOrm()
	user := models.User{EtherumAddress: etherumAddress}
	err := o.Read(&user, "etherum_address")
	if err == orm.ErrNoRows {
		beego.Error("No result found.")
		return nil, err
	} else if err == orm.ErrMissPK {
		beego.Error("No primary key found.")
		return nil, err
	}
	o.LoadRelated(&user, "Roles")
	setTVD(&user)
	setEtherBalance(&user)
	setAPaymentTokenBalance(&user)
	return &user, nil
}

func setTVD(user *models.User) {
	personAddressResult, err := tvd.GetPersonAddressFromTVD()
	if err != nil {
		beego.Error("Error while fetching PersonAddress from TVD. ", err)
	}
	user.PersonAddressResult = personAddressResult

	return
	if user.TVD == 0 {
		return
	}
	animalHusbandryDetailResult, err := tvd.GetAnimalHusbandryDetailFromTVD(user.TVD)
	if err != nil {
		beego.Error("Error while fetching AnimalHusbandryDetail from TVD. ", err)
	}
	user.AnimalHusbandryDetailResult = animalHusbandryDetailResult
}

func setEtherBalance(user *models.User) {
	ethereumController := ethereum.GetEthereumController()
	ctx := context.Background()
	latestBlock, err := ethereumController.Client.BlockByNumber(ctx, nil)
	if err != nil {
		beego.Critical("Failed to get latest block: ", err)
	}
	balance, err := ethereumController.Client.BalanceAt(ctx, common.HexToAddress(user.EtherumAddress), latestBlock.Number())
	user.EthereumBalance = balance
}

func setAPaymentTokenBalance(user *models.User) {
	balance, err := GetBalanceOf(common.HexToAddress(user.EtherumAddress))
	if err != nil {
		beego.Error("Error while getting balance", err)
	}
	user.APaymentTokenBalance = balance
}

func CountUsers() (int64, error) {
	o := orm.NewOrm()
	cnt, err := o.QueryTable(new(models.User)).Count() // SELECT COUNT(*) FROM USE
	return cnt, err
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
