// Package ids encapsulates common id generation uses in namespacelabs projects.
package ids

import (
	"crypto/rand"
)

// NewID returns a random ID in base32, which is not as compact as say base62 but is
// usable in scenarios which are case agnostic, and only allow alphanumeric values,
// such as subdomain segments.
func NewRandomBase32ID(len int) string {
	b := make([]byte, len)
	rand.Read(b) // rand.Read is guaranteed to not return errors.
	return EncodeToBase32String(b)
}

// Returns a base62-encoded ID from an alphanumeric dictionary, best suitable in
// scenarios which retain case. For IDs that need to managed manually by people,
// or need to be embedded into domain names or URLs, prefer `NewRandomBase32ID`.
func NewRandomBase62ID(len int) string {
	b := make([]byte, len)
	rand.Read(b) // rand.Read is guaranteed to not return errors.
	return EncodeToBase62String(b)
}
