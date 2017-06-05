package test

import (
	"testing"
	_ "github.com/scmo/apayment-backend/routers"
	"github.com/astaxie/beego"

	"path/filepath"
	"runtime"
	"github.com/scmo/apayment-backend/services/tvd"
	"encoding/xml"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// Test soap request
func TestSoap(t *testing.T) {
	auth := tvd.BasicAuth{
		Login: "2000319",
		Password:"qikCloud2017",
	}
	animalTracingPortType := tvd.NewAnimalTracingPortType("https://ws-in.wbf.admin.ch/Livestock/AnimalTracing/1", true, &auth)

	versionRequest := tvd.Version{
		P_ManufacturerKey:"bebc4e6a-2477-4ec6-8837-d503a87e85f2",
	}
	versionResponse, err := animalTracingPortType.Version(&versionRequest)

	beego.Debug(versionResponse.VersionResult, err)
}

type MyRespEnvelope struct {
	XMLName xml.Name
	Body    Body
}

type Body struct {
	XMLName     xml.Name
	GetResponse completeResponse
}

type completeResponse struct {
	XMLName xml.Name `xml:"VersionResponse"`
	MyVar   string   `xml:"VersionResult,omitempty"`
}

func Test_ParsingSoap(t *testing.T) {

	Soap := []byte(`<?xml version="1.0" encoding="UTF-8"?><s:Envelope xmlns:a="http://www.w3.org/2005/08/addressing" xmlns:s="http://www.w3.org/2003/05/soap-envelope"><s:Header><a:Action s:mustUnderstand="1">http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1/AnimalTracingPortType/VersionResponse</a:Action></s:Header><s:Body><VersionResponse xmlns="http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1"><VersionResult>1.15.0</VersionResult></VersionResponse></s:Body></s:Envelope>`)
	res := &MyRespEnvelope{}
	err := xml.Unmarshal(Soap, res)

	beego.Debug(res.Body.GetResponse.MyVar, err)
}