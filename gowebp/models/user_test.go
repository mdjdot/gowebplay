package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AddUser(t *testing.T) {
	id, err := (&User{}).AddUser("王五", "1qaz@WSX", "123asdd@144.com")
	assert.NoError(t, err)
	assert.Greater(t, id, int64(0))
}
