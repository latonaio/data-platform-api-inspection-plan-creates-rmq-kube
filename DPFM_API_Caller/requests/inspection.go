package requests

type Inspection struct {
	InspectionPlan                           int      `json:"InspectionPlan"`
	Inspection                               int      `json:"Inspection"`
	InspectionType                           string   `json:"InspectionType"`
	InspectionTypeValueUnit                  *string  `json:"InspectionTypeValueUnit"`
	InspectionTypePlannedValue               *float32 `json:"InspectionTypePlannedValue"`
	InspectionTypeCertificateType            *string  `json:"InspectionTypeCertificateType"`
	InspectionTypeCertificateValueInText     *string  `json:"InspectionTypeCertificateValueInText"`
	InspectionTypeCertificateValueInQuantity *float32 `json:"InspectionTypeCertificateValueInQuantity"`
	InspectionPlanInspectionText             *string  `json:"InspectionPlanInspectionText"`
	CreationDate                             string   `json:"CreationDate"`
	LastChangeDate                           string   `json:"LastChangeDate"`
	IsMarkedForDeletion                      *int     `json:"IsMarkedForDeletion"`
}
