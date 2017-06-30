package models

type Inspection struct {
	RequestId   int64             `json:"requestId"`
	InspectorId int64             `json:"inspectorId"`
	Lacks       []*InspectionLack `json:"lacks"`
}
