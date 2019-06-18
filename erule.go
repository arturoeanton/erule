package erule

import (
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

func Fire(p politics, data map[string]interface{})  ( int,  []string,  error) {
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
		}
	}
	return risk, path, nil
}
