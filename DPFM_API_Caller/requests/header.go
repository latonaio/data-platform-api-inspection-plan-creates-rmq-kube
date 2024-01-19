package requests

type Header struct {
	InspectionPlan                 int     `json:"InspectionPlan"`
	InspectionPlantBusinessPartner int     `json:"InspectionPlantBusinessPartner"`
	InspectionPlant                string  `json:"InspectionPlant"`
	Product                        *string `json:"Product"`
	ValidityStartDate			   *string `json:"ValidityStartDate"`
	ValidityEndDate				   *string `json:"ValidityEndDate"`
	ProductSpecification           *string `json:"ProductSpecification"`
	InspectionSpecification        *string `json:"InspectionSpecification"`
	InspectionPlanHeaderText       *string `json:"InspectionPlanHeaderText"`
	CertificateAuthorityChain      *string  `json:"CertificateAuthorityChain"`
	UsageControlChain        	   *string  `json:"UsageControlChain"`
	CreationDate                   string   `json:"CreationDate"`
	LastChangeDate                 string   `json:"LastChangeDate"`
	IsMarkedForDeletion            *bool    `json:"IsMarkedForDeletion"`
}
