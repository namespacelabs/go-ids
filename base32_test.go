package ids

import (
	"strings"
	"testing"
)

func TestBase32Encoding(t *testing.T) {
	input := "Hello, world!"

	encoded := EncodeToBase32String([]byte(input))

	if strings.ToLower(encoded) != encoded {
		t.Fatalf("expected %q to be a lowercase string", encoded)
	}

	decoded, err := DecodeBase32String(encoded)
	if err != nil {
		t.Fatalf("failed to decode %q: %v", encoded, err)
	}

	if string(decoded) != input {
		t.Fatalf("Decoded output %q does not match input %q", string(decoded), input)
	}
}
