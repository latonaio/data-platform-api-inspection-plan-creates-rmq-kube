package dpfm_api_processing_formatter

import dpfm_api_input_reader "data-platform-api-inspection-plan-creates-rmq-kube/DPFM_API_Input_Reader"

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		InspectionPlan:           data.InspectionPlan,
		InspectionPlanHeaderText: data.InspectionPlanHeaderText,
	}
}

func ConvertToInspectionUpdates(inspectionUpdates dpfm_api_input_reader.Inspection) *InspectionUpdates {
	data := inspectionUpdates

	return &InspectionUpdates{
		InspectionPlan:               data.InspectionPlan,
		InspectionPlanInspectionText: data.InspectionPlanInspectionText,
	}
}

func ConvertToOperationUpdates(operationUpdates dpfm_api_input_reader.Operation) *OperationUpdates {
	data := operationUpdates

	return &OperationUpdates{
		InspectionPlan: data.InspectionPlan,
		Operations:     data.Operations,
		OperationsItem: data.OperationsItem,
	}
}
