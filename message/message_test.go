package message

import (
	"testing"

	"gotest.tools/assert"
)

func TestNew(t *testing.T) {
	msg := New("", "")
	assert.Equal(t, msg.Name, "")
	assert.Equal(t, msg.Content, "")
}
