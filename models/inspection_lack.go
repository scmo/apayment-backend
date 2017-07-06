package models

type InspectionLack struct {
	ContributionCode  uint16 `json:"contributionCode"`
	ControlCategoryId int64  `json:"controlCategoryId"`
	PointGroupCode    uint16 `json:"pointGroupCode"`
	ControlPointId    int64  `json:"controlPointId"`
	LackId            int64  `json:"lackId"`
	Points            uint8  `json:"points"`
}
