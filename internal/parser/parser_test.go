package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePkg(t *testing.T) {
	err := ParsePkg("github.com/go-po/generator/examples/counter/...")
	assert.NoError(t, err)
}
