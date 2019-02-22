package hydros

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestDefaultMeterService_init(t *testing.T) {
	defaultMeterService := (&DefaultMeterService{DefaultService: &DefaultService{}}).Init(&ServiceSpec{ServiceName: "test"})
	assert.NotNil(t, defaultMeterService, "Service should not be nil")
	assert.Equal(t, "test", defaultMeterService.Spec.ServiceName)

	assert.NotNil(t, defaultMeterService.GetFunc, "GetFun should not be null")
	assert.Equal(t, reflect.TypeOf(defaultMeterService.GetFunc).Kind(), reflect.Func, "GetFunc should be func")
	assert.NotNil(t, defaultMeterService.CreateFunc, "CreateFunc should not be null")
	assert.Equal(t, reflect.TypeOf(defaultMeterService.CreateFunc).Kind(), reflect.Func, "CreateFunc should be func")
	assert.NotNil(t, defaultMeterService.UpdateFunc, "UpdateFunc should not be null")
	assert.Equal(t, reflect.TypeOf(defaultMeterService.UpdateFunc).Kind(), reflect.Func, "UpdateFunc should be func")
	assert.NotNil(t, defaultMeterService.DecommissionFunc, "DecommissionFunc should not be null")
	assert.Equal(t, reflect.TypeOf(defaultMeterService.DecommissionFunc).Kind(), reflect.Func, "DecommissionFunc should be func")
}

func TestDefaultMeterServiceCreateFunc(t *testing.T) {

	defaultMeterService := (&DefaultMeterService{DefaultService: &DefaultService{}}).
		Init(&ServiceSpec{
			ServiceName:      "test",
			PayloadModelType: reflect.TypeOf(MeterModel{}),
		})

	defaultMeterService.CreateFunc = func(model *MeterModel) (*MeterModel, error) {
		return model, nil
	}
	returnedModel, err := defaultMeterService.Create(&MeterModel{DefaultModelBase: &DefaultModelBase{ID: 3332}})
	assert.Nil(t, err, "Error should be nil.")
	assert.Equal(t, uint(3332), returnedModel.ID)
}

func TestDefaultMeterServiceGetFunc(t *testing.T) {

	defaultMeterService := (&DefaultMeterService{DefaultService: &DefaultService{}}).
		Init(&ServiceSpec{
			ServiceName:      "test",
			PayloadModelType: reflect.TypeOf(MeterModel{}),
		})

	defaultMeterService.GetFunc = func(wellID uint, ID uint) (*MeterModel, error) {
		return &MeterModel{DefaultModelBase: &DefaultModelBase{ID: ID}}, nil
	}
	returnedModel, err := defaultMeterService.Get(1, 3333)
	assert.Nil(t, err, "Error should be nil.")
	assert.Equal(t, uint(3333), returnedModel.ID)
}

func TestDefaultMeterServiceUpdateFunc(t *testing.T) {

	defaultMeterService := (&DefaultMeterService{DefaultService: &DefaultService{}}).
		Init(&ServiceSpec{
			ServiceName:      "test",
			PayloadModelType: reflect.TypeOf(MeterModel{}),
		})

	defaultMeterService.UpdateFunc = func(model *MeterModel) (*MeterModel, error) {
		return model, nil
	}
	returnedModel, err := defaultMeterService.Update(&MeterModel{DefaultModelBase: &DefaultModelBase{ID: 2}})
	assert.Nil(t, err, "Error should be nil.")
	assert.Equal(t, uint(2), returnedModel.ID)
}

func TestDefaultMeterServiceDecommissionFunc(t *testing.T) {

	defaultMeterService := (&DefaultMeterService{DefaultService: &DefaultService{}}).
		Init(&ServiceSpec{
			ServiceName:      "test",
			PayloadModelType: reflect.TypeOf(MeterModel{}),
		})

	defaultMeterService.DecommissionFunc = func(id uint, decommissionDate time.Time) (*MeterModel, error) {
		model := MeterModel{DefaultModelBase: &DefaultModelBase{ID: id}, DecomissionDate: decommissionDate}
		return &model, nil
	}

	decomTime := time.Now()

	returnedModel, err := defaultMeterService.Decommission(2, decomTime)
	assert.Nil(t, err, "Error should be nil.")
	assert.Equal(t, uint(2), returnedModel.ID)
	assert.Equal(t, decomTime, returnedModel.DecomissionDate)
}
