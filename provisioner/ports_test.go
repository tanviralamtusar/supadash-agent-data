package provisioner

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewPortAllocator(t *testing.T) {
	pa := NewPortAllocator(54320, 8000)
	assert.NotNil(t, pa)
	assert.Equal(t, 54320, pa.baseDBPort)
	assert.Equal(t, 8000, pa.baseAPIPort)
	assert.Empty(t, pa.allocatedPorts)
}

func TestAllocatePorts(t *testing.T) {
	// Use high ports to avoid conflicts with real services
	pa := NewPortAllocator(60000, 61000)

	t.Run("returns valid allocation", func(t *testing.T) {
		alloc, err := pa.AllocatePorts("project-1")
		require.NoError(t, err)
		require.NotNil(t, alloc)

		assert.Greater(t, alloc.DBPort, 0)
		assert.Greater(t, alloc.APIPort, 0)
		assert.Greater(t, alloc.StudioPort, 0)
		assert.Greater(t, alloc.PoolerPort, 0)
	})

	t.Run("two projects get different ports", func(t *testing.T) {
		pa2 := NewPortAllocator(60000, 61000)
		alloc1, err := pa2.AllocatePorts("proj-a")
		require.NoError(t, err)
		alloc2, err := pa2.AllocatePorts("proj-b")
		require.NoError(t, err)

		// DB ports must differ
		assert.NotEqual(t, alloc1.DBPort, alloc2.DBPort)
		// API ports must differ
		assert.NotEqual(t, alloc1.APIPort, alloc2.APIPort)
	})
}

func TestReleasePorts(t *testing.T) {
	pa := NewPortAllocator(60000, 61000)

	alloc, err := pa.AllocatePorts("test-release")
	require.NoError(t, err)

	// Verify ports are tracked
	assert.Contains(t, pa.allocatedPorts, alloc.DBPort)

	// Release
	pa.ReleasePorts("test-release")

	// Verify ports are freed
	assert.NotContains(t, pa.allocatedPorts, alloc.DBPort)
	assert.NotContains(t, pa.allocatedPorts, alloc.APIPort)
}

func TestRegisterExistingPorts(t *testing.T) {
	pa := NewPortAllocator(60000, 61000)

	existing := PortAllocation{
		DBPort:        60000,
		APIPort:       61000,
		APIPortHTTPS:  61001,
		StudioPort:    61002,
		PoolerPort:    60001,
		AnalyticsPort: 61003,
	}
	pa.RegisterExistingPorts("existing-proj", existing)

	// All ports should be tracked
	assert.Equal(t, "existing-proj", pa.allocatedPorts[60000])
	assert.Equal(t, "existing-proj", pa.allocatedPorts[61000])

	// New allocation should skip the registered block
	alloc, err := pa.AllocatePorts("new-proj")
	require.NoError(t, err)
	assert.NotEqual(t, 60000, alloc.DBPort)
	assert.NotEqual(t, 61000, alloc.APIPort)
}
