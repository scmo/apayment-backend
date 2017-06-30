package models

type InspectionLack struct {
	ContributionCode  uint16 `json:"contributionCode"`
	ControlCategoryId string `json:"controlCategoryId"`
	PointGroupCode    uint16 `json:"pointGroupCode"`
	ControlPointId    string `json:"controlPointId"`
	LackId            int64  `json:"lackId"`
	Points            uint8  `json:"points"`
}
