//go:build unit

package pocket

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddFloat(t *testing.T) {
	z := AddFloat(float64(0.2), float64(0.1))
	assert.Equal(t, float64(0.3), z)
}

func TestMinusFloat(t *testing.T) {
	z := MinusFloat(float64(0.2), float64(0.1))
	assert.Equal(t, float64(0.1), z)
}
