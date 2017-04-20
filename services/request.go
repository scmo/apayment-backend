package services

import (
	"github.com/scmo/foodchain-backend/models"
	"github.com/astaxie/beego"
	"github.com/scmo/foodchain-backend/smart-contracts/request"
	"github.com/scmo/foodchain-backend/ethereum"
	"github.com/astaxie/beego/orm"
)

func CreateRequest(r *models.Request) error {

	ethereumController := ethereum.GetEthereumController()
	address, tx, _, err := smartcontracts.DeployRequestContract(ethereumController.Auth, ethereumController.Client)
	if err != nil {
		beego.Critical("Failed to deploy new token contract: ", err)
	}
	beego.Info("Contract pending deploy: ", address.String())
	beego.Info("Transaction waiting to be mined: ", tx.Hash().String())

	r.Address = address.String()

	o := orm.NewOrm()
	_, err = o.Insert(r)

	m2m := o.QueryM2M(r, "Contributions")
	for _, contributionPtr := range r.Contributions {
		_, err := m2m.Add(contributionPtr)
		if err != nil {
			beego.Error("Many2Many Add ", err.Error())
			return err
		}
	}
	return err
}

func GetAllRequests() []*models.Request {
	o := orm.NewOrm()
	var requests []*models.Request

	o.QueryTable(new(models.Request)).All(&requests)

	return requests
}

/*
func GetAllContributions() ([]*models.Contribution) {
	o := orm.NewOrm()
	var contributions []*models.Contribution
	o.QueryTable(new(models.Contribution)).All(&contributions)
	for _, contribution := range contributions {
		o.LoadRelated(contribution, "InspectionCriteria")
	}
	return contributions
}
 */