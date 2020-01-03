// Package elastic registers a request modifier for generating parametrized queries
// to an elastic search service
package elastic

import (
	"github.com/google/martian/parse"
	"github.com/rsasidar/plugin-test/body/elastic-search/modifier"
)

func init() {
	parse.Register("body.ESQuery", modifier.FromJSON)
}
