package hydros

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestDefaultWellService_Init(t *testing.T) {

	defaultWellService := (&DefaultWellService{DefaultService: &DefaultService{}}).Init(&ServiceSpec{ServiceName: "test"})
	assert.NotNil(t, defaultWellService, "Service should not be nil")
	assert.Equal(t, "test", defaultWellService.Spec.ServiceName)

	assert.NotNil(t, defaultWellService.CreateFunc, "CreateFunc should not be null")
	assert.Equal(t, reflect.TypeOf(defaultWellService.CreateFunc).Kind(), reflect.Func, "CreateFunc should be func")
	assert.NotNil(t, defaultWellService.GetFunc, "GetFunc should not be null")
	assert.Equal(t, reflect.TypeOf(defaultWellService.GetFunc).Kind(), reflect.Func, "GetFunc should be func")
	assert.NotNil(t, defaultWellService.CountFunc, "CountFunc should not be null")
	assert.Equal(t, reflect.TypeOf(defaultWellService.CountFunc).Kind(), reflect.Func, "CountFunc should be func")
	assert.NotNil(t, defaultWellService.ListFunc, "ListFunc should not be null")
	assert.Equal(t, reflect.TypeOf(defaultWellService.ListFunc).Kind(), reflect.Func, "ListFunc should be func")
}

func TestDefaultWellServiceCountFunc(t *testing.T) {

	defaultWellService := (&DefaultWellService{DefaultService: &DefaultService{}}).Init(&ServiceSpec{ServiceName: "test"})
	defaultWellService.CountFunc = func() (int, error) {
		return 314, nil
	}
	count, err := defaultWellService.CountFunc()
	assert.Nil(t, err, "Error should be nil.")
	assert.Equal(t, 314, count)
}

func TestDefaultWellServiceCreateFunc(t *testing.T) {

	defaultWellService := (&DefaultWellService{DefaultService: &DefaultService{}}).
		Init(&ServiceSpec{
			ServiceName:      "test",
			PayloadModelType: reflect.TypeOf(WellModel{}),
		})

	defaultWellService.CreateFunc = func(model *WellModel) (*WellModel, error) {
		return model, nil
	}
	returnedModel, err := defaultWellService.Create(&WellModel{DefaultModelBase: &DefaultModelBase{ID: 2718}})
	assert.Nil(t, err, "Error should be nil.")
	assert.Equal(t, uint(2718), returnedModel.ID)
}

func TestDefaultWellServiceGetFunc(t *testing.T) {

	defaultWellService := (&DefaultWellService{DefaultService: &DefaultService{}}).
		Init(&ServiceSpec{
			ServiceName:      "test",
			PayloadModelType: reflect.TypeOf(WellModel{}),
		})

	defaultWellService.GetFunc = func(ID uint) (*WellModel, error) {
		return &WellModel{DefaultModelBase: &DefaultModelBase{ID: ID}}, nil
	}
	returnedModel, err := defaultWellService.Get(2718)
	assert.Nil(t, err, "Error should be nil.")
	assert.Equal(t, uint(2718), returnedModel.ID)
}

func TestDefaultWellServiceListFunc(t *testing.T) {

	defaultWellService := (&DefaultWellService{DefaultService: &DefaultService{}}).
		Init(&ServiceSpec{
			ServiceName:      "test",
			PayloadModelType: reflect.TypeOf(WellModel{}),
		})

	defaultWellService.ListFunc = func(from int, size int, sort []Sort, ids []uint) ([]*WellModel, error) {
		list := make([]*WellModel, 1)
		list[0] = &WellModel{DefaultModelBase: &DefaultModelBase{ID: 235711}}
		return list, nil
	}
	returnedModels, err := defaultWellService.List(0, 1, nil, nil)
	assert.Nil(t, err, "Error should be nil.")
	assert.Equal(t, uint(235711), returnedModels[0].ID)
}
