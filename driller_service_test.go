package hydros

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestDefaultDrillerService_Init(t *testing.T) {

	defaultDrillerService := (&DefaultDrillerService{DefaultService: &DefaultService{}}).Init(&ServiceSpec{ServiceName: "test"})
	assert.NotNil(t, defaultDrillerService, "Service should not be nil")
	assert.Equal(t, "test", defaultDrillerService.Spec.ServiceName)

	assert.NotNil(t, defaultDrillerService.CreateFunc, "CreateFunc should not be null")
	assert.Equal(t, reflect.TypeOf(defaultDrillerService.CreateFunc).Kind(), reflect.Func, "CreateFunc should be func")
	assert.NotNil(t, defaultDrillerService.GetFunc, "GetFunc should not be null")
	assert.Equal(t, reflect.TypeOf(defaultDrillerService.GetFunc).Kind(), reflect.Func, "GetFunc should be func")
	assert.NotNil(t, defaultDrillerService.CountFunc, "CountFunc should not be null")
	assert.Equal(t, reflect.TypeOf(defaultDrillerService.CountFunc).Kind(), reflect.Func, "CountFunc should be func")
	assert.NotNil(t, defaultDrillerService.ListFunc, "ListFunc should not be null")
	assert.Equal(t, reflect.TypeOf(defaultDrillerService.ListFunc).Kind(), reflect.Func, "ListFunc should be func")
}

func TestDefaultDrillerServiceCountFunc(t *testing.T) {

	defaultDrillerService := (&DefaultDrillerService{DefaultService: &DefaultService{}}).Init(&ServiceSpec{ServiceName: "test"})
	defaultDrillerService.CountFunc = func() (int, error) {
		return 314, nil
	}
	count, err := defaultDrillerService.CountFunc()
	assert.Nil(t, err, "Error should be nil.")
	assert.Equal(t, 314, count)
}

func TestDefaultDrillerServiceCreateFunc(t *testing.T) {

	defaultDrillerService := (&DefaultDrillerService{DefaultService: &DefaultService{}}).
		Init(&ServiceSpec{
			ServiceName:      "test",
			PayloadModelType: reflect.TypeOf(DrillerModel{}),
		})

	defaultDrillerService.CreateFunc = func(model *DrillerModel) (*DrillerModel, error) {
		return model, nil
	}
	returnedModel, err := defaultDrillerService.Create(&DrillerModel{DefaultModelBase: &DefaultModelBase{ID: 2718}})
	assert.Nil(t, err, "Error should be nil.")
	assert.Equal(t, uint(2718), returnedModel.ID)
}

func TestDefaultDrillerServiceGetFunc(t *testing.T) {

	defaultDrillerService := (&DefaultDrillerService{DefaultService: &DefaultService{}}).
		Init(&ServiceSpec{
			ServiceName:      "test",
			PayloadModelType: reflect.TypeOf(DrillerModel{}),
		})

	defaultDrillerService.GetFunc = func(ID uint) (*DrillerModel, error) {
		return &DrillerModel{DefaultModelBase: &DefaultModelBase{ID: ID}}, nil
	}
	returnedModel, err := defaultDrillerService.Get(2718)
	assert.Nil(t, err, "Error should be nil.")
	assert.Equal(t, uint(2718), returnedModel.ID)
}

func TestDefaultDrillerServiceListFunc(t *testing.T) {

	defaultDrillerService := (&DefaultDrillerService{DefaultService: &DefaultService{}}).
		Init(&ServiceSpec{
			ServiceName:      "test",
			PayloadModelType: reflect.TypeOf(DrillerModel{}),
		})

	defaultDrillerService.ListFunc = func(from int, size int, sort []Sort, ids []uint) ([]*DrillerModel, error) {
		list := make([]*DrillerModel, 1)
		list[0] = &DrillerModel{DefaultModelBase: &DefaultModelBase{ID: 235711}}
		return list, nil
	}
	returnedModels, err := defaultDrillerService.List(0, 1, nil, nil)
	assert.Nil(t, err, "Error should be nil.")
	assert.Equal(t, uint(235711), returnedModels[0].ID)
}
