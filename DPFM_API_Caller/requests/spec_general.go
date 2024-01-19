package requests

type SpecGeneral struct {
	InspectionPlan	    	int	     `json:"InspectionPlan"`
    HeatNumber	        	string   `json:"HeatNumber"`
	CreationDate            string   `json:"CreationDate"`
	LastChangeDate          string   `json:"LastChangeDate"`
	IsMarkedForDeletion     *bool    `json:"IsMarkedForDeletion"`
}
