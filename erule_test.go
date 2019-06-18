package erule_test

import (
	"github.com/arturoeanton/erule"
	"testing"
)

func TestFire1(t *testing.T) {
	p := `{ 
	"name": "first example", 
	"rules":[	 
		{"name": "rule A", "point": 1, "code": "data.expo == 1"}, 
		{"name": "rule B", "point": 10, "code": "data.ou == \"payroll\""},
		{"name": "rule C", "point": 11, "code": "data.expo == 2"}
		]
	}`
	risk, path, _ :=  erule.Fire(p, `{"expo":1, "ou":"payroll","list":["a","b"]}`)
	t.Log(path)
	if risk != 11 {
		t.Errorf("risk bad value")
	}

}


func TestFire2(t *testing.T) {
	p := `{ 
	"name": "first example", 
	"rules":[	 
		{"name": "rule A", "point": 1, "code": "data.expo != 1"}, 
		{"name": "rule B", "point": 10, "code": "data.ou == \"payroll\""},
		{"name": "rule C", "point": 11, "code": "data.expo == 2"}
		]
	}`
	risk, path, _ :=  erule.Fire(p, `{"expo":1, "ou":"payroll","list":["a","b"]}`)
	t.Log(path)
	if risk != 10 {
		t.Errorf("risk bad value")
	}

}

func TestFire3(t *testing.T) {

	p := `{ 
	"name": "first example", 
	"rules":[	 
		{"name": "rule A", "point": 1, "code": "data.expo == 1 &&  data.ou == \"payroll\""}, 
		{"name": "rule B", "point": 11, "code": "data.expo != 2"}
		]
	}`

	risk, path, _ :=  erule.Fire(p, `{"expo":1, "ou":"payroll","list":["a","b"]}`)
	t.Log(path)
	if risk != 12 {
		t.Errorf("risk bad value")
	}

}