package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AddOne(t *testing.T) {

	assert.Equal(t, AddOne(2), 3, "AddOne should be 3")

}
