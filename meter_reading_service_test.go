package hydros

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestDefaultMeterReadingService_Init(t *testing.T) {
	defaultMeterReadingService := (&DefaultMeterReadingService{DefaultService: &DefaultService{}}).Init(&ServiceSpec{ServiceName: "test"})
	assert.NotNil(t, defaultMeterReadingService, "Service should not be nil")
	assert.Equal(t, "test", defaultMeterReadingService.Spec.ServiceName)

	assert.NotNil(t, defaultMeterReadingService.GetProductionByWellFunc, "GetProductionByWellFunc should not be null")
	assert.Equal(t, reflect.TypeOf(defaultMeterReadingService.GetProductionByWellFunc).Kind(), reflect.Func, "GetProductionByWellFunc should be func")
}

func TestDefaultMeterReadingServiceGetFunc(t *testing.T) {
	wellId := uint(1)
	meterId := uint(1)
	id := uint(1)

	defaultMeterReadingService := (&DefaultMeterReadingService{DefaultService: &DefaultService{}}).
		Init(&ServiceSpec{
			ServiceName:      "test",
			PayloadModelType: reflect.TypeOf(PermitModel{}),
		})

	defaultMeterReadingService.GetFunc = func(wellID uint, meterID uint, ID uint) (*MeterReadingModel, error) {
		return &MeterReadingModel{DefaultModelBase: &DefaultModelBase{ID: ID}}, nil
	}
	returnedModel, err := defaultMeterReadingService.Get(wellId, meterId, id)
	assert.Nil(t, err, "Error should be nil.")
	assert.Equal(t, id, returnedModel.ID)
}

func TestDefaultMeterReadingServiceGetProductionByWellFunc(t *testing.T) {

	defaultMeterReadingService := (&DefaultMeterReadingService{DefaultService: &DefaultService{}}).
		Init(&ServiceSpec{
			ServiceName:      "test",
			PayloadModelType: reflect.TypeOf(PermitModel{}),
		})

	defaultMeterReadingService.GetProductionByWellFunc = func(wellID uint, fromDate *time.Time, toDate *time.Time, estimateBounds bool) ([]ProductionModel, error) {
		list := make([]ProductionModel, 1)
		list[0] = ProductionModel{MeterID: uint(100)}
		return list, nil
	}
	returnedModels, err := defaultMeterReadingService.GetProductionByWell(1, nil, nil, false)
	assert.Nil(t, err, "Error should be nil.")
	assert.Equal(t, uint(100), returnedModels[0].MeterID)
}


func TestDefaultMeterReadingServiceGetProductionByWellAndMeterFunc(t *testing.T) {

	defaultMeterReadingService := (&DefaultMeterReadingService{DefaultService: &DefaultService{}}).
		Init(&ServiceSpec{
			ServiceName:      "test",
			PayloadModelType: reflect.TypeOf(PermitModel{}),
		})

	defaultMeterReadingService.GetProductionByWellAndMeterFunc = func(wellID uint, meterID uint, fromDate *time.Time, toDate *time.Time, estimateBounds bool) (*ProductionModel, error) {
		productionModel := ProductionModel{MeterID: uint(101)}
		return &productionModel, nil
	}
	returnedModel, err := defaultMeterReadingService.GetProductionByWellAndMeter(1, 1, nil, nil, false)
	assert.Nil(t, err, "Error should be nil.")
	assert.Equal(t, uint(101), returnedModel.MeterID)
}