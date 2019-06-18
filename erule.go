package erule

import (
	"encoding/json"
	"fmt"
	"github.com/mattn/anko/vm"
)

type Politics struct {
	Name  string `json:"name"`
	Rules []Rule `json:"rules"`
}
type Rule struct {
	Name  string `json:"name"`
	Point int    `json:"point"`
	Code  string `json:"code"`
}

const (
	RISK    	int  = 0
	FIRST_RULE  int  = 1
)

func Fire(str string, jsonString string) (int, []string, error) {
	return fireMode(str, jsonString, RISK )
}

func FireFirstRule(str string, jsonString string) (string, error) {
	_ , path , err := fireMode(str, jsonString, FIRST_RULE )
	return  path[0], err
}

func fireMode(str string, jsonString string, mode int ) (int, []string, error) {
	p := Politics{}
	if err := json.Unmarshal([]byte(str), &p); err != nil {
		return 0, nil, fmt.Errorf("Execute error: %v\n", err)
	}
	dataByt := []byte(jsonString)
	var data map[string]interface{}
	if err := json.Unmarshal(dataByt, &data); err != nil {
		return 0, nil, fmt.Errorf("Execute error: %v\n", err)
	}

	return fire(p, data, mode )
}

func fire(p Politics, data interface{}, mode int) (int, []string, error) {
	risk := 0
	var path []string
	env := vm.NewEnv()
	err := env.Define("data", data)
	if err != nil {
		return 0, nil, fmt.Errorf("Execute error: %v\n", err)
	}
	for _, rule := range p.Rules {
		script := rule.Code
		flag, err := env.Execute(script)
		if err != nil {
			return 0, nil, fmt.Errorf("Execute error: %v\n", err)
		}
		if flag.(bool) {

			risk += rule.Point
			path = append(path, rule.Name)
			if mode == FIRST_RULE {
				return 0, path, err
			}
		}

	}
	return risk, path, nil
}
