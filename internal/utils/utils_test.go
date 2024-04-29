package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildEnpoint(t *testing.T) {
	assert.Equal(t, "0.0.0.0:8000", BuildEndpoint())

}

func TestBuildEnpointSpecifiedPortInvalid(t *testing.T) {
	t.Setenv("LOCO_PASTER_API_PORT", "foo")

	assert.Equal(t, "0.0.0.0:8000", BuildEndpoint())

}

func TestBuildEnpointSpecifiedPortValid(t *testing.T) {
	t.Setenv("LOCO_PASTER_API_PORT", "8001")

	assert.Equal(t, "0.0.0.0:8001", BuildEndpoint())

}
