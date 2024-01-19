package requests

type SpecDetail struct {
	InspectionPlan	        int	    `json:"InspectionPlan"`
    SpecType	            string	`json:"SpecType"`
    UpperLimitValue	        float32	`json:"UpperLimitValue"`
    LowerLimitValue	        float32	`json:"LowerLimitValue"`
    StandardValue	        float32	`json:"StandardValue"`
    SpecTypeUnit	        string	`json:"SpecTypeUnit"`
    Formula			        *string	`json:"Formula"`
	CreationDate            string   `json:"CreationDate"`
	LastChangeDate          string   `json:"LastChangeDate"`
	IsMarkedForDeletion     *bool    `json:"IsMarkedForDeletion"`
}
