package uplink

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	client := New("http://localhost:1234", "warehouse", "schema")
	assert.NotNil(t, client)
}
