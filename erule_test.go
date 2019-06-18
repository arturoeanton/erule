package erule_test

import (
	"github.com/arturoeanton/erule"
	"testing"
)

func TestFire(t *testing.T) {
	p := erule.Politics{
		Name: "first rule",
		Rules: []erule.Rule{
			{Name: "first rule", Point: 1, Code: `data.expo == 1`},
			{Name: "second rule", Point: 10, Code: `data.ou == "payroll"`},
			{Name: "third rule", Point: 11, Code: `data.expo == 2`},
		},
	}
	risk, _, _ :=  erule.Fire(p, `{"expo":1, "ou":"payroll","list":["a","b"]}`)
	if risk != 11 {

	}

}

/*
package main

import (
	"encoding/json"
	"fmt"
)



func main() {
	byt := []byte(`{"expo":1, "ou":"payroll","list":["a","b"]}`)
	var data map[string]interface{}
	if err := json.Unmarshal(byt, &data); err != nil {
		panic(err)
	}
	p := erule.Politics{
		Name: "first rule",
		Rules: []rule{
			{Name: "first rule", Point: 1, Code: `data.expo == 1`},
			{Name: "second rule", Point: 10, Code: `data.ou == "payroll"`},
			{Name: "third rule", Point: 11, Code: `data.expo == 2`},
		},
	}
	risk, path, _ := erule.Fire(p,data)
	fmt.Printf("risk: %d \npath:%q\n",risk,path)
}

*/
