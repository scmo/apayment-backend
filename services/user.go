package services

import (
	"github.com/scmo/foodchain-backend/models"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
	"errors"
	"github.com/scmo/foodchain-backend/ethereum"
	"net/http"
	"context"
	"github.com/ethereum/go-ethereum/common"
)

func CreateUser(u *models.User) error {
	o := orm.NewOrm()
	hash, err := hashPassword(u.Password)
	if err != nil {
		beego.Error("HashPassword ", err.Error())
		return err
	}
	u.Password = hash

	if (beego.BConfig.RunMode != "test") {
		accountAddress, err := createNewEthereumAccount()
		if err != nil {
			beego.Error("Creating Ethereum Account ", err.Error())
			return err
		}
		u.Address = accountAddress
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
		if err != nil {
			beego.Error("Many2Many Add ", err.Error())
			return err
		}
	}
	return err
}

func createNewEthereumAccount() (string, error) {
	ethereumController := ethereum.GetEthereumController()
	account, err := ethereumController.Keystore.NewAccount(beego.AppConfig.String("userAccountPassword"))
	if beego.BConfig.RunMode == "dev" {
		ethereum.SendEther("0x88f1e48e11864bfc4685f9f9d8b79b18450764ef", account.Address.String(), 1)
		//addEthers(account.Address.String())
	}

	return account.Address.String(), err
}

func addEthers(address string) {
	faucetURL := "http://faucet.ropsten.be:3001/donate/" + address
	beego.Info(faucetURL)
	resp, err := http.Get(faucetURL)
	if err != nil {
		beego.Error("Error while calling faucet. ", err)
	}
	beego.Info("Calling faucet. Response Code: ", resp.Status)
}

func CheckLogin(_username string, _password string) (models.User, error) {
	o := orm.NewOrm()

	user := models.User{Username: _username}
	err := o.Read(&user, "Username")
	if checkPasswordHash(_password, user.Password) == false {
		return user, errors.New("Wrong password")
	}
	o.LoadRelated(&user, "Roles")
	return user, err;
}

func GetAllUsers() ([]*models.User) {
	o := orm.NewOrm()
	var users []*models.User
	o.QueryTable(new(models.User)).All(&users)
	for _, user := range users {
		o.LoadRelated(user, "Roles")
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
	setEtherBalance(&user)
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

	setEtherBalance(&user)
	return &user, nil
}

func setEtherBalance(user *models.User) {
	ethereumController := ethereum.GetEthereumController()
	ctx := context.Background()
	latestBlock, err := ethereumController.Client.BlockByNumber(ctx, nil)
	if err != nil {
		beego.Critical("Failed to get latest block: ", err)
	}
	balance, err := ethereumController.Client.BalanceAt(ctx, common.HexToAddress(user.Address), latestBlock.Number())
	beego.Debug(balance)
	user.EthereumBalance = balance
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