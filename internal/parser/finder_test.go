package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/go/packages"
)

func TestFindPackgePO(t *testing.T) {
	// setup
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedTypes | packages.NeedTypesInfo,
	}, "github.com/go-po/generator/examples/counter")
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	// execute
	po, err := findPackagePo(pkgs)

	// verify
	if assert.NoError(t, err) && assert.NotNil(t, po) {
		assert.Equal(t, "po", po.Name())
		assert.Equal(t, "github.com/go-po/po", po.Path())
		assert.NotNil(t, po.Scope())
		assert.True(t, po.Complete())
	}
}
