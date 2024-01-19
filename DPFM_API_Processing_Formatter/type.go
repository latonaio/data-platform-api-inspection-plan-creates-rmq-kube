package dpfm_api_processing_formatter

type HeaderUpdates struct {
	InspectionPlan           	int     `json:"InspectionPlan"`
	InspectionSpecification		*string	`json:"InspectionSpecification"`
	ValidityStartDate			*string	`json:"ValidityStartDate"`
	ValidityEndDate				*string	`json:"ValidityEndDate"`
	InspectionPlanHeaderText 	*string `json:"InspectionPlanHeaderText"`
}

type SpecGeneralUpdates struct {
	InspectionPlan	int    `json:"InspectionPlan"`
	HeatNumber		string `json:"HeatNumber"`
}

type SpecDetailUpdates struct {
	InspectionPlan			int     `json:"InspectionPlan"`
    SpecType	            string	`json:"SpecType"`
    UpperLimitValue	        float32	`json:"UpperLimitValue"`
    LowerLimitValue	        float32	`json:"LowerLimitValue"`
    StandardValue	        float32	`json:"StandardValue"`
    SpecTypeUnit	        string	`json:"SpecTypeUnit"`
    Formula			        *string	`json:"Formula"`
}

type ComponentCompositionUpdates struct {
	InspectionPlan								int     `json:"InspectionPlan"`
	ComponentCompositionType					string	`json:"ComponentCompositionType"`
	ComponentCompositionUpperLimitInPercent		float32	`json:"ComponentCompositionUpperLimitInPercent"`
	ComponentCompositionLowerLimitInPercent		float32	`json:"ComponentCompositionLowerLimitInPercent"`
	ComponentCompositionStandardValueInPercent	float32	`json:"ComponentCompositionStandardValueInPercent"`
}

type InspectionUpdates struct {
	InspectionPlan               				int      `json:"InspectionPlan"`
	Inspection                   				int      `json:"Inspection"`
    InspectionType                            	string	 `json:"InspectionType"`
    InspectionTypeValueUnit	                    *string	 `json:"InspectionTypeValueUnit"`
    InspectionTypePlannedValue	                *float32 `json:"InspectionTypePlannedValue"`
    InspectionTypeCertificateType	            *string	 `json:"InspectionTypeCertificateType"`
    InspectionTypeCertificateValueInText	    *string	 `json:"InspectionTypeCertificateValueInText"`
    InspectionTypeCertificateValueInQuantity	*float32 `json:"InspectionTypeCertificateValueInQuantity"`
    InspectionPlanInspectionText                *string	 `json:"InspectionPlanInspectionText"`
}

type OperationUpdates struct {
	InspectionPlan                           int      `json:"InspectionPlan"`
	Operations                               int      `json:"Operations"`
	OperationsItem                           int      `json:"OperationsItem"`
}
