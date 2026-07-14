package main

import (
	"testing"

	corev2 "github.com/sensu/sensu-go/api/core/v2"
	"github.com/sensu/sensu-plugin-sdk/sensu"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
}

func TestCheckArgs(t *testing.T) {
	assert := assert.New(t)
	event := corev2.FixtureEvent("entity1", "check1")

	// Valid thresholds
	plugin.Warning = 1
	plugin.Critical = 2
	i, e := checkArgs(event)
	assert.Equal(sensu.CheckStateOK, i)
	assert.NoError(e)

	// warning not set
	plugin.Warning = 0
	plugin.Critical = 2
	i, e = checkArgs(event)
	assert.Equal(sensu.CheckStateCritical, i)
	assert.Error(e)

	// critical not set
	plugin.Warning = 1
	plugin.Critical = 0
	i, e = checkArgs(event)
	assert.Equal(sensu.CheckStateCritical, i)
	assert.Error(e)

	// critical < warning
	plugin.Warning = 5
	plugin.Critical = 2
	i, e = checkArgs(event)
	assert.Equal(sensu.CheckStateCritical, i)
	assert.Error(e)

	// equal thresholds are valid
	plugin.Warning = 3
	plugin.Critical = 3
	i, e = checkArgs(event)
	assert.Equal(sensu.CheckStateOK, i)
	assert.NoError(e)
}
