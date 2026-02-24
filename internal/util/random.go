package util

import (
	"math/rand/v2"
	"strings"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateFastString is thread-safe and extremely fast.
// Ideal for: IDs, filenames, or non-sensitive tokens.
func GenerateFastString(n int) string {
	var sb strings.Builder
	sb.Grow(n)
	for range n {
		// rand.IntN is thread-safe in v2
		//nolint:gosec // G404: Non-crypto random is intentional for fast, non-sensitive string generation
		idx := rand.IntN(len(charset))
		sb.WriteByte(charset[idx])
	}

	return sb.String()
}
