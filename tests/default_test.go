package test

import (
	"runtime"
	"path/filepath"
	_ "github.com/scmo/foodchain-backend/routers"
	"github.com/astaxie/beego"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/scmo/foodchain-backend/services"
	"github.com/scmo/foodchain-backend/tests/db_test"
	"github.com/scmo/foodchain-backend/models"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)

	db_test.Setup()
}

func TestAddLackToContributions(t *testing.T) {
	contributions := make([] *models.Contribution, 0)
	iLack := models.InspectionLack{ContributionCode:5416, ControlCategoryId:"12.01_2017", PointGroupId:"A1", ControlPointId:"01", LackId:1}
	contribution := services.GetContributionByInspectionLack(&iLack)
	contributions = services.AddContributionToContributions(contributions, contribution)

	iLack = models.InspectionLack{ContributionCode:5416, ControlCategoryId:"12.01_2017", PointGroupId:"A2", ControlPointId:"01", LackId:7}
	contribution = services.GetContributionByInspectionLack(&iLack)
	contributions2 := services.AddContributionToContributions(contributions, contribution)

	iLack = models.InspectionLack{ContributionCode:5416, ControlCategoryId:"12.01_2017", PointGroupId:"A1", ControlPointId:"02", LackId:4}
	contribution = services.GetContributionByInspectionLack(&iLack)
	contributions3 := services.AddContributionToContributions(contributions, contribution)

	iLack = models.InspectionLack{ContributionCode:5416, ControlCategoryId:"12.01_2017", PointGroupId:"A1", ControlPointId:"01", LackId:2}
	contribution = services.GetContributionByInspectionLack(&iLack)
	contributions4 := services.AddContributionToContributions(contributions, contribution)

	iLack = models.InspectionLack{ContributionCode:5417, ControlCategoryId:"12.01_2017", PointGroupId:"A1", ControlPointId:"01", LackId:28}
	contribution = services.GetContributionByInspectionLack(&iLack)
	contributions5 := services.AddContributionToContributions(contributions, contribution)

	Convey("Subject: Test Generating Contributions with Lacks \n", t, func() {
		Convey("Length should 1", func() {
			So(len(contributions), ShouldEqual, 1)
		})
		Convey("Length should still be 1", func() {
			So(len(contributions2), ShouldEqual, 1)
		})
		Convey("ControlCategory should have two PointGroups", func() {
			So(len(contributions2[0].ControlCategories[0].PointGroups), ShouldEqual, 2)
		})
		Convey("PointGroup[0] should have two ControlPoints", func() {
			So(len(contributions3[0].ControlCategories[0].PointGroups[0].ControlPoints), ShouldEqual, 2)
		})
		Convey("ControlPoint[0] should have two Lacks", func() {
			So(len(contributions4[0].ControlCategories[0].PointGroups[0].ControlPoints[0].Lacks), ShouldEqual, 2)
		})
		Convey("Length should be 2", func() {
			So(len(contributions5), ShouldEqual, 2)
		})
	})
}


