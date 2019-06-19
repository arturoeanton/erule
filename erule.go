package erule

import (
	"encoding/json"
	"fmt"
	"github.com/robertkrimen/otto"
)

const (
	MODE_FIRST     string = "first_rule"
	ERROR          string = "error"
	OTTO_VALUE            = "otto.value"
	RESPONSE_FIRST        = "erule.ResponseFirst"
	RESPONSE_RISK         = "erule.ResponseRisk"
)

type Politics struct {
	Name           string `json:"name"`
	Rules          []Rule `json:"rules"`
	Mode           string `jons:"mode"`
	After          string `json:"after"`
	Before         string `json:"before"`
	CustomResponse string `json:"custom_response"`
}
type Rule struct {
	Name   string `json:"name"`
	Point  int    `json:"point"`
	Code   string `json:"code"`
	After  string `json:"after"`
	Before string `json:"before"`
}

type ResponseFirst struct {
	Value string `json:"value"`
}
type ResponseRisk struct {
	Value int `json:"value"`
	Path  []string `json:"path"`
}

func Fire(str string, jsonString string, params ...interface{}) (string, interface{}, error) {
	return fireMode(str, jsonString, params)
}

func fireMode(str string, jsonString string, params ...interface{}) (string, interface{}, error) {
	p := Politics{}
	if err := json.Unmarshal([]byte(str), &p); err != nil {
		return ERROR, nil, fmt.Errorf("Execute error: %v\n", err)
	}
	dataByt := []byte(jsonString)
	var data map[string]interface{}
	if err := json.Unmarshal(dataByt, &data); err != nil {
		return ERROR, nil, fmt.Errorf("Execute error: %v\n", err)
	}

	return fire(p, data, params)
}

func fire(p Politics, data interface{}, params ...interface{}) (string, interface{}, error) {
	risk := 0
	var path []string
	vm := otto.New()

	err := vm.Set("data", data)
	if err != nil {
		return ERROR, nil, fmt.Errorf("Execute error(data): %v\n", err)
	}

	err = vm.Set("params", params)
	if err != nil {
		return ERROR, nil, fmt.Errorf("Execute error(params): %v \n", err)
	}

	if p.Before != "" {
		vm.Run(p.Before)
	}
	for _, rule := range p.Rules {

		if rule.Before != "" {
			vm.Run(rule.Before)
		}

		script := rule.Code
		ret, err := vm.Run(script)
		if err != nil {
			return ERROR, nil, fmt.Errorf("Execute error: %v\n", err)
		}
		flag, _ := ret.ToBoolean()

		if rule.After != "" {
			vm.Run(rule.After)
		}

		if flag {
			risk += rule.Point
			path = append(path, rule.Name)

			vm.Set("value", risk)
			vm.Set("path", path)

			if p.Mode == MODE_FIRST {
				break
			}
		}

	}
	if p.After != "" {
		vm.Run(p.After)
	}

	if p.CustomResponse != "" {
		v, err := vm.Run(p.CustomResponse)

		return OTTO_VALUE, v, err
	}

	if p.Mode == MODE_FIRST {
		return RESPONSE_FIRST, ResponseFirst{Value: path[0]}, err
	}

	return RESPONSE_RISK, ResponseRisk{Value: risk, Path: path}, nil
}
