package utils

import (
	"testing"

	"github.com/jfgrea27/loco-paster/internal/models"
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

func TestFindPasteObjIndex(t *testing.T) {
	pobjs := []models.PasteObj{}

	assert.Equal(t, -1, FindPasteObjIndex(1, pobjs))

	pobjs = append(pobjs, models.PasteObj{Id: 1, Blob: "hello world"})
	assert.Equal(t, 0, FindPasteObjIndex(1, pobjs))
	assert.Equal(t, -1, FindPasteObjIndex(2, pobjs))
}
