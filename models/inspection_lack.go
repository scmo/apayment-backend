package models

type InspectionLack struct {
	ContributionCode  uint16 `json:"contributionCode"`
	ControlCategoryId string `json:"controlCategoryId"`
	PointGroupId      string `json:"pointGroupId"`
	ControlPointId    string `json:"controlPointId"`
	LackId            int64  `json:"lackId"`
}
