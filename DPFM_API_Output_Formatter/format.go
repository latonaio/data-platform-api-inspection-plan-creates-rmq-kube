package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-inspection-plan-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "data-platform-api-inspection-plan-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-inspection-plan-creates-rmq-kube/sub_func_complementer"
	"encoding/json"

	"golang.org/x/xerrors"
)

func ConvertToHeaderCreates(subfuncSDC *sub_func_complementer.SDC) (*Header, error) {
	data := subfuncSDC.Message.Header

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToInspectionCreates(subfuncSDC *sub_func_complementer.SDC) (*[]Inspection, error) {
	inspections := make([]Inspection, 0)

	for _, data := range *subfuncSDC.Message.Inspection {
		inspection, err := TypeConverter[*Inspection](data)
		if err != nil {
			return nil, err
		}

		inspections = append(inspections, *inspection)
	}

	return &inspections, nil
}

func ConvertToOperationCreates(subfuncSDC *sub_func_complementer.SDC) (*[]Operation, error) {
	operations := make([]Operation, 0)

	for _, data := range *subfuncSDC.Message.Operation {
		operation, err := TypeConverter[*Operation](data)
		if err != nil {
			return nil, err
		}

		operations = append(operations, *operation)
	}

	return &operations, nil
}

func ConvertToHeaderUpdates(headerData dpfm_api_input_reader.Header) (*Header, error) {
	data := headerData

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToInspectionUpdates(itemUpdates *[]dpfm_api_processing_formatter.InspectionUpdates) (*[]Inspection, error) {
	inspections := make([]Inspection, 0)

	for _, data := range *inspectionUpdates {
		inspection, err := TypeConverter[*Inspection](data)
		if err != nil {
			return nil, err
		}

		inspections = append(inspections, *inspection)
	}

	return &inspections, nil
}

func TypeConverter[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}
