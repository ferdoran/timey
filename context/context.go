package context

import (
	"errors"
	"fmt"
)

var bindings = make(map[string]interface{})

func Bind[T interface{}](qualifier string, instance *T) {
	if bindings[qualifier] != nil {
		panic(fmt.Sprintf("binding for qualifier %s already exists", qualifier))
	}

	bindings[qualifier] = instance
}

func Get[T interface{}](qualifier string) (*T, error) {
	binding := bindings[qualifier]
	if binding == nil {
		return nil, errors.New(fmt.Sprintf("no binding for qualifier %s", qualifier))
	}

	instance, ok := bindings[qualifier].(*T)

	if !ok {
		return nil, errors.New(fmt.Sprintf("binding for qualifier %s is not of type %T", qualifier, binding))
	}

	return instance, nil
}
