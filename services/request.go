package services

import (
	"github.com/scmo/apayment-backend/models"
	"github.com/astaxie/beego"
	"github.com/scmo/apayment-backend/ethereum"
	"github.com/astaxie/beego/orm"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"github.com/scmo/apayment-backend/smart-contracts/direct-payment-request"
	"github.com/scmo/apayment-backend/services/tvd"
)

func CreateRequest(request *models.Request, auth *bind.TransactOpts) error {
	ethereumController := ethereum.GetEthereumController()
	// TODO: Uncomment this when TVD is again up
	//gvesMap, err := tvd.GetNumberOfGVE(request.User.TVD, beego.AppConfig.String("agate_username"), beego.AppConfig.String("agate_password"))
	//if ( err != nil ) {
	//	beego.Error("Failed to get GVE. ", err)
	//	return err
	//}
	//var gvesList [9]uint16
	//for  i, value := range gvesMap {
	//	gvesList[i] = value
	//}
	gvesList := [9]uint16{29, 0, 18, 21, 0, 0, 5, 9, 1}

	address, tx, _, err := directpaymentrequest.DeployRequestContract(auth, ethereumController.Client, getContributionCodes(request), request.Remark, gvesList)
	if err != nil {
		beego.Error("Failed to deploy new token contract: ", err)
		return err
	}
	beego.Info("Contract pending deploy: ", address.String())
	beego.Info("Transaction waiting to be mined: ", tx.Hash().String())

	request.Address = address.String()

	o := orm.NewOrm()
	_, err = o.Insert(request)
	if err != nil {
		beego.Error("Failed to insert new Request: ", err)
	}
	return err
}

func GetAllRequests() []*models.Request {
	o := orm.NewOrm()
	var requests []*models.Request
	o.QueryTable(new(models.Request)).RelatedSel().All(&requests)
	for _, request := range requests {
		requestContract, err := getRequestContractByAddress(request.Address)
		if err != nil {
			beego.Error("Failed to instantiate a Token contract: %v", err)
		}
		assignRequest(request, requestContract, false)
	}
	return requests
}

func GetAllRequestsByUserId(userId int64) []*models.Request {
	o := orm.NewOrm()
	var requests []*models.Request
	o.QueryTable(new(models.Request)).Filter("user", userId).RelatedSel().All(&requests)
	for _, request := range requests {
		requestContract, err := getRequestContractByAddress(request.Address)
		if err != nil {
			beego.Error("Failed to instantiate a Token contract: %v", err)
		}
		assignRequest(request, requestContract, false)
	}
	return requests
}

func GetRequestById(requestId int64, smartContract bool) *models.Request {
	o := orm.NewOrm()
	var request models.Request
	err := o.QueryTable(new(models.Request)).Filter("Id", requestId).RelatedSel().One(&request)
	if err != nil {
		beego.Error("Error while fetching Request.", err)
	}
	if ( smartContract) {
		requestContract, err := getRequestContractByAddress(request.Address)
		if err != nil {
			beego.Error("Failed to instantiate a Token contract: %v", err)
		}
		assignRequest(&request, requestContract, true)
	}

	return &request
}

func GetRequestAddressById(requestId int64) string {
	o := orm.NewOrm()
	var request models.Request
	err := o.QueryTable(new(models.Request)).Filter("Id", requestId).RelatedSel().One(&request)
	if err != nil {
		beego.Error("Error while fetching Request.", err)
	}
	return request.Address
}

func GetRequestIdByAddress(requestAddress string) int64 {
	o := orm.NewOrm()
	var request models.Request
	err := o.QueryTable(new(models.Request)).Filter("Address", requestAddress).RelatedSel().One(&request)
	if err != nil {
		beego.Error("Error while fetching Request.", err)
	}
	return request.Id
}

func GetAllRequestsForInspection() []*models.Request {
	o := orm.NewOrm()
	var requests []*models.Request
	o.QueryTable(new(models.Request)).Filter("inspector__isnull", false).RelatedSel().All(&requests)
	for _, request := range requests {
		requestContract, err := getRequestContractByAddress(request.Address)
		if err != nil {
			beego.Error("Failed to instantiate a Token contract: %v", err)
		}
		assignRequest(request, requestContract, false)
	}
	return requests
}

func GetAllRequestsForInspectionByInspectorId(inspectorId int64) []*models.Request {
	o := orm.NewOrm()
	var requests []*models.Request
	o.QueryTable(new(models.Request)).Filter("inspector", inspectorId).RelatedSel().All(&requests)
	for _, request := range requests {
		requestContract, err := getRequestContractByAddress(request.Address)
		if err != nil {
			beego.Error("Failed to instantiate a Token contract: %v", err)
		}
		assignRequest(request, requestContract, false)
	}
	return requests
}

func AddInspectorToRequest(request *models.Request, auth *bind.TransactOpts) error {
	// Add to DB
	o := orm.NewOrm()
	o.Update(request, "Inspector")

	// Add to the SmartContract
	requestContract, err := getRequestContractByAddress(request.Address)
	if err != nil {
		beego.Error("Error while fetching RequestContract by Address: ", err)
		return err
	}
	session := getRequestContractSession(requestContract)
	session.TransactOpts.From = auth.From
	session.TransactOpts.Signer = auth.Signer

	tx, err := session.SetInspectorId(common.HexToAddress(request.Inspector.Address))
	if err != nil {
		beego.Error("Failed to update request (add inspector): ", err)
		return err
	}
	beego.Info("Transaction waiting to be mined: ", tx.Hash().String())
	return err
}

// Add inspection Lacks to Request
func AddLacksToRequest(inspection *models.Inspection, auth *bind.TransactOpts) error {
	requestAddress := GetRequestAddressById(inspection.RequestId)
	// Add to the SmartContract
	requestContract, err := getRequestContractByAddress(requestAddress)

	if err != nil {
		beego.Error("Error while fetching RequestContract by Address: ", err)
		return err
	}
	session := getRequestContractSession(requestContract)
	session.TransactOpts.From = auth.From
	session.TransactOpts.Signer = auth.Signer
	for _, lack := range inspection.Lacks {
		beego.Debug("Add Lack", lack.LackId, lack.Points, lack.PointGroupCode)
		tx, err := session.AddLack(lack.ContributionCode, lack.ControlCategoryId, lack.PointGroupCode, lack.ControlPointId, lack.LackId, lack.Points)
		if err != nil {
			beego.Critical("Failed to update name: ", err)
			return err
		}
		beego.Info("Transaction waiting to be mined: ", tx.Hash().String())
	}
	return err
}

func SetGVE(request *models.Request) (error) {
	ethereumController := ethereum.GetEthereumController()
	gves, err := tvd.GetNumberOfGVE(request.User.TVD, beego.AppConfig.String("agate_username"), beego.AppConfig.String("agate_password"))
	if ( err != nil ) {
		beego.Error("Failed to get GVE. ", err)
		return err
	}
	requestContract, err := getRequestContractByAddress(request.Address)
	if err != nil {
		beego.Error("Error while fetching RequestContract by Address: ", err)
		return err
	}
	session := getRequestContractSession(requestContract)
	session.TransactOpts.From = ethereumController.Auth.From
	session.TransactOpts.Signer = ethereumController.Auth.Signer

	tx, err := session.SetGVE(gves[1110], gves[1150], gves[1128], gves[1141], gves[1142], gves[1124], gves[1129], gves[1143], gves[1144])
	if (err != nil) {
		beego.Critical("Failed to update GVE: ", err)
		return err
	}
	beego.Info("Transaction waiting to be mined: ", tx.Hash().String())

	setGVE(request, session)

	return err
}

//func AddPayment(request *models.Request, from common.Address, amount *big.Int) (error) {
//	ethereumController := ethereum.GetEthereumController()
//	requestContract, err := getRequestByAddress(request.Address)
//	if err != nil {
//		beego.Error("Error while fetching RequestContract by Address: ", err)
//		return err
//	}
//	session := getRequestContractSession(requestContract)
//	session.TransactOpts.From = ethereumController.Auth.From
//	session.TransactOpts.Signer = ethereumController.Auth.Signer
//
//	tx, err := session.AddPayment(from, amount)
//	beego.Info("Transaction waiting to be mined: ", tx.Hash().String())
//	return err
//}

func GetFirstPaymentAmount(request *models.Request) (*big.Int, error) {
	ethereumController := ethereum.GetEthereumController()

	requestContract, err := getRequestContractByAddress(request.Address)
	if err != nil {
		beego.Error("Error while fetching RequestContract by Address: ", err)
		return nil, err
	}
	session := getRequestContractSession(requestContract)
	session.TransactOpts.From = ethereumController.Auth.From
	session.TransactOpts.Signer = ethereumController.Auth.Signer

	amount, err := session.GetFirstPaymentAmount()
	if err != nil {
		beego.Error("Error while first payment amount: ", err)
		return nil, err
	}
	return amount, err
}

func getRequestContractSession(requestContract *directpaymentrequest.RequestContract) (*directpaymentrequest.RequestContractSession) {
	//ethereumController := ethereum.GetEthereumController()
	requestContractSesssion := &directpaymentrequest.RequestContractSession{
		Contract:requestContract,
		CallOpts:bind.CallOpts{Pending:true},
		//TransactOpts:bind.TransactOpts{
		//	From: ethereumController.Auth.From,
		//	Signer: ethereumController.Auth.Signer,
		//	GasLimit: big.NewInt(3141592),
		//},
	}
	return requestContractSesssion

}

func getContributionCodes(request *models.Request) []uint16 {
	var codes = make([]uint16, len(request.Contributions))
	for i, contribution := range request.Contributions {
		codes[i] = contribution.Code
	}
	return codes
}

func getRequestContractByAddress(address string) (*directpaymentrequest.RequestContract, error) {
	ethereumController := ethereum.GetEthereumController()
	return directpaymentrequest.NewRequestContract(common.HexToAddress(address), ethereumController.Client)
}

func assignRequest(request *models.Request, requestContract *directpaymentrequest.RequestContract, full bool) {
	session := getRequestContractSession(requestContract)
	remark, err := session.Remark()
	if err != nil {
		beego.Error("Failed to instantiate a Token contract: ", err)
	}
	request.Remark = remark
	setInspector(request, session)
	setTimestamps(request, session)
	if (full) {
		setContributions(request, session)
		setLacksInspected(request, session)
		setGVE(request, session)
		setPayments(request)
	}
}

func setInspector(request *models.Request, session *directpaymentrequest.RequestContractSession) {
	if (request.Inspector != nil) {
		inspectorAddress, err := session.InspectorAddress()
		if err != nil {
			beego.Error("Error while reading InspectorId from Contract: ", err)
		}
		if (inspectorAddress.String() != request.Inspector.Address) {
			beego.Info("Blockchain: ", inspectorAddress.String())
			beego.Info("DB: ", request.Inspector.Address)
			beego.Error("Request Inspector Id and Contract InspectorId does not match!")
		}
	}
}

func setContributions(request *models.Request, session *directpaymentrequest.RequestContractSession) {
	request.Contributions = make([]*models.Contribution, 0)
	next := true
	index := big.NewInt(0)
	for next {
		code, err := session.ContributionCodes(index)
		if err != nil {
			next = false
		}
		if err == nil {
			contribution, err := GetContributionByCode(code)
			if err != nil {
				beego.Error("Error getting Contribution", err)
			}
			request.Contributions = append(request.Contributions, contribution)
		}
		index.Add(index, big.NewInt(1))
	}

}

func setLacksInspected(request *models.Request, session *directpaymentrequest.RequestContractSession) {
	contributions := make([] *models.Contribution, 0)
	numLack, err := session.NumLacks()
	if err != nil {
		beego.Error("Error getting NumLacks", err)
		return
	}
	for i := new(big.Int).Set(big.NewInt(0)); i.Cmp(numLack) < 0; i.Add(i, big.NewInt(1)) {
		lack, err := session.Lacks(i)
		if err != nil {
			beego.Error("Error getting Lack", err)
			return
		}

		inspectionLack := models.InspectionLack(lack)
		contribution := GetContributionByInspectionLack(&inspectionLack)
		contributions = AddContributionToContributions(contributions, contribution)
	}

	request.ContributionsWithLacks = contributions
}

func setPayments(request *models.Request) {
	request.Payments = GetTransactionForRequest(request.Address)
}

func setTimestamps(request *models.Request, session *directpaymentrequest.RequestContractSession) {

	// Created
	createdTimestamp, err := session.Created()
	if err != nil {
		beego.Error("Error while reading createdTimestamp from Contract: ", err)
	}
	request.Created = createdTimestamp

	// Modified
	modifiedTimestamp, err := session.Modified()
	if err != nil {
		beego.Error("Error while reading updatedTimestamp from Contract: ", err)
	}
	request.Modified = modifiedTimestamp
}

// Function fetches GVE values from the smart contract and adds it to the 'request' model
func setGVE(request *models.Request, session *directpaymentrequest.RequestContractSession) {
	pointGroupCodes := tvd.GetPointGroupCodes()

	for i := range pointGroupCodes {
		requestGVE, err := session.PointGroups(pointGroupCodes[i])
		if ( err != err ) {
			beego.Critical("Error get GVE", err)
		}
		pointGroup, err := GetAllPointGroupByCode(pointGroupCodes[i])
		gve := models.GVE{PointGroup: pointGroup, Amount: requestGVE.Gve}
		request.GVE = append(request.GVE, &gve)
	}
}
// Function adds a Contribution to a slice of contributions if
// contribution does not exist yet in the slice, if contribution exist,
// it checks if the contribution alreayd contains the ControlCategory and so on...
func AddContributionToContributions(contributions []*models.Contribution, contribution *models.Contribution) []*models.Contribution {
	for i := range contributions {
		if contributions[i].Id == contribution.Id {
			// Contribution exist! Check if ControlCategory exists
			for j := range contributions[i].ControlCategories {
				if contributions[i].Id == contribution.Id &&
					contributions[i].ControlCategories[j].Id == contribution.ControlCategories[0].Id {
					// ControlCategory exists! Check if PointGroup exist
					for k := range contributions[i].ControlCategories[j].PointGroups {
						if contributions[i].Id == contribution.Id &&
							contributions[i].ControlCategories[j].Id == contribution.ControlCategories[0].Id &&
							contributions[i].ControlCategories[j].PointGroups[k].Id == contribution.ControlCategories[0].PointGroups[0].Id {
							// PointGroup exists! Check if ControlPoint exist
							for l := range contributions[i].ControlCategories[j].PointGroups[k].ControlPoints {
								if contributions[i].Id == contribution.Id &&
									contributions[i].ControlCategories[j].Id == contribution.ControlCategories[0].Id &&
									contributions[i].ControlCategories[j].PointGroups[k].Id == contribution.ControlCategories[0].PointGroups[0].Id &&
									contributions[i].ControlCategories[j].PointGroups[k].ControlPoints[l].Id == contribution.ControlCategories[0].PointGroups[0].ControlPoints[0].Id {
									// ControlPoint exists! Check if Lack exist
									for m := range contributions[i].ControlCategories[j].PointGroups[k].ControlPoints[l].Lacks {
										if contributions[i].Id == contribution.Id &&
											contributions[i].ControlCategories[j].Id == contribution.ControlCategories[0].Id &&
											contributions[i].ControlCategories[j].PointGroups[k].Id == contribution.ControlCategories[0].PointGroups[0].Id &&
											contributions[i].ControlCategories[j].PointGroups[k].ControlPoints[l].Id == contribution.ControlCategories[0].PointGroups[0].ControlPoints[0].Id &&
											contributions[i].ControlCategories[j].PointGroups[k].ControlPoints[l].Lacks[m].Id == contribution.ControlCategories[0].PointGroups[0].ControlPoints[0].Lacks[0].Id {
											// Lack exists! Check if Lack exist
											return contributions
										}
									}
									contributions[i].ControlCategories[j].PointGroups[k].ControlPoints[l].Lacks = append(contributions[i].ControlCategories[j].PointGroups[k].ControlPoints[l].Lacks, contribution.ControlCategories[0].PointGroups[0].ControlPoints[0].Lacks[0])
									return contributions
								}
							}
							contributions[i].ControlCategories[j].PointGroups[k].ControlPoints = append(contributions[i].ControlCategories[j].PointGroups[k].ControlPoints, contribution.ControlCategories[0].PointGroups[0].ControlPoints[0])
							return contributions
						}
					}
					contributions[i].ControlCategories[j].PointGroups = append(contributions[i].ControlCategories[j].PointGroups, contribution.ControlCategories[0].PointGroups[0])
					return contributions
				}
			}
			contributions[i].ControlCategories = append(contributions[i].ControlCategories, contribution.ControlCategories[0])
			return contributions
		}
	}
	return append(contributions, contribution)
}

