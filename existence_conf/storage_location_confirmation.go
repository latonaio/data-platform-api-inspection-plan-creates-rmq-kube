package existence_conf

import (
	dpfm_api_input_reader "data-platform-api-inspection-plan-creates-rmq-kube/DPFM_API_Input_Reader"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *ExistenceConf) headerStorageLocationExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}
	exReqTimes := 0

	headers := make([]dpfm_api_input_reader.Header, 0, 1)
	headers = append(headers, input.Header)
	for _, header := range headers {
		bpID, plant, storageLocation := getHeaderStorageLocationExistenceConfKey(mapper, &header, exconfErrMsg)
		wg2.Add(1)
		exReqTimes++
		go func() {
			if isZero(bpID) || isZero(plant) || isZero(storageLocation) {
				wg2.Done()
				return
			}
			res, err := c.storageLocationExistenceConfRequest(bpID, plant, storageLocation, mapper, input, existenceMap, mtx, log)
			if err != nil {
				mtx.Lock()
				*errs = append(*errs, err)
				mtx.Unlock()
			}
			if res != "" {
				*exconfErrMsg = res
			}
			wg2.Done()
		}()
	}
	wg2.Wait()
	if exReqTimes == 0 {
		*existenceMap = append(*existenceMap, false)
	}
}

func (c *ExistenceConf) itemStorageLocationExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}
	exReqTimes := 0

	items := input.Header.Item
	for _, item := range items {
		bpID, plant, storageLocation := getItemStorageLocationExistenceConfKey(mapper, &item, exconfErrMsg)
		wg2.Add(1)
		exReqTimes++
		go func() {
			if isZero(bpID) || isZero(plant) || isZero(storageLocation) {
				wg2.Done()
				return
			}
			res, err := c.storageLocationExistenceConfRequest(bpID, plant, storageLocation, mapper, input, existenceMap, mtx, log)
			if err != nil {
				mtx.Lock()
				*errs = append(*errs, err)
				mtx.Unlock()
			}
			if res != "" {
				*exconfErrMsg = res
			}
			wg2.Done()
		}()
	}
	wg2.Wait()
	if exReqTimes == 0 {
		*existenceMap = append(*existenceMap, false)
	}
}

func (c *ExistenceConf) storageLocationExistenceConfRequest(bpID int, plant string, storageLocation string, mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
	keys := newResult(map[string]interface{}{
		"BusinessPartner": bpID,
		"Plant":           plant,
		"StorageLocation": storageLocation,
	})
	exist := false
	defer func() {
		mtx.Lock()
		*existenceMap = append(*existenceMap, exist)
		mtx.Unlock()
	}()

	req, err := jsonTypeConversion[Returns](input)
	if err != nil {
		return "", xerrors.Errorf("request create error: %w", err)
	}
	req.StorageLocationReturn.BusinessPartner = bpID
	req.StorageLocationReturn.Plant = plant
	req.StorageLocationReturn.StorageLocation = storageLocation

	exist, err = c.exconfRequest(req, mapper, log)
	if err != nil {
		return "", err
	}
	if !exist {
		return keys.fail(), nil
	}
	return "", nil
}

func getHeaderStorageLocationExistenceConfKey(mapper ExConfMapper, header *dpfm_api_input_reader.Header, exconfErrMsg *string) (int, string, string) {
	var bpID int
	var plant string
	var storageLocation string

	switch mapper.Field {
	case "OwnerProductionPlantStorageLocation":
		if header.OwnerProductionPlantBusinessPartner == nil ||
			header.OwnerProductionPlant == nil ||
			header.OwnerProductionPlantStorageLocation == nil {
			bpID = 0
			plant = ""
			storageLocation = ""
		} else {
			bpID = *header.OwnerProductionPlantBusinessPartner
			plant = *header.OwnerProductionPlant
			storageLocation = *header.OwnerProductionPlantStorageLocation
		}
	}

	return bpID, plant, storageLocation
}

func getItemStorageLocationExistenceConfKey(mapper ExConfMapper, item *dpfm_api_input_reader.Item, exconfErrMsg *string) (int, string, string) {
	var bpID int
	var plant string
	var storageLocation string

	switch mapper.Field {
	case "ProductionPlantStorageLocation":
		if item.ProductionPlantBusinessPartner == nil ||
			item.ProductionPlant == nil ||
			item.ProductionPlantStorageLocation == nil {
			bpID = 0
			plant = ""
			storageLocation = ""
		} else {
			bpID = *item.ProductionPlantBusinessPartner
			plant = *item.ProductionPlant
			storageLocation = *item.ProductionPlantStorageLocation
		}
	}

	return bpID, plant, storageLocation
}
