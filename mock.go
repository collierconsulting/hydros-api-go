package hydros

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// ModelServiceCallMock holds mock definition for service instance
type ModelServiceCallMock struct {
	MockFunc interface{}
}

func MockModelServiceMethod(service interface{}, targetModelMethod string, mockFunc interface{}) error {
	if _, ok := service.(Service); !ok {
		return errors.New("service arguments must implement Service interface")
	}
	//modelType := service.(Service)._ServiceSpec().PayloadModelType
	//if _, ok := modelType.MethodByName(targetModelMethod); !ok {
	//	return fmt.Errorf("could not find method '%s' in service payload model '%s'", targetModelMethod, modelType.Name())
	//}
	service.(Service)._SetModelServiceCallMock(targetModelMethod, &ModelServiceCallMock{MockFunc: mockFunc})
	return nil
}

// MockServiceMethod mock a method on a service
func MockServiceMethod(client interface{}, targetServiceMethod string, mockFunc interface{}) error {

	targetMethodParts := strings.Split(targetServiceMethod, ".")
	if len(targetMethodParts) != 2 {
		return errors.New("targetServiceMethod arguments should be of format Service.Method (e.g. 'Well.Get')")
	}

	// Verify service
	serviceName := targetMethodParts[0]

	service, err := extractServiceFromClient(client, serviceName)
	if err != nil {
		return err
	}

	// Check method exists
	targetMethodName := targetMethodParts[1]
	//var targetMethod reflect.Method
	if _, ok := reflect.TypeOf(service).MethodByName(targetMethodName); !ok {
		return fmt.Errorf("could not find method '%s' in service '%s'", targetMethodName, serviceName)
	}

	// Look for method backing function
	mockableFunctionName := fmt.Sprintf("%sFunc", targetMethodName)
	serviceElem := reflect.ValueOf(service).Elem()
	var mockableFunction interface{}
	for i := 0; i < serviceElem.NumField(); i++ {
		if serviceElem.Type().Field(i).Name == mockableFunctionName {
			mockableFunction = serviceElem.Field(i).Interface()
		}
	}

	if mockableFunction == nil {
		return fmt.Errorf("could not find backing function '%s' for %s service", mockableFunctionName, serviceName)
	}

	targetMethodType := reflect.TypeOf(mockableFunction)
	if targetMethodType != reflect.TypeOf(mockFunc) {
		return fmt.Errorf("mock function is wrong type: expected: %s but got: %s",
			targetMethodType, reflect.TypeOf(mockFunc))
	}

	// Replace function
	if !serviceElem.FieldByName(mockableFunctionName).CanSet() {
		return fmt.Errorf("can not set backing function %s on service %s", mockableFunctionName, serviceName)
	}
	serviceElem.FieldByName(mockableFunctionName).Set(reflect.ValueOf(mockFunc))

	return nil
}

func extractServiceFromClient(client interface{}, serviceName string) (interface{}, error) {
	clientElem := reflect.ValueOf(client).Elem()
	var service interface{}
	for i := 0; i < clientElem.NumField(); i++ {
		if clientElem.Type().Field(i).Name == serviceName {
			service = clientElem.Field(i).Interface()
		}
	}

	if service == nil {
		// Could not find at struct top level...let's check if the Client exists as anonymous field
		for i := 0; i < clientElem.NumField(); i++ {
			if clientElem.Type().Field(i).Name == "Client" {
				return extractServiceFromClient(clientElem.Field(i).Interface(), serviceName)
			}
		}
		return nil, fmt.Errorf("could not find service '%s'", serviceName)
	} else if _, ok := service.(Service); !ok {
		return nil, fmt.Errorf("could not find service '%s' that implements Service", serviceName)
	}

	return service, nil

}
