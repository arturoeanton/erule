package erule_test

import (
	"github.com/arturoeanton/erule"
	"github.com/robertkrimen/otto"
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
	_, r, _ := erule.Fire(p, `{"expo":1, "ou":"payroll","list":["a","b"]}`)
	t.Log(r.(erule.ResponseRisk).Path)
	if r.(erule.ResponseRisk).Value != 11 {
		t.Errorf("risk bad value")
	}

}

func TestFire3(t *testing.T) {
	p := `{ 
	"name": "first example", 
	"before": "console.log(\"p.Before\")",
	"after": "console.log(\"p.After\")",
	"custom_response": "s = {hola:\"chau\"}",
	"rules":[	 
		{"name": "rule A", "point": 1, "code": "data.expo != 1"}, 
		{"name": "rule B", "point": 10, "code": "data.ou == \"payroll\""},
		{"name": "rule C", "point": 11, "code": "data.expo == 2"}
		]
	}`
	_, r, _ := erule.Fire(p, `{"expo":1, "ou":"payroll","list":["a","b"]}`)
	d, _ := r.(otto.Value).Object().Get("hola")
	if d.String() != "chau" {
		t.Errorf("risk bad value")
	}

}

func TestFire4(t *testing.T) {
	p := `{ 
	"name": "first example", 
	"before": "console.log(\"p.Before\")",
	"after": "console.log(\"p.After\")",
	"rules":[	 
		{"name": "rule A", "point": 1, "code": "data.expo != 1"}, 
		{"name": "rule B", "point": 10, "code": "data.ou == \"payroll\""},
		{"name": "rule C", "point": 11, "code": "data.expo == 2"}
		]
	}`
	_, r, _ := erule.Fire(p, `{"expo":1, "ou":"payroll","list":["a","b"]}`)
	if r.(erule.ResponseRisk).Value!= 10 {
		t.Errorf("risk bad value")
	}

}

func TestFire5(t *testing.T) {

	p := `{ 
	"name": "first example", 
	"rules":[	 
		{"name": "rule A", "point": 1, "code": "data.expo == 1 &&  data.ou == \"payroll\""}, 
		{"name": "rule B", "point": 11, "code": "data.expo != 2"}
		]
	}`

	_, r, _ := erule.Fire(p, `{"expo":1, "ou":"payroll","list":["a","b"]}`)

	t.Log(r.(erule.ResponseRisk).Path)
	t.Log(r.(erule.ResponseRisk).Path)
	if r.(erule.ResponseRisk).Value != 12 {
		t.Errorf("risk bad value")
	}

}

func TestFire6(t *testing.T) {

	p := `{ 
	"name": "first example",
	"mode": "first_rule",
	"rules":[	 
		{"name": "rule A", "point": 1, "code": "data.expo != 1 &&  data.ou == \"payroll\""}, 
		{"name": "rule B", "point": 11, "code": "data.expo != 2"}
		]
	}`

	_, r, _ := erule.Fire(p, `{"expo":1, "ou":"payroll","list":["a","b"]}`)
	if r.(erule.ResponseFirst).Value != "rule B" {
		t.Errorf("risk bad value")
	}

}
