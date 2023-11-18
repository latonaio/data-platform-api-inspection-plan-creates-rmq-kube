package dpfm_api_processing_formatter

type HeaderUpdates struct {
	InspectionPlan           int     `json:"InspectionPlan"`
	InspectionPlanHeaderText *string `json:"InspectionPlanHeaderText"`
}

type InspectionUpdates struct {
	InspectionPlan               int     `json:"InspectionPlan"`
	Inspection                   int     `json:"Inspection"`
	InspectionPlanInspectionText *string `json:"InspectionPlanInspectionText"`
}

type OperationUpdates struct {
	InspectionPlan                           int      `json:"InspectionPlan"`
	Operations                               int      `json:"Operations"`
	OperationsItem                           int      `json:"OperationsItem"`
}
