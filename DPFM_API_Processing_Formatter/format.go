package dpfm_api_processing_formatter

import dpfm_api_input_reader "data-platform-api-inspection-plan-creates-rmq-kube/DPFM_API_Input_Reader"

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		InspectionPlan:           	data.InspectionPlan,
		InspectionSpecification:	data.InspectionSpecification,
		ValidityStartDate:			data.ValidityStartDate,
		ValidityEndDate:			data.ValidityEndDate,
		InspectionPlanHeaderText: 	data.InspectionPlanHeaderText,
	}
}

func ConvertToSpecGeneralUpdates(header dpfm_api_input_reader.SpecGeneral) *SpecGeneralUpdates {
	data := specGeneralUpdates

	return &SpecGeneralUpdates{
		InspectionPlan:		data.InspectionPlan,
		HeatNumber:			data.HeatNumber,
	}
}

func ConvertToComponentCompositionUpdates(header dpfm_api_input_reader.ComponentComposition) *ComponentCompositionUpdates {
	data := componentCompositionUpdates

	return &ComponentCompositionUpdates{
		InspectionPlan:								data.InspectionPlan,
		ComponentCompositionType:					data.ComponentCompositionType,
		ComponentCompositionUpperLimitInPercent:	data.ComponentCompositionUpperLimitInPercent,
		ComponentCompositionLowerLimitInPercent:	data.ComponentCompositionLowerLimitInPercent,
		ComponentCompositionStandardValueInPercent:	data.ComponentCompositionStandardValueInPercent,
	}
}

func ConvertToSpecDetailUpdates(header dpfm_api_input_reader.SpecDetail) *SpecDetailUpdates {
	data := specDetailUpdates

	return &SpecDetailUpdates{
		InspectionPlan:		data.InspectionPlan,
		SpecType:			data.SpecType,
		UpperLimitValue:	data.UpperLimitValue,
		LowerLimitValue:	data.LowerLimitValue,
		StandardValue:		data.StandardValue,
		SpecTypeUnit:		data.SpecTypeUnit,
		Formula:			data.Formula,
	}
}

func ConvertToInspectionUpdates(inspectionUpdates dpfm_api_input_reader.Inspection) *InspectionUpdates {
	data := inspectionUpdates

	return &InspectionUpdates{
		InspectionPlan:								data.InspectionPlan,
		Inspection:									data.InspectionPlanInspectionText,
		InspectionType:								data.InspectionType,
		InspectionTypeValueUnit:					data.InspectionTypeValueUnit,
		InspectionTypePlannedValue:					data.InspectionTypePlannedValue,
		InspectionTypeCertificateType:				data.InspectionTypeCertificateType,
		InspectionTypeCertificateValueInText:		data.InspectionTypeCertificateValueInText,
		InspectionTypeCertificateValueInQuantity:	data.InspectionTypeCertificateValueInQuantity,
		InspectionPlanInspectionText:				data.InspectionPlanInspectionText,
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
