package requests

type ComponentComposition struct {
	InspectionPlan								int		`json:"InspectionPlan"`
	ComponentCompositionType					string	`json:"ComponentCompositionType"`
	ComponentCompositionUpperLimitInPercent		float32	`json:"ComponentCompositionUpperLimitInPercent"`
	ComponentCompositionLowerLimitInPercent		float32	`json:"ComponentCompositionLowerLimitInPercent"`
	ComponentCompositionStandardValueInPercent	float32	`json:"ComponentCompositionStandardValueInPercent"`
	CreationDate            					string   `json:"CreationDate"`
	LastChangeDate          					string   `json:"LastChangeDate"`
	IsMarkedForDeletion     					*bool    `json:"IsMarkedForDeletion"`
}
