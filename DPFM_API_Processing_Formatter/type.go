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
