package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestControlMethod(t *testing.T) {
	cm := ControlMethod(5, 5, 5, 5)
	assert.Equal(t, 50, int(cm["Round_About"]))
	assert.Equal(t, 20, int(cm["Stop_Sign"]))
	assert.Equal(t, 90, int(cm["Traffic_Light"]))
	cm = ControlMethod(4, 5, 5, 5)
	assert.Equal(t, 75+10, int(cm["Round_About"]))
	assert.Equal(t, 30, int(cm["Stop_Sign"]))
	assert.Equal(t, 75, int(cm["Traffic_Light"]))
	cm = ControlMethod(2, 1, 1, 5)
	assert.Equal(t, 90+10, int(cm["Round_About"]))
	assert.Equal(t, 40, int(cm["Stop_Sign"]))
	assert.Equal(t, 30, int(cm["Traffic_Light"]))

}
