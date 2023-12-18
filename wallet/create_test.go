package wallet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreate(t *testing.T) {
	err := Create("whf")
	assert.Nil(t, err)
}
