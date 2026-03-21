package utils

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateProjectRef(t *testing.T) {
	t.Run("basic format slug-6hex", func(t *testing.T) {
		ref := GenerateProjectRef("My Project")
		// Should be lowercase
		assert.Equal(t, ref, strings.ToLower(ref))
		// Should contain a hyphen-separated random suffix
		parts := strings.Split(ref, "-")
		assert.GreaterOrEqual(t, len(parts), 2)
		// Last part should be 6 hex chars
		suffix := parts[len(parts)-1]
		assert.Len(t, suffix, 6)
	})

	t.Run("special characters stripped", func(t *testing.T) {
		ref := GenerateProjectRef("Hello! @World #123")
		// Should not contain special chars except hyphens
		for _, c := range ref {
			assert.True(t, (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '-',
				"unexpected char: %c", c)
		}
	})

	t.Run("long name truncated", func(t *testing.T) {
		ref := GenerateProjectRef("This is a very long project name that exceeds twenty characters")
		// Slug part (before random suffix) should be <= 20 chars
		parts := strings.Split(ref, "-")
		slug := strings.Join(parts[:len(parts)-1], "-")
		assert.LessOrEqual(t, len(slug), 20)
	})

	t.Run("uniqueness", func(t *testing.T) {
		a := GenerateProjectRef("test")
		b := GenerateProjectRef("test")
		// Same input should produce different refs due to random suffix
		assert.NotEqual(t, a, b)
	})

	t.Run("empty slug handled", func(t *testing.T) {
		ref := GenerateProjectRef("!!!")
		// Should still produce a valid ref with just the random suffix
		assert.NotEmpty(t, ref)
	})
}
