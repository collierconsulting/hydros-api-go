package client

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestDefaultDrillerService_Init(t *testing.T) {

	defaultDrillerService := (&DefaultDrillerService{DefaultService: &DefaultService{}}).Init(&ServiceSpec{ServiceName: "test"})
	assert.NotNil(t, defaultDrillerService, "Service should not be nil")
	assert.Equal(t, "test", defaultDrillerService.Spec.ServiceName)

	assert.NotNil(t, defaultDrillerService._Create, "_Create should not be null")
	assert.Equal(t, reflect.TypeOf(defaultDrillerService._Create).Kind(), reflect.Func, "_Create should be func")
	assert.NotNil(t, defaultDrillerService._Get, "_Get should not be null")
	assert.Equal(t, reflect.TypeOf(defaultDrillerService._Get).Kind(), reflect.Func, "_Get should be func")
	assert.NotNil(t, defaultDrillerService._Count, "_Count should not be null")
	assert.Equal(t, reflect.TypeOf(defaultDrillerService._Count).Kind(), reflect.Func, "_Count should be func")
	assert.NotNil(t, defaultDrillerService._List, "_List should not be null")
	assert.Equal(t, reflect.TypeOf(defaultDrillerService._List).Kind(), reflect.Func, "_List should be func")
}

func TestDefaultDrillerService_Count(t *testing.T) {

	defaultDrillerService := (&DefaultDrillerService{DefaultService: &DefaultService{}}).Init(&ServiceSpec{ServiceName: "test"})
	defaultDrillerService._Count = func() (int, error) {
		return 314, nil
	}
	count, err := defaultDrillerService._Count()
	assert.Nil(t, err, "Error should be nil.")
	assert.Equal(t, 314, count)
}

func TestDefaultDrillerService_Create(t *testing.T) {

	defaultDrillerService := (&DefaultDrillerService{DefaultService: &DefaultService{}}).
		Init(&ServiceSpec{
			ServiceName:      "test",
			PayloadModelType: reflect.TypeOf(DrillerModel{}),
		})

	defaultDrillerService._Create = func(model *DrillerModel) (*DrillerModel, error) {
		return model, nil
	}
	returnedModel, err := defaultDrillerService.Create(&DrillerModel{DefaultModelBase: &DefaultModelBase{ID: 2718}})
	assert.Nil(t, err, "Error should be nil.")
	assert.Equal(t, uint(2718), returnedModel.ID)
}

func TestDefaultDrillerService_Get(t *testing.T) {

	defaultDrillerService := (&DefaultDrillerService{DefaultService: &DefaultService{}}).
		Init(&ServiceSpec{
			ServiceName:      "test",
			PayloadModelType: reflect.TypeOf(DrillerModel{}),
		})

	defaultDrillerService._Get = func(ID uint) (*DrillerModel, error) {
		return &DrillerModel{DefaultModelBase: &DefaultModelBase{ID: ID}}, nil
	}
	returnedModel, err := defaultDrillerService.Get(2718)
	assert.Nil(t, err, "Error should be nil.")
	assert.Equal(t, uint(2718), returnedModel.ID)
}

func TestDefaultDrillerService_List(t *testing.T) {

	defaultDrillerService := (&DefaultDrillerService{DefaultService: &DefaultService{}}).
		Init(&ServiceSpec{
			ServiceName:      "test",
			PayloadModelType: reflect.TypeOf(DrillerModel{}),
		})

	defaultDrillerService._List = func(from int, size int, sort []Sort, ids []int) ([]*DrillerModel, error) {
		list := make([]*DrillerModel, 1)
		list[0] = &DrillerModel{DefaultModelBase: &DefaultModelBase{ID: 235711}}
		return list, nil
	}
	returnedModels, err := defaultDrillerService.List(0, 1, nil, nil)
	assert.Nil(t, err, "Error should be nil.")
	assert.Equal(t, uint(235711), returnedModels[0].ID)
}
