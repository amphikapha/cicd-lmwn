package calculator_test

import (
	"testing"

	. "github.com/amphikapha/cicd-lmwn/internal/calculator"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	actual := Add(1, 2)
	assert.Equal(t, 3, actual)
}
