package test

import (
	"github.com/astaxie/beego"
	_ "github.com/scmo/apayment-backend/routers"
	"testing"

	"github.com/scmo/apayment-backend/services/tvd"
	. "github.com/smartystreets/goconvey/convey"
	"path/filepath"
	"runtime"
	"os"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func Test_TvdVersion(t *testing.T) {
	auth := tvd.BasicAuth{
		Login:    os.Getenv("AGATE_USERNAME"),
		Password: os.Getenv("AGATE_PASSWORD"),
	}

	animalTracingPortType := tvd.NewAnimalTracingPortType("https://ws-in.wbf.admin.ch/Livestock/AnimalTracing/1", true, &auth)

	versionRequest := tvd.Version{
		PManufacturerKey: os.Getenv("TVD_MANUFACTURER_KEY"),
		Action:           "http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1/AnimalTracingPortType/Version",
	}


	versionResponse, err := animalTracingPortType.Version(&versionRequest)
	Convey("Subject: Test TVD Version\n", t, func() {
		Convey("Error should be nil", func() {
			So(err, ShouldBeNil)
		})
		Convey("Version should not be null (= '')", func() {
			So(len(versionResponse.VersionResult), ShouldBeGreaterThan, 0)
		})
	})
}

func Test_SetCategory(t *testing.T) {
	cowA1 := tvd.CattleLivestockDataV2{
		Gender:           "1",
		BirthDate:        "2016-04-05T00:00:00",
		FirstCalvingDate: "2017-02-05T00:00:00",
	}
	cowA3 := tvd.CattleLivestockDataV2{
		Gender:    "1",
		BirthDate: "2016-04-05T00:00:00",
	}
	cowA4 := tvd.CattleLivestockDataV2{
		Gender:    "1",
		BirthDate: "2016-12-26T00:00:00",
	}
	cowA5 := tvd.CattleLivestockDataV2{
		Gender:    "1",
		BirthDate: "2017-06-05T00:00:00",
	}
	cowA6 := tvd.CattleLivestockDataV2{
		Gender:    "2",
		BirthDate: "2004-10-05T00:00:00",
	}
	cowA7 := tvd.CattleLivestockDataV2{
		Gender:    "2",
		BirthDate: "2016-04-05T00:00:00",
	}
	cowA8 := tvd.CattleLivestockDataV2{
		Gender:    "2",
		BirthDate: "2016-12-26T00:00:00",
	}
	cowA9 := tvd.CattleLivestockDataV2{
		Gender:    "2",
		BirthDate: "2017-06-05T00:00:00",
	}

	Convey("Subject: Test Categorize Cow\n", t, func() {
		Convey("Cow 1 should be A1", func() {
			cat, _ := tvd.GetAnimalCategory(&cowA1)
			So(cat, ShouldEqual, 1)
		})
		Convey("Cow 3 should be A3", func() {
			cat, _ := tvd.GetAnimalCategory(&cowA3)
			So(cat, ShouldEqual, 3)
		})
		Convey("Cow 4 should be A4", func() {
			cat, _ := tvd.GetAnimalCategory(&cowA4)
			So(cat, ShouldEqual, 4)
		})
		Convey("Cow 5 should be A5", func() {
			cat, _ := tvd.GetAnimalCategory(&cowA5)
			So(cat, ShouldEqual, 5)
		})
		Convey("Cow 6 should be A6", func() {
			cat, _ := tvd.GetAnimalCategory(&cowA6)
			So(cat, ShouldEqual, 6)
		})
		Convey("Cow 7 should be A7", func() {
			cat, _ := tvd.GetAnimalCategory(&cowA7)
			So(cat, ShouldEqual, 7)
		})
		Convey("Cow 8 should be A8", func() {
			cat, _ := tvd.GetAnimalCategory(&cowA8)
			So(cat, ShouldEqual, 8)
		})
		Convey("Cow 9 should be A9", func() {
			cat, _ := tvd.GetAnimalCategory(&cowA9)
			So(cat, ShouldEqual, 9)
		})
	})
}

func Test_GetPersonAddress(t *testing.T) {

	auth := tvd.BasicAuth{
		Login:    os.Getenv("AGATE_USERNAME"),
		Password: os.Getenv("AGATE_PASSWORD"),
	}
	animalTracingPortType := tvd.NewAnimalTracingPortType("https://ws-in.wbf.admin.ch/Livestock/AnimalTracing/1", true, &auth)

	getPersonAddressRequest := tvd.GetPersonAddress{
		Action:           "http://www.admin.ch/xmlns/Services/evd/Livestock/AnimalTracing/1/AnimalTracingPortType/GetPersonAddress",
		PManufacturerKey: os.Getenv("TVD_MANUFACTURER_KEY"),
		PLCID:            2055,
		PAgateNumber:     os.Getenv("AGATE_USERNAME"),
	}

	resp, err := animalTracingPortType.GetPersonAddress(&getPersonAddressRequest)
	if err != nil {
		beego.Error("Failed to fetch info about the person: ", err)
	}
	beego.Debug(resp.GetPersonAddressResult.PostAddress)
	Convey("Subject: Retrieve correct Address from TVD\n", t, func() {
		Convey("Error should be nil", func() {
			So(err, ShouldBeNil)
		})
		Convey("Name is Meinrad", func() {
			So(resp.GetPersonAddressResult.PostAddress.FirstName, ShouldEqual, "Meinrad")
		})
		Convey("Street is Rüeggerstrasse 8", func() {
			So(resp.GetPersonAddressResult.PostAddress.Street, ShouldEqual, "Rüeggerstrasse 8")
		})
		Convey("Postcode is 9108", func() {
			So(resp.GetPersonAddressResult.PostAddress.PostCode, ShouldEqual, "9108")
		})
		Convey("City is Gonten", func() {
			So(resp.GetPersonAddressResult.PostAddress.City, ShouldEqual, "Gonten")
		})
	})
}
