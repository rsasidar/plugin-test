package main

import (
	"github.com/devopsfaith/krakend-martian/register"
	"github.com/rsasidar/plugin-test/body/elastic-search/modifier"
)

func init() {
	register.Set("body.ESQuery", []register.Scope{register.ScopeRequest}, func(b []byte) (interface{}, error) {
		return modifier.FromJSON(b)
	})
}

func main() {

}
